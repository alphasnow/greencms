// @author AlphaSnow

package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/admin/handler"
	"server/internal/admin/middleware"
)

func Register(r *gin.Engine) {
	// middleware
	tokenMiddleware := middleware.NewAuthToken([]string{
		// 跳过登录
		"/api/admin/login/account",
	})
	accessMiddleware := middleware.NewUserAccess()

	// 公共服务
	{
		h := handler.NewAntdProHandler()
		r.GET("/api/admin/currentUser", tokenMiddleware, h.CurrentUser)
		r.POST("/api/admin/login/outLogin", tokenMiddleware, h.OutLogin)
		r.POST("/api/admin/account/settings", tokenMiddleware, h.AccountSettings)
		r.POST("/api/admin/login/account", h.Account)
	}
	{
		h := handler.NewUploadHandler()
		r.POST("/api/admin/upload/form-file", tokenMiddleware, h.FormFile)
	}

	// group
	adminRouter := r.Group("/api/admin", tokenMiddleware, accessMiddleware)

	// antd pro 模型路由
	handler.NewArticleCategoryHandler().Routes(adminRouter)
	handler.NewArticleHandler().Routes(adminRouter)
	handler.NewArticleTagHandler().Routes(adminRouter)

	handler.NewWebBannerHandler().Routes(adminRouter)
	handler.NewWebMetaHandler().Routes(adminRouter)

	handler.NewAdminUserHandler().Routes(adminRouter)
}
