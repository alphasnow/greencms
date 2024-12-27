// @author AlphaSnow

package middleware

import (
	"github.com/casbin/casbin/v2"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/gin-gonic/gin"
	"log"
	"regexp"
	"server/internal/admin/schema"
	"server/pkg/g"
)

// UserAccessMiddleware
// 参考 https://github.com/gin-contrib/authz/blob/master/authz.go
type UserAccessMiddleware struct {
	enforcer *casbin.Enforcer
	regexp   *regexp.Regexp
}

func NewUserAccess() gin.HandlerFunc {
	m := new(UserAccessMiddleware)
	a := fileadapter.NewAdapter(g.Path("storage/app/casbin/admin_policy.csv"))
	e, err := casbin.NewEnforcer(g.Path("storage/app/casbin/admin_model.conf"), a)
	if err != nil {
		log.Fatal(err)
	}
	m.enforcer = e
	m.regexp = regexp.MustCompile(`/api/admin/([a-zA-Z\-_]+)/`)
	return m.Handle
}

func (m *UserAccessMiddleware) Handle(c *gin.Context) {
	res := m.regexp.FindStringSubmatch(c.Request.URL.Path)
	if res != nil && len(res) == 2 {
		user := GetUser(c) // access=editor
		obj := res[1]      // article
		ok, err := m.enforcer.Enforce(user.Access, obj)
		if err != nil {
			panic(err)
		}
		if ok == false {
			schema.ErrorJson(c, "权限受限")
			return
		}
	}
	c.Next()
}
