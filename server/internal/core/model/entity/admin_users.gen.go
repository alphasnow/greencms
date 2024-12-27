// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"server/internal/core/model/accessor"

	"gorm.io/gorm"
)

const TableNameAdminUser = "admin_users"

// AdminUser 管理员
type AdminUser struct {
	ID        uint               `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time          `gorm:"column:created_at;<-:create" json:"created_at"`
	UpdatedAt time.Time          `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt     `gorm:"column:deleted_at" json:"-"`
	Username  string             `gorm:"column:username" json:"username"`
	Password  string             `gorm:"column:password" json:"-"`
	Realname  string             `gorm:"column:realname" json:"realname"`
	AvatarURL accessor.StaticUrl `gorm:"column:avatar_url" json:"avatar_url"`
	Access    string             `gorm:"column:access" json:"access"`
}

// TableName AdminUser's table name
func (*AdminUser) TableName() string {
	return TableNameAdminUser
}
