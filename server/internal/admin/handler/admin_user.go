// @author AlphaSnow

package handler

import (
	"github.com/gin-gonic/gin"
	"server/internal/admin/middleware"
	"server/internal/admin/schema"
	"server/internal/admin/service"
	"server/internal/admin/utils"
	"server/internal/core/model/dao"
	"server/internal/core/model/entity"
	utils2 "server/pkg/utils"
)

type AdminUserHandler struct {
	AccessService *service.AccessService
}

func NewAdminUserHandler() *AdminUserHandler {
	h := new(AdminUserHandler)
	h.AccessService = new(service.AccessService)
	return h
}

func (h *AdminUserHandler) Routes(g gin.IRouter) {

	//checkAccess := func(c *gin.Context) {
	//	if access := h.AccessService.GetRoleByID(c, middleware.GetUserID(c)); access != service.AdminRole {
	//		schema.ErrorJson(c, "仅限超级管理员操作")
	//		return
	//	}
	//	c.Next()
	//}

	g.GET("/admin-user/index", h.Index)
	g.GET("/admin-user/show/:id", h.Show)
	g.POST("/admin-user/create", h.Store)
	g.POST("/admin-user/edit/:id", h.Update)
	g.POST("/admin-user/delete/:id", h.Delete)

	g.GET("/admin-user/options", h.Options)
}

func (h *AdminUserHandler) Index(c *gin.Context) {
	req := new(schema.PageParams)
	if err := c.ShouldBindQuery(req); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	searches := utils.GetPageListConditions(c, dao.AdminUser.TableName())
	orders := utils.GetPageListOrders(c, dao.AdminUser.TableName(), dao.AdminUser.ID.Desc())
	data, total, _ := dao.AdminUser.
		WithContext(c).
		Where(searches...).
		Order(orders...).
		FindByPage(req.Offset(), req.Limit())

	schema.PageListJson(c, data, total)
}

func (h *AdminUserHandler) Show(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	data, err := dao.AdminUser.WithContext(c).Where(dao.AdminUser.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "数据查询为空", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, data)
}

func (h *AdminUserHandler) Store(c *gin.Context) {
	data := new(entity.AdminUser)
	if err := c.ShouldBindJSON(data); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	// 超级管理员才能创建root
	if data.Access == "root" && middleware.GetUserID(c) != service.AdminSuperID {
		schema.ErrorJson(c, "禁止创建超级管理")
		return
	}

	// username 不能重复
	if _, err := dao.AdminUser.WithContext(c).Where(dao.AdminUser.Username.Eq(data.Username)).Take(); err == nil {
		schema.ErrorJson(c, "账号已被使用,请更换重试", schema.WithError(err))
		return
	}
	data.Password, _ = utils2.PasswordHash(data.Password)

	if err := dao.AdminUser.WithContext(c).Create(data); err != nil {
		schema.ErrorJson(c, "数据存储失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, data)
}

func (h *AdminUserHandler) Update(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	// 超级管理员不能被自己外的其他人修改
	if id == service.AdminSuperID && middleware.GetUserID(c) != service.AdminSuperID {
		schema.ErrorJson(c, "禁止修改超级管理")
		return
	}

	data := new(entity.AdminUser)
	if err = c.ShouldBindJSON(data); err != nil {
		schema.ErrorJson(c, "请求数据错误", schema.WithError(err))
		return
	}

	// username 不能重复
	if _, err = dao.AdminUser.WithContext(c).Where(dao.AdminUser.ID.Neq(id), dao.AdminUser.Username.Eq(data.Username)).Take(); err == nil {
		schema.ErrorJson(c, "账号已被使用,请更换重试", schema.WithError(err))
		return
	}

	db := dao.AdminUser.WithContext(c)
	if data.Password != "" {
		data.Password, _ = utils2.PasswordHash(data.Password)
	} else {
		db = db.Omit(dao.AdminUser.Password)
	}

	_, err = db.Where(dao.AdminUser.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "数据查询为空", schema.WithError(err))
		return
	}

	_, err = dao.AdminUser.WithContext(c).Where(dao.AdminUser.ID.Eq(id)).Updates(data)
	if err != nil {
		schema.ErrorJson(c, "数据更新失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, data)
}

func (h *AdminUserHandler) Delete(c *gin.Context) {
	id, err := utils.GetParamID(c)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	// 超级管理员不能删除
	if id == service.AdminSuperID {
		schema.ErrorJson(c, "不能删除超级管理")
		return
	}

	if _, err = dao.AdminUser.WithContext(c).Where(dao.AdminUser.ID.Eq(id)).Delete(); err != nil {
		schema.ErrorJson(c, "数据删除失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, nil)
}

func (h *AdminUserHandler) Options(c *gin.Context) {
	data := []schema.SelectOption{
		{"后台管理", "admin"},
		{"数据管理", "manager"},
		{"文章编辑", "editor"},
	}
	if middleware.GetUserID(c) != service.AdminSuperID {
		schema.SuccessJson(c, gin.H{
			"access": data,
		})
		return
	}

	newData := make([]schema.SelectOption, len(data)+1)
	newData[0] = schema.SelectOption{
		"超级管理员", "root",
	}
	copy(newData[1:], data)
	schema.SuccessJson(c, gin.H{
		"access": newData,
	})
}
