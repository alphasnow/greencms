// @author AlphaSnow

package handler

import (
	"github.com/gin-gonic/gin"
	"server/internal/admin/schema"
	"server/internal/admin/utils"
	"server/internal/core/constant"
	"server/internal/core/model/dao"
	"server/internal/core/model/entity"
	"server/pkg/g"
)

type WebMetaHandler struct {
}

func NewWebMetaHandler() *WebMetaHandler {
	h := new(WebMetaHandler)
	return h
}

func (h *WebMetaHandler) Routes(r gin.IRouter) {
	// 增删改后清理缓存
	cc := func(c *gin.Context) {
		_ = g.Cache().Delete(c, constant.WebGlobalStateCacheKey)
	}

	r.GET("/web-meta/index", h.Index)
	r.GET("/web-meta/show/:id", h.Show)
	r.POST("/web-meta/create", h.Store, cc)
	r.POST("/web-meta/edit/:id", h.Update, cc)
	r.POST("/web-meta/delete/:id", h.Delete, cc)

	r.GET("/web-meta/options", h.Options)
}

func (h *WebMetaHandler) Index(c *gin.Context) {
	req := new(schema.PageParams)
	if err := c.ShouldBindQuery(req); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	searches := utils.GetPageListConditions(c, dao.WebMeta.TableName())
	orders := utils.GetPageListOrders(c, dao.WebMeta.TableName(), dao.WebMeta.ID.Desc())
	data, total, _ := dao.WebMeta.
		WithContext(c).
		Where(searches...).
		Order(orders...).
		FindByPage(req.Offset(), req.Limit())

	schema.PageListJson(c, data, total)
}

func (h *WebMetaHandler) Show(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	data, err := dao.WebMeta.WithContext(c).Where(dao.WebMeta.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "数据查询为空", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, data)
}

func (h *WebMetaHandler) Store(c *gin.Context) {
	data := new(entity.WebMeta)
	if err := c.ShouldBindJSON(data); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	if err := dao.WebMeta.WithContext(c).Create(data); err != nil {
		schema.ErrorJson(c, "数据存储失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, data)
}

func (h *WebMetaHandler) Update(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	data := new(entity.WebMeta)
	if err = c.ShouldBindJSON(data); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	_, err = dao.WebMeta.WithContext(c).Where(dao.WebMeta.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "数据查询为空", schema.WithError(err))
		return
	}

	_, err = dao.WebMeta.WithContext(c).Where(dao.WebMeta.ID.Eq(id)).Updates(data)
	if err != nil {
		schema.ErrorJson(c, "数据更新失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, nil)
}

func (h *WebMetaHandler) Delete(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	info, err := dao.WebMeta.WithContext(c).Where(dao.WebMeta.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "数据查询失败", schema.WithError(err))
		return
	}

	if info.MetaGroup == constant.WebMetaDefaultGroup {
		schema.ErrorJson(c, "默认数据不能删除")
		return
	}

	if _, err = dao.WebMeta.WithContext(c).Where(dao.WebMeta.ID.Eq(id)).Delete(); err != nil {
		schema.ErrorJson(c, "数据删除失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, nil)
}

func (h *WebMetaHandler) Options(c *gin.Context) {
	data := []schema.SelectOption{
		// {"默认数据", constant.WebMetaDefaultGroup},
		{"友情链接", "friendly_link"},
		{"自定义数据", "custom_data"},
	}
	schema.SuccessJson(c, gin.H{
		"meta_group": data,
	})
}
