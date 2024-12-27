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

type ArticleTagHandler struct {
}

func NewArticleTagHandler() *ArticleTagHandler {
	h := new(ArticleTagHandler)
	return h
}

func (h *ArticleTagHandler) Routes(g gin.IRouter) {
	g.GET("/article-tag/index", h.Index)
	g.GET("/article-tag/show/:id", h.Show)
	g.POST("/article-tag/create", h.Store)
	g.POST("/article-tag/edit/:id", h.Update)
	g.POST("/article-tag/delete/:id", h.Delete)

	g.GET("/article-tag/options", h.Options)
}

func (h *ArticleTagHandler) Index(c *gin.Context) {
	req := new(schema.PageParams)
	if err := c.ShouldBindQuery(req); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	searches := utils.GetPageListConditions(c, dao.ArticleTag.TableName())
	orders := utils.GetPageListOrders(c, dao.ArticleTag.TableName(), dao.ArticleTag.ID.Desc())
	data, total, _ := dao.ArticleTag.
		WithContext(c).
		Where(searches...).
		Order(orders...).
		FindByPage(req.Offset(), req.Limit())

	schema.PageListJson(c, data, total)
}

func (h *ArticleTagHandler) Show(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	data, err := dao.ArticleTag.WithContext(c).Where(dao.ArticleTag.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "数据查询为空", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, data)
}

func (h *ArticleTagHandler) Store(c *gin.Context) {
	data := new(entity.ArticleTag)
	if err := c.ShouldBindJSON(data); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	if err := dao.ArticleTag.WithContext(c).Create(data); err != nil {
		schema.ErrorJson(c, "数据存储失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, data)
}

func (h *ArticleTagHandler) Update(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	data := new(entity.ArticleTag)
	if err = c.ShouldBindJSON(data); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	_, err = dao.ArticleTag.WithContext(c).Where(dao.ArticleTag.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "数据查询为空", schema.WithError(err))
		return
	}

	_, err = dao.ArticleTag.WithContext(c).Where(dao.ArticleTag.ID.Eq(id)).Updates(data)
	if err != nil {
		schema.ErrorJson(c, "数据更新失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, nil)
}

func (h *ArticleTagHandler) Delete(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	if _, err = dao.ArticleTag.WithContext(c).Where(dao.ArticleTag.ID.Eq(id)).Delete(); err != nil {
		schema.ErrorJson(c, "数据删除失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, nil)
}

func (h *ArticleTagHandler) Options(c *gin.Context) {
	res, _ := dao.ArticleTag.WithContext(c).Select(
		dao.ArticleTag.ID,
		dao.ArticleTag.Name,
	).Order(dao.ArticleTag.ID.Desc()).Find()

	data := make([]schema.SelectOption, len(res))
	opt := transform.SelectCopyOption(entity.ArticleTag{}, schema.SelectOption{}, "Name", "ID")
	_ = copier.CopyWithOption(&data, &res, opt)

	schema.SuccessJson(c, data)
}
