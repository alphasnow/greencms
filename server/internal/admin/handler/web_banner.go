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

type WebBannerHandler struct {
}

func NewWebBannerHandler() *WebBannerHandler {
	h := new(WebBannerHandler)
	return h
}

func (h *WebBannerHandler) Routes(r gin.IRouter) {
	// 增删改后清理缓存
	cc := func(c *gin.Context) {
		_ = g.Cache().Delete(c, constant.WebGlobalStateCacheKey)
	}

	r.GET("/web-banner/index", h.Index)
	r.GET("/web-banner/show/:id", h.Show)
	r.POST("/web-banner/create", h.Store, cc)
	r.POST("/web-banner/edit/:id", h.Update, cc)
	r.POST("/web-banner/delete/:id", h.Delete, cc)

	r.GET("/web-banner/options", h.Options)
}

func (h *WebBannerHandler) Index(c *gin.Context) {
	req := new(schema.PageParams)
	if err := c.ShouldBindQuery(req); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	searches := utils.GetPageListConditions(c, dao.WebBanner.TableName())
	orders := utils.GetPageListOrders(c, dao.WebBanner.TableName(), dao.WebBanner.ID.Desc())
	data, total, _ := dao.WebBanner.
		WithContext(c).
		Where(searches...).
		Order(orders...).
		FindByPage(req.Offset(), req.Limit())

	schema.PageListJson(c, data, total)
}

func (h *WebBannerHandler) Show(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	data, err := dao.WebBanner.WithContext(c).Where(dao.WebBanner.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "数据查询为空", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, data)
}

func (h *WebBannerHandler) Store(c *gin.Context) {
	data := new(entity.WebBanner)
	if err := c.ShouldBindJSON(data); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	if err := dao.WebBanner.WithContext(c).Create(data); err != nil {
		schema.ErrorJson(c, "数据存储失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, data)
}

func (h *WebBannerHandler) Update(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	data := new(entity.WebBanner)
	if err = c.ShouldBindJSON(data); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	_, err = dao.WebBanner.WithContext(c).Where(dao.WebBanner.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "数据查询为空", schema.WithError(err))
		return
	}

	_, err = dao.WebBanner.WithContext(c).Where(dao.WebBanner.ID.Eq(id)).Updates(data)
	if err != nil {
		schema.ErrorJson(c, "数据更新失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, nil)
}

func (h *WebBannerHandler) Delete(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	if _, err = dao.WebBanner.WithContext(c).Where(dao.WebBanner.ID.Eq(id)).Delete(); err != nil {
		schema.ErrorJson(c, "数据删除失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, nil)
}

func (h *WebBannerHandler) Options(c *gin.Context) {
	data := []schema.SelectOption{
		{"首页展示", constant.WebBannerHomeRight},
	}
	schema.SuccessJson(c, gin.H{
		"banner_group": data,
	})
}
