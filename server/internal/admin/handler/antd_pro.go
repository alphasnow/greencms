// @author AlphaSnow

package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/internal/admin/middleware"
	"server/internal/admin/schema"
	"server/internal/admin/service"
	"server/internal/core/model/dao"
	"server/internal/core/model/entity"
	"server/pkg/g"
	utils2 "server/pkg/utils"
	"server/pkg/xjwt"
)

type AntdProHandler struct {
	AuthService   *service.AuthService
	AccessService *service.AccessService
	UserToken     *xjwt.UserToken
}

func NewAntdProHandler() *AntdProHandler {
	h := new(AntdProHandler)
	h.AuthService = new(service.AuthService)
	h.UserToken = g.JWT(dao.AdminUser.TableName())
	h.AccessService = new(service.AccessService)
	return h
}

//func (h *AntdProHandler) Routes(g gin.IRouter) {
//	// get /api/currentUser
//	g.GET("/currentUser", h.CurrentUser)
//	// post /api/login/outLogin
//	g.POST("/login/outLogin", h.OutLogin)
//	// post /api/login/account
//	g.POST("/login/account", h.Account)
//	//// get /api/notices
//	//g.GET("/notices", h.Notices)
//	//// get /api/rule
//	//g.GET("/rule", h.GetRule)
//	//// post /api/rule
//	//g.POST("/rule", h.UpdateRule)
//}

func (h *AntdProHandler) CurrentUser(c *gin.Context) {
	id := middleware.GetUserID(c)
	user, err := dao.AdminUser.WithContext(c).Where(dao.AdminUser.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorResp("账号失效", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, schema.CurrentUser{
		Name:     user.Realname,
		Avatar:   user.AvatarURL.String(),
		UserID:   fmt.Sprintf("%d", user.ID),
		Access:   user.Access,
		Username: user.Username,
	})
}
func (h *AntdProHandler) Account(c *gin.Context) {
	req := new(schema.LoginParams)
	if err := c.ShouldBindBodyWithJSON(req); err != nil {
		schema.ErrorJson(c, "请求参数错误", schema.WithError(err))
		return
	}

	user, err := h.AuthService.Login(c, req)
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}

	token, err := h.UserToken.GenerateID(user.ID)
	if err != nil {
		schema.ErrorJson(c, "授权失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, schema.LoginResult{
		Status:           "ok",
		Token:            token,
		Type:             "account",
		CurrentAuthority: user.Access,
	})
}
func (h *AntdProHandler) OutLogin(c *gin.Context) {
	schema.SuccessJson(c, map[string]any{})
}

func (h *AntdProHandler) AccountSettings(c *gin.Context) {
	req := new(schema.AccountSettingsReq)
	if err := c.ShouldBindBodyWithJSON(req); err != nil {
		schema.ErrorJson(c, "请求参数错误", schema.WithError(err))
		return
	}

	user := entity.AdminUser{}
	if req.Password != "" {
		user.Password, _ = utils2.PasswordHash(req.Password)
	}

	userId := middleware.GetUserID(c)
	if _, err := dao.AdminUser.WithContext(c).Where(dao.AdminUser.ID.Eq(userId)).Updates(user); err != nil {
		schema.ErrorJson(c, "设置失败", schema.WithError(err))
		return
	}

	schema.SuccessJson(c, nil)
}

//func (h *AntdProHandler) Notices(c *gin.Context) {
//	schema.SuccessJson(c, []any{})
//}
//
//func (h *AntdProHandler) GetRule(c *gin.Context) {
//	schema.SuccessJson(c, []any{})
//}
//
//func (h *AntdProHandler) UpdateRule(c *gin.Context) {
//	schema.SuccessJson(c, []any{})
//}
