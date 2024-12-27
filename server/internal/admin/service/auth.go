// @author AlphaSnow

package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"server/internal/admin/schema"
	"server/internal/core/model/dao"
	"server/internal/core/model/entity"
	"server/pkg/utils"
)

type AuthService struct{}

func (s *AuthService) Login(c *gin.Context, req *schema.LoginParams) (*entity.AdminUser, error) {
	user, err := dao.AdminUser.WithContext(c).Where(dao.AdminUser.Username.Eq(req.Username)).Take()
	if err != nil {
		return nil, errors.New("账号不存在")
	}

	if utils.PasswordVerify(req.Password, user.Password) == false {
		return nil, errors.New("密码错误")
	}

	return user, nil
}
