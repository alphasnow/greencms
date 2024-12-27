// @author AlphaSnow

package schema

import (
	"server/internal/core/model/entity"
)

type StateDataResp struct {
	//ErrorResp  `json:",inline"`
	Metas      map[string]string         `json:"metas"`
	Categories []*entity.ArticleCategory `json:"categories"`
}

type HomeDataResp struct {
	//ErrorResp    `json:",inline"`
	Banners      []*entity.WebBanner  `json:"banners"`
	Articles     []*entity.Article    `json:"articles"`
	MoreArticles bool                 `json:"more_articles"`
	Tags         []*entity.ArticleTag `json:"tags"`
	HotArticles  []*entity.Article    `json:"hot_articles"`
}
