// @author AlphaSnow

package handler

import (
	"github.com/gin-gonic/gin"
	"server/internal/core/model/dao"
	"server/internal/core/model/entity"
	"server/internal/web/ecode"
	"server/internal/web/schema"
)

type HomeHandler struct {
}

// Data
// @Summary 首页数据
// @Description 一次性获取首页所需所有数据
// @Tags Web
// @Success 200	{object} schema.HomeDataResp "ok"
// @Router /api/web/home [get]
func (h *HomeHandler) Data(c *gin.Context) {
	resp, err := h.getDatabaseData(c)
	if err != nil {
		schema.Error(c, ecode.InternalServerError, err)
		return
	}

	schema.Success(c, resp)
}

func (h *HomeHandler) getDatabaseData(c *gin.Context) (resp *schema.HomeDataResp, err error) {
	// 焦点图 5个
	resp = new(schema.HomeDataResp)
	banners, err := dao.WebBanner.WithContext(c).Order(dao.WebBanner.Sort.Asc(), dao.WebBanner.ID.Desc()).Limit(5).Find()
	if err != nil {
		return nil, err
	}
	resp.Banners = banners

	// tag热榜 10个
	tags, err := dao.ArticleTag.WithContext(c).Order(dao.ArticleTag.ID.Desc()).Limit(20).Find()
	if err != nil {
		return nil, err
	}
	resp.Tags = tags

	// 热榜 5个
	hotArticleIDs := make([]uint, 5)
	err = dao.ArticleStatistic.WithContext(c).
		Order(dao.ArticleStatistic.Views.Desc()).
		Limit(5).
		Pluck(dao.ArticleStatistic.ArticleID, &hotArticleIDs)
	if err != nil {
		return nil, err
	}
	hotArticles, err := dao.Article.WithContext(c).Preload(dao.Article.ArticleStatistic).Where(dao.Article.ID.In(hotArticleIDs...)).Find()
	if err != nil {
		return nil, err
	}
	// 按ID排序
	sortHotArticles := make([]*entity.Article, len(hotArticles))
	for k, id := range hotArticleIDs {
		for _, art := range hotArticles {
			if art.ID == id {
				sortHotArticles[k] = art
				break
			}
		}
	}
	resp.HotArticles = sortHotArticles

	// 文章 5个 以及是否 有更多
	var articles []*entity.Article
	articles, err = dao.Article.WithContext(c).
		Preload(dao.Article.ArticleStatistic, dao.Article.ArticleTags, dao.Article.ArticleCategory).
		Order(dao.Article.Sort.Asc(), dao.Article.ID.Desc()).
		Limit(6).
		Find()
	if err != nil {
		return nil, err
	}
	more := false
	if len(articles) > 5 {
		more = true
		articles = articles[:len(articles)-1]
	}

	resp.Articles = articles
	resp.MoreArticles = more
	return resp, nil
}
