// @author AlphaSnow

package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"server/internal/admin/schema"
	"server/internal/core/model/dao"
	"server/internal/core/model/entity"
	"server/pkg/g"
	"server/pkg/xjwt"
	"strings"
)

// "_gin-gonic/gin/bodybyteskey"
const userIDKey = "_server/middleware/user_id"
const userKey = "_server/middleware/user"

type AuthTokenMiddleware struct {
	skipPaths []string
	userToken *xjwt.UserToken
}

//func NewAuthTokenMiddleware(skipPaths []string) *AuthTokenMiddleware {
//	t := g.JWT(dao.AdminUser.TableName())
//	return &AuthTokenMiddleware{SkipPaths: skipPaths, UserToken: t}
//}

func NewAuthToken(skipPaths []string) gin.HandlerFunc {
	t := g.JWT(dao.AdminUser.TableName())
	m := &AuthTokenMiddleware{skipPaths: skipPaths, userToken: t}
	return m.Handle
}

func (m *AuthTokenMiddleware) Handle(c *gin.Context) {
	if len(m.skipPaths) == 0 || lo.Contains[string](m.skipPaths, c.Request.URL.Path) {
		c.Next()
		return
	}

	bearerToken, err := m.parseBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		schema.ErrorJson(c, err.Error())
		return
	}
	id, err := m.userToken.ParseID(bearerToken)
	if err != nil {
		//{
		//data: {
		//isLogin: false,
		//},
		//errorCode: '401',
		//	errorMessage: '请先登录！',
		//	success: true,
		//}
		schema.ErrorJson(c, "授权失败,请重新登陆", schema.WithErrorCode(401), schema.WithData(gin.H{"isLogin": false}))
		return
	}

	user, err := dao.AdminUser.WithContext(c).Where(dao.AdminUser.ID.Eq(id)).Take()
	if err != nil {
		schema.ErrorJson(c, "账号失效,请更换登录", schema.WithErrorCode(401), schema.WithData(gin.H{"isLogin": false}))
		return
	}

	c.Set(userKey, user)
	c.Set(userIDKey, id)
	c.Next()
}

func GetUserID(c *gin.Context) uint {
	return c.GetUint(userIDKey)
}
func GetUser(c *gin.Context) *entity.AdminUser {
	val, ok := c.Get(userKey)
	if !ok {
		return nil
	}
	return val.(*entity.AdminUser)
}

func (m *AuthTokenMiddleware) parseBearerToken(token string) (string, error) {
	if strings.HasPrefix(token, "Bearer") == false {
		return "", errors.New("未授权请求")
	}
	tk := strings.TrimPrefix(token, "Bearer ")
	if tk == "" {
		return "", errors.New("请重新登录")
	}
	return tk, nil
}
