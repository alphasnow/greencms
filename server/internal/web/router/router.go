// @author AlphaSnow

package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/web/handler"
)

func Register(r *gin.Engine) {
	wr := r.Group("/api/web")

	wr.GET("/state", new(handler.StateHandler).Data)
	wr.GET("/home", new(handler.HomeHandler).Data)
	wr.GET("/article", new(handler.ArticleHandler).List)
	wr.PUT("/statistic/:id/:type", new(handler.ArticleHandler).Statistic)
	wr.GET("/article/:id", new(handler.ArticleHandler).Show)
	//wr.GET("/sitemap", new(handler.SEOHandler).Sitemap)
	wr.GET("/category/:id", new(handler.CategoryHandler).Show)
	wr.GET("/tag/:id", new(handler.TagHandler).Show)
}
