// @author AlphaSnow

package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"server/internal/core/model/dao"
	"server/internal/web/router"
	"server/pkg/g"
	"testing"
)

var engine *gin.Engine

func TestMain(m *testing.M) {
	_ = g.Config()
	dao.SetDefault(g.DB())

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	//r.Static("/upload", g.Path("storage/upload"))
	router.Register(r)
	//router2.Register(r)

	engine = r
	code := m.Run()
	os.Exit(code)
}

func postJson(t *testing.T, path string, params any) string {
	paramsJson, _ := json.Marshal(params)
	body := bytes.NewBuffer(paramsJson)
	req, _ := http.NewRequest(http.MethodPost, path, body)
	req.Header.Set("Accept", "application/json")
	resp := httptest.NewRecorder()
	engine.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("POST %s failed: %d %s", path, resp.Code, resp.Body.String())
	}
	return resp.Body.String()
}

func getQuery(t *testing.T, path string, params url.Values) string {
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	req.Header.Set("Accept", "application/json")
	req.URL.RawQuery = params.Encode()
	resp := httptest.NewRecorder()
	engine.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("GET %s failed: %d %s", path, resp.Code, resp.Body.String())
	}
	return resp.Body.String()
}
