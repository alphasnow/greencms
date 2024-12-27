// @author AlphaSnow

package handler

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func Test_HomeHandler_Data(t *testing.T) {
	resp := getQuery(t, "/api/web/home", url.Values{})

	assert.Equal(t, resp, true)
}
