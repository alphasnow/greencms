// @author AlphaSnow

package schema

import "server/internal/core/model/entity"

type ArticleListReq struct {
	Keywords   string `form:"keywords" json:"keywords"`       // 关键词
	CategoryID uint   `form:"category_id" json:"category_id"` // 分类ID
	TagID      uint   `form:"tag_id" json:"tag_id"`           // 标签ID
	Page       int    `form:"page" json:"page" default:"1"`   // 分页编号
	Size       int    `form:"size" json:"size" default:"5"`   // 分页数量
}

func (r *ArticleListReq) GetOffset() int {
	if r.Page <= 1 {
		return 0
	}
	return (r.Page - 1) * r.GetLimit()
}
func (r *ArticleListReq) GetLimit() int {
	if r.Size <= 0 {
		return 5
	}
	return r.Size
}

type ArticleListResp struct {
	//ErrorResp    `json:",inline"`
	Articles     []*entity.Article `json:"articles"`      // 文章数据
	MoreArticles bool              `json:"more_articles"` // 是否有更多文章
}

//type ArticleShowResp struct {
//	ErrorResp    `json:",inline"`
//	Articles     *entity.Article `json:"articles"`
//}
//type ArticleCategoryShowResp struct {
//	ErrorResp    `json:",inline"`
//	Articles     []*entity.Article `json:"articles"`
//	MoreArticles bool              `json:"more_articles"`
//}
//type ArticleTagShowResp struct {
//	ErrorResp    `json:",inline"`
//	Articles     []*entity.Article `json:"articles"`
//	MoreArticles bool              `json:"more_articles"`
//}
