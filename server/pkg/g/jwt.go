package g

import (
	"server/pkg/xjwt"
)

const AdminUser = "admin_users"

func JWT(roles ...string) (t *xjwt.UserToken) {
	var role string
	if len(roles) == 0 {
		role = "user"
	} else {
		role = roles[0]
	}

	insKey := "jwt." + role
	if globalContainer.Has(insKey) {
		t = globalContainer.Get(insKey).(*xjwt.UserToken)
	} else {
		t = xjwt.NewUserToken(
			Config().GetString("auth."+role+".secret"),
			Config().GetInt("auth."+role+".expire"),
			role,
		)
		globalContainer.Set(insKey, t)
	}
	return t

}
