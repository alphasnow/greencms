// @author AlphaSnow

package schema

import (
	"server/internal/core/model/accessor"
	"server/internal/core/model/entity"
	"strings"
	"time"
)

type ArticleIndexResp struct {
	ID           uint               `json:"id"`
	Title        string             `json:"title"`
	ImageURL     accessor.StaticUrl `json:"image_url"`
	CategoryName string             `json:"category_name"`
	TagNames     string             `json:"tag_names"`
	Favourites   uint               `json:"favourites"`
	Views        uint               `json:"views"`
	Sort         uint               `json:"sort"`
	UpdatedAt    time.Time          `json:"updated_at"` // updated_at
}

func (a *ArticleIndexResp) ArticleTags(tags []*entity.ArticleTag) {
	names := make([]string, len(tags))
	for i, tag := range tags {
		names[i] = tag.Name
	}
	a.TagNames = strings.Join(names, ",")
}

func (a *ArticleIndexResp) ArticleCategory(category *entity.ArticleCategory) {
	a.CategoryName = category.Title
}

func (a *ArticleIndexResp) ArticleStatistic(statistic *entity.ArticleStatistic) {
	if statistic == nil {
		return
	}
	a.Views = statistic.Views
	a.Favourites = statistic.Favourites
}
