// @author AlphaSnow

package handler

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func Test_GlobalHandler_Data(t *testing.T) {
	resp := getQuery(t, "/api/web/global", url.Values{})

	assert.Equal(t, resp, true)
}
