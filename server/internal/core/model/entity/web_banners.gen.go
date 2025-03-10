// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"server/internal/core/model/accessor"

	"gorm.io/gorm"
)

const TableNameWebBanner = "web_banners"

// WebBanner 轮播图
type WebBanner struct {
	ID          uint               `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time          `gorm:"column:created_at;<-:create" json:"created_at"`
	UpdatedAt   time.Time          `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt     `gorm:"column:deleted_at" json:"-"`
	ImageURL    accessor.StaticUrl `gorm:"column:image_url" json:"image_url"`
	RedirectURL string             `gorm:"column:redirect_url" json:"redirect_url"`
	BannerGroup string             `gorm:"column:banner_group" json:"banner_group"`
	Sort        *int32             `gorm:"column:sort;default:255" json:"sort"`
	Remark      string             `gorm:"column:remark" json:"remark"`
	Title       string             `gorm:"column:title" json:"title"`
	Description string             `gorm:"column:description" json:"description"`
}

// TableName WebBanner's table name
func (*WebBanner) TableName() string {
	return TableNameWebBanner
}
