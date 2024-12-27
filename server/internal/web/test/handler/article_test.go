// @author AlphaSnow

package handler

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func Test_ArticleHandler_List(t *testing.T) {
	resp := getQuery(t, "/api/web/article", url.Values{
		//"tag_id":      []string{"1"},
		//"category_id": []string{"2"},
		"keywords": []string{"微信"},
		"page":     []string{"20"},
	})

	assert.Equal(t, resp, true)
}

func Test_ArticleHandler_Statistic(t *testing.T) {
	resp := getQuery(t, "/api/web/statistic/views/1", url.Values{})

	assert.Equal(t, resp, true)
}
