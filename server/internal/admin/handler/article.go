// @author AlphaSnow

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"server/internal/admin/middleware"
	"server/internal/admin/schema"
	"server/internal/admin/utils"
	"server/internal/core/model/dao"
	"server/internal/core/model/entity"
)

type ArticleHandler struct {
}

func NewArticleHandler() *ArticleHandler {
	h := new(ArticleHandler)
	return h
}

func (h *ArticleHandler) Routes(g gin.IRouter) {
	g.GET("/article/index", h.Index)
	g.GET("/article/show/:id", h.Show)
	g.POST("/article/create", h.Store)
	g.POST("/article/edit/:id", h.Update)
	g.POST("/article/delete/:id", h.Delete)
}

func (h *ArticleHandler) Index(c *gin.Context) {
	req := new(schema.PageParams)
	if err := c.ShouldBindQuery(req); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	// 按分类 按标题查询
	// searches := utils.GetPageListConditions(c, dao.Article.TableName())
	searches := make([]gen.Condition, 0)
	if tit, _ := c.GetQuery("title"); tit != "" {
		search := dao.Article.Title.Like("%" + tit + "%")
		searches = append(searches, search)
	}
	if cat, _ := c.GetQuery("category_name"); cat != "" {
		search := dao.ArticleCategory.WithContext(c).Where(
			dao.Article.WithContext(c).Columns(dao.Article.CategoryID).Eq(
				dao.ArticleCategory.WithContext(c).Select(dao.ArticleCategory.ID).Where(dao.ArticleCategory.Title.Eq(cat)),
			),
		)
		searches = append(searches, search)
	}

	orders := utils.GetPageListOrders(c, dao.Article.TableName(), dao.Article.ID.Desc())
	data, total, _ := dao.Article.
		WithContext(c).
		Preload(
			dao.Article.ArticleCategory,
			dao.Article.ArticleTags,
			dao.Article.ArticleStatistic,
		).
		Where(searches...).
		Order(orders...).
		FindByPage(req.Offset(), req.Limit())

	var resp []*schema.ArticleIndexResp
	_ = copier.Copy(&resp, data)

	schema.PageListJson(c, resp, total)
}

func (h *ArticleHandler) Show(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	data, err := dao.Article.WithContext(c).
		Preload(
			dao.Article.ArticleCategory,
			dao.Article.ArticleTags,
			dao.Article.ArticleContent,
		).Where(dao.Article.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "数据查询为空", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, data)
}

func (h *ArticleHandler) Store(c *gin.Context) {
	data := new(entity.Article)
	if err := c.ShouldBindJSON(data); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	userID := middleware.GetUserID(c)
	data.AdminID = userID

	if err := dao.Article.WithContext(c).Create(data); err != nil {
		schema.ErrorJson(c, "数据存储失败", schema.WithError(err))
		return
	}

	if data.ArticleStatistic == nil {
		statistic := &entity.ArticleStatistic{ArticleID: data.ID}
		_ = dao.ArticleStatistic.WithContext(c).Create(statistic)
		data.ArticleStatistic = statistic
	}

	schema.SuccessJson(c, data)
}

func (h *ArticleHandler) Update(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	data := new(entity.Article)
	if err = c.ShouldBindJSON(data); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	info, err := dao.Article.WithContext(c).Where(dao.Article.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "数据查询为空", schema.WithError(err))
		return
	}

	//err = dao.Q.Transaction(func(tx *dao.Query) error {
	//	if _, err = tx.ArticleContent.WithContext(c).
	//		Where(tx.ArticleContent.ArticleID.Eq(id)).
	//		Updates(data.ArticleContent); err != nil {
	//		return err
	//	}
	//	// 使用 sqlite 发生 database is locked
	//	if err = tx.Article.ArticleTags.Model(info).
	//		Replace(data.ArticleTags...); err != nil {
	//		return err
	//	}
	//	if _, err = tx.Article.WithContext(c).
	//		Omit(field.AssociationFields).
	//		Where(tx.Article.ID.Eq(id)).
	//		Updates(data); err != nil {
	//		return err
	//	}
	//	return nil
	//})
	err = func() error {
		if _, err = dao.ArticleContent.WithContext(c).
			Where(dao.ArticleContent.ArticleID.Eq(id)).
			Updates(data.ArticleContent); err != nil {
			return err
		}
		if err = dao.Article.ArticleTags.Model(info).
			Replace(data.ArticleTags...); err != nil {
			return err
		}
		if _, err = dao.Article.WithContext(c).
			Omit(field.AssociationFields).
			Where(dao.Article.ID.Eq(id)).
			Updates(data); err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		schema.ErrorJson(c, "数据更新失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, nil)
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	if _, err = dao.Article.WithContext(c).Where(dao.Article.ID.Eq(id)).Delete(); err != nil {
		schema.ErrorJson(c, "数据删除失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, nil)
}
