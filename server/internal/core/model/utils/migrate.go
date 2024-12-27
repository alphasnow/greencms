package utils

import (
	"gorm.io/gorm"
	"server/internal/core/model/entity"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		entity.AdminUser{},
		entity.Article{},
		entity.ArticleContent{},
		entity.ArticleCategory{},
		entity.ArticleTag{},
		entity.ArticleTagRelate{},
		entity.ArticleStatistic{},
		entity.WebMeta{},
		entity.WebBanner{},
	)
	return err
}
