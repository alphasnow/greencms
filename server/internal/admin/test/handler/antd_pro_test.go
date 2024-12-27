// @author AlphaSnow

package handler

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestLoginAccount(t *testing.T) {
	resp := postJson(t, "/api/admin/login/account", map[string]string{
		"username": "admin",
		"password": "admin",
	})

	assert.Equal(t, resp.Success, true)
}

func TestCurrentUser(t *testing.T) {
	resp := getQuery(t, "/api/admin/currentUser", url.Values{})

	assert.Equal(t, resp.Success, true)
}
