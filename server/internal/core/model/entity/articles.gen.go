// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"server/internal/core/model/accessor"

	"gorm.io/gorm"
)

const TableNameArticle = "articles"

// Article 文章
type Article struct {
	ID               uint               `gorm:"column:id;primaryKey;autoIncrement:true;comment:文章管理ID" json:"id"` // 文章管理ID
	CreatedAt        time.Time          `gorm:"column:created_at;<-:create;comment:添加时间" json:"created_at"`       // 添加时间
	UpdatedAt        time.Time          `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        gorm.DeletedAt     `gorm:"column:deleted_at" json:"-"`
	CategoryID       uint               `gorm:"column:category_id;comment:分类id" json:"category_id"`     // 分类id
	Title            string             `gorm:"column:title;comment:文章标题" json:"title"`                 // 文章标题
	ImageURL         accessor.StaticUrl `gorm:"column:image_url;comment:文章图片" json:"image_url"`         // 文章图片
	OriginURL        string             `gorm:"column:origin_url;comment:原文链接" json:"origin_url"`       // 原文链接
	OriginAuthor     string             `gorm:"column:origin_author;comment:原文作者" json:"origin_author"` // 原文作者
	AdminID          uint               `gorm:"column:admin_id;comment:管理员id" json:"admin_id"`          // 管理员id
	Keywords         string             `gorm:"column:keywords;comment:SEO关键词" json:"keywords"`         // SEO关键词
	Description      string             `gorm:"column:description;comment:SEO描述" json:"description"`    // SEO描述
	Sort             *int32             `gorm:"column:sort;default:255;comment:排序" json:"sort"`         // 排序
	ArticleCategory  *ArticleCategory   `gorm:"foreignKey:CategoryID" json:"article_category,omitempty"`
	ArticleContent   *ArticleContent    `gorm:"foreignKey:ArticleID" json:"article_content,omitempty"`
	ArticleStatistic *ArticleStatistic  `gorm:"foreignKey:ArticleID" json:"article_statistic,omitempty"`
	ArticleTags      []*ArticleTag      `gorm:"Many2many:article_tag_relates;foreignKey:ID;joinForeignKey:ArticleID;joinReferences:TagID;references:ID" json:"article_tags,omitempty"`
}

// TableName Article's table name
func (*Article) TableName() string {
	return TableNameArticle
}
