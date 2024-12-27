// @author AlphaSnow

package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"server/internal/admin/ecode"
	"strconv"
	"strings"
)

func GetParamID(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, ecode.WrongID
	}
	return uint(id), nil
}

var tablePresetFields = []string{
	"current",
	"pageSize",
	"filter",
	"sort",
}

func GetPageListConditions(c *gin.Context, tableName string) []gen.Condition {
	searches := make([]gen.Condition, 0)
	queries := c.Request.URL.Query()
	for k, v := range queries {
		if v[0] == "" {
			continue
		}
		if lo.ContainsBy[string](tablePresetFields, func(item string) bool {
			return strings.HasPrefix(k, item)
		}) {
			continue
		}

		w := field.NewString(tableName, k).Eq(strings.Trim(v[0], " "))
		searches = append(searches, w)
	}
	return searches
}

func GetPageListOrders(c *gin.Context, tableName string, defaultSort ...field.Expr) []field.Expr {
	searches := make([]field.Expr, 0)
	queries := c.QueryMap("sort")
	// {updated_at: 'ascend'} {updated_at: 'descend'}
	for k, v := range queries {
		if v == "ascend" {
			searches = append(searches, field.NewString(tableName, k).Asc())
		} else {
			searches = append(searches, field.NewString(tableName, k).Desc())
		}
	}
	if len(searches) == 0 {
		return defaultSort
	}
	return searches
}
