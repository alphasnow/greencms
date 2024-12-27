// @author AlphaSnow

package handler

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func Test_WebMeta_Index(t *testing.T) {
	//u := "/api/admin/article-category/index?current=1&pageSize=10&sort%5Bupdated_at%5D=ascend"
	//up, _ := url.Parse(u)
	//upr, _ := url.ParseQuery(up.RawQuery)
	//fmt.Println(upr)

	resp := getQuery(t, "/api/admin/web-meta/index", url.Values{
		"current":  []string{"2"},
		"pageSize": []string{"5"},
		//"title":            []string{"标题"},
		//"sort[updated_at]": []string{"ascend"},
	})

	assert.Equal(t, resp.Success, true)
}
