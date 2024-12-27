// @author AlphaSnow

package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"server/internal/core/model/dao"
	"server/internal/web/ecode"
	"server/internal/web/schema"
	"strconv"
)

type ArticleHandler struct {
}

// List
// @Summary 文章列表
// @Description 按分页或分类或标签或关键词获取文章列表数据
// @Tags Web
// @Param request query schema.ArticleListReq true "query"
// @Success 200	{object} schema.ArticleListResp "ok"
// @Router /api/web/article [get]
func (h *ArticleHandler) List(c *gin.Context) {
	req := new(schema.ArticleListReq)
	if err := c.ShouldBindQuery(req); err != nil {
		schema.Error(c, ecode.BadRequest, err)
		return
	}

	wheres := []gen.Condition{}
	if req.Keywords != "" {
		// https://gorm.io/gen/query.html#Group-Conditions
		pd := dao.Article.WithContext(c)
		where := pd.Where(pd.Where(dao.Article.Title.Like("%" + req.Keywords + "%"))).
			Or(pd.Where(dao.Article.Description.Like("%" + req.Keywords + "%")))
		wheres = append(wheres, where)
	}
	if req.CategoryID != 0 {
		where := dao.Article.CategoryID.Eq(req.CategoryID)
		wheres = append(wheres, where)
	}
	if req.TagID != 0 {
		where := dao.Article.WithContext(c).Columns(dao.Article.ID).In(
			dao.ArticleTagRelate.WithContext(c).Select(dao.ArticleTagRelate.ArticleID).Where(dao.ArticleTagRelate.TagID.Eq(req.TagID)),
		)
		wheres = append(wheres, where)
	}

	articles, err := dao.Article.WithContext(c).
		Preload(dao.Article.ArticleStatistic, dao.Article.ArticleTags, dao.Article.ArticleCategory).
		Order(dao.Article.Sort.Asc(), dao.Article.ID.Desc()).
		Where(wheres...).
		Limit(req.GetLimit() + 1).
		Offset(req.GetOffset()).
		Find()
	if err != nil {
		schema.Error(c, ecode.DatabaseError, err)
		return
	}

	moreArticles := false
	if len(articles) > req.GetLimit() {
		moreArticles = true
	}

	resp := new(schema.ArticleListResp)
	resp.Articles = articles
	resp.MoreArticles = moreArticles
	schema.Success(c, resp)
}

// Show
// @Summary 查看文章
// @Description 根据文章序号获取文章详细数据
// @Tags Web
// @Param id path int true "文章ID"
// @Success 200	{object} entity.Article "ok"
// @Router /api/web/article/{id} [get]
func (h *ArticleHandler) Show(c *gin.Context) {
	idP := c.Param("id")
	id, err := strconv.ParseUint(idP, 10, 0)
	if err != nil {
		schema.Error(c, ecode.BadRequest, err)
		return
	}

	// id获取当前文章
	article, err := dao.Article.WithContext(c).
		Preload(field.Associations).
		Where(dao.Article.ID.Eq(uint(id))).
		Take()
	if err != nil {
		schema.Error(c, ecode.ModelNotFound, err)
		return
	}

	schema.Success(c, article)
}

// Statistic
// @Summary 文章统计
// @Description 更新文章统计数据,如查看量
// @Tags Web
// @Param id path int true "文章ID"
// @Param type path string true "统计类型:views,favourites"
// @Success 200	{object} schema.ErrorResp "ok"
// @Router /api/web/statistic/{id}/{type} [put]
func (h *ArticleHandler) Statistic(c *gin.Context) {
	idP := c.Param("id")
	id, err := strconv.ParseUint(idP, 10, 0)
	if err != nil {
		schema.Error(c, ecode.BadRequest, err)
		return
	}
	typeField := c.Param("type")

	_, err = dao.ArticleStatistic.WithContext(c).
		Where(dao.ArticleStatistic.ArticleID.Eq(uint(id))).
		Update(field.NewString(dao.ArticleStatistic.TableName(), typeField), gorm.Expr(typeField+" +?", 1))
	if err != nil {
		schema.Error(c, ecode.ModelNotFound, err)
		return
	}

	schema.Success(c, schema.ErrorResp{})
}
