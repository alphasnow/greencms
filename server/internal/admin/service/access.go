// @author AlphaSnow

package service

import (
	"github.com/gin-gonic/gin"
	"server/internal/core/model/entity"
)

const EditorRole = "editor"
const AdminRole = "admin"
const AdminSuperID = 1

type AccessService struct{}

func (s *AccessService) GetRoleByID(c *gin.Context, id uint) string {
	access := EditorRole
	if id == AdminSuperID {
		access = AdminRole
	}
	return access
}

func (s *AccessService) GetRole(c *gin.Context, user *entity.AdminUser) string {
	access := EditorRole
	if user.ID == AdminSuperID {
		access = AdminRole
	}
	return access
}
