// @author AlphaSnow

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"server/internal/admin/schema"
	"server/internal/admin/transform"
	"server/internal/admin/utils"
	"server/internal/core/model/dao"
	"server/internal/core/model/entity"
)

type ArticleCategoryHandler struct {
}

func NewArticleCategoryHandler() *ArticleCategoryHandler {
	h := new(ArticleCategoryHandler)
	return h
}

func (h *ArticleCategoryHandler) Routes(g gin.IRouter) {
	g.GET("/article-category/index", h.Index)
	g.GET("/article-category/show/:id", h.Show)
	g.POST("/article-category/create", h.Store)
	g.POST("/article-category/edit/:id", h.Update)
	g.POST("/article-category/delete/:id", h.Delete)

	g.GET("/article-category/options", h.Options)
}

func (h *ArticleCategoryHandler) Index(c *gin.Context) {
	req := new(schema.PageParams)
	if err := c.ShouldBindQuery(req); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	searches := utils.GetPageListConditions(c, dao.ArticleCategory.TableName())
	orders := utils.GetPageListOrders(c, dao.ArticleCategory.TableName(), dao.ArticleCategory.ID.Desc())
	data, total, _ := dao.ArticleCategory.
		WithContext(c).
		Where(searches...).
		Order(orders...).
		FindByPage(req.Offset(), req.Limit())

	schema.PageListJson(c, data, total)
}

func (h *ArticleCategoryHandler) Show(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	data, err := dao.ArticleCategory.WithContext(c).Where(dao.ArticleCategory.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "数据查询为空", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, data)
}

func (h *ArticleCategoryHandler) Store(c *gin.Context) {
	data := new(entity.ArticleCategory)
	if err := c.ShouldBindJSON(data); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	if err := dao.ArticleCategory.WithContext(c).Create(data); err != nil {
		schema.ErrorJson(c, "数据存储失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, data)
}

func (h *ArticleCategoryHandler) Update(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	data := new(entity.ArticleCategory)
	if err = c.ShouldBindJSON(data); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	_, err = dao.ArticleCategory.WithContext(c).Where(dao.ArticleCategory.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "数据查询为空", schema.WithError(err))
		return
	}

	_, err = dao.ArticleCategory.WithContext(c).Where(dao.ArticleCategory.ID.Eq(id)).Updates(data)
	if err != nil {
		schema.ErrorJson(c, "数据更新失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, nil)
}

func (h *ArticleCategoryHandler) Delete(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	if _, err = dao.ArticleCategory.WithContext(c).Where(dao.ArticleCategory.ID.Eq(id)).Delete(); err != nil {
		schema.ErrorJson(c, "数据删除失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, nil)
}

func (h *ArticleCategoryHandler) Options(c *gin.Context) {
	res, _ := dao.ArticleCategory.WithContext(c).Select(
		dao.ArticleCategory.ID,
		dao.ArticleCategory.Title,
	).Order(dao.ArticleCategory.ID.Desc()).Find()

	data := make([]schema.SelectOption, len(res))
	opt := transform.SelectCopyOption(entity.ArticleCategory{}, schema.SelectOption{}, "Title", "ID")
	_ = copier.CopyWithOption(&data, &res, opt)

	schema.SuccessJson(c, data)
}
