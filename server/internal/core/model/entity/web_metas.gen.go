// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNameWebMeta = "web_metas"

// WebMeta 元数据
type WebMeta struct {
	ID        uint           `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	CreatedAt time.Time      `gorm:"column:created_at;<-:create" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
	MetaKey   string         `gorm:"column:meta_key" json:"meta_key"`
	MetaValue string         `gorm:"column:meta_value" json:"meta_value"`
	MetaGroup string         `gorm:"column:meta_group" json:"meta_group"`
	MetaName  string         `gorm:"column:meta_name" json:"meta_name"`
	Remark    string         `gorm:"column:remark" json:"remark"`
}

// TableName WebMeta's table name
func (*WebMeta) TableName() string {
	return TableNameWebMeta
}