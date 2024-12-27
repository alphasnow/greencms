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
	"server/internal/admin/router"
	"server/internal/admin/schema"
	"server/internal/core/model/dao"
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

func postJson(t *testing.T, path string, params any) *schema.AntdProResp {
	tk, err := generateToken()
	if err != nil {
		t.Fatal(err)
	}
	paramsJson, _ := json.Marshal(params)
	body := bytes.NewBuffer(paramsJson)
	req, _ := http.NewRequest(http.MethodPost, path, body)
	req.Header.Set("Authorization", "Bearer "+tk)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp := httptest.NewRecorder()
	engine.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("POST %s failed: %d %s", path, resp.Code, resp.Body.String())
	}
	antdproresp := new(schema.AntdProResp)
	_ = json.Unmarshal(resp.Body.Bytes(), antdproresp)
	return antdproresp
}

// var adminId uint = 1 // root
// var adminId uint = 2 // admin
// var adminId uint = 3 // manager
var adminId uint = 4 // editor

func generateToken() (string, error) {

	return g.JWT(dao.AdminUser.TableName()).GenerateID(adminId)

}

func getQuery(t *testing.T, path string, params url.Values) *schema.AntdProResp {
	tk, err := generateToken()
	if err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest(http.MethodGet, path, nil)
	req.Header.Set("Authorization", "Bearer "+tk)
	req.Header.Set("Accept", "application/json")
	req.URL.RawQuery = params.Encode()
	resp := httptest.NewRecorder()
	engine.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("GET %s failed: %d %s", path, resp.Code, resp.Body.String())
	}
	antdproresp := new(schema.AntdProResp)
	_ = json.Unmarshal(resp.Body.Bytes(), antdproresp)
	return antdproresp
}
