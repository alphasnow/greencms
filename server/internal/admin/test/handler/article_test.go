// @author AlphaSnow

package handler

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func Test_Article_Index(t *testing.T) {
	//u := "/api/admin/article-category/index?current=1&pageSize=10&sort%5Bupdated_at%5D=ascend"
	//up, _ := url.Parse(u)
	//upr, _ := url.ParseQuery(up.RawQuery)
	//fmt.Println(upr)

	resp := getQuery(t, "/api/admin/article/index", url.Values{
		"current":       []string{"1"},
		"pageSize":      []string{"5"},
		"title":         []string{"标题"},
		"category_name": []string{"生活心得"},
		//"sort[updated_at]": []string{"ascend"},
	})

	assert.Equal(t, resp.Success, true)
}

func Test_Article_Store(t *testing.T) {
	resp := postJson(t, "/api/admin/article/edit/2", map[string]interface{}{
		"title":           "123456",
		"article_content": map[string]string{"content": "article content"},
		"article_tags":    []map[string]int{{"id": 1}, {"id": 2}},
		"image_url":       "http://127.0.0.1:8080/upload/article-image/a51ba754977a2eff2cbfdc0db15ca6c4.png",
	})

	assert.Equal(t, resp.Success, true)
}
