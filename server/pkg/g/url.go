package g

import (
	"fmt"
	"strings"
)

func Url(ps ...string) string {
	var b, p string
	if len(ps) == 2 {
		b = ps[1]
		p = ps[0]
	} else if len(ps) == 1 {
		b = "api"
		p = ps[0]
	} else {
		b = "api"
		p = ""
	}
	base := Config().GetString("server." + b + ".url")
	if p == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", base, strings.TrimLeft(p, "/"))
}
