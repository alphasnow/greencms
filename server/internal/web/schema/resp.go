// @author AlphaSnow

package schema

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/web/ecode"
	"server/pkg/g"
)

type ErrorResp struct {
	ErrCode uint   `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

func Error(c *gin.Context, code ecode.ErrCode, errs ...error) {
	for _, err := range errs {
		g.Log().Error(err.Error())
	}
	c.Abort()
	c.JSON(http.StatusOK, ErrorResp{ErrCode: uint(code), ErrMsg: code.String()})
}

// Deprecated: inline 不支持 any
// https://github.com/golang/go/issues/6213#issuecomment-1500639389
type SuccessResp struct {
	ErrorResp `json:",inline"`
	Data      any `json:",inline"`
}

func Success(c *gin.Context, resp any) {
	c.JSON(http.StatusOK, resp)
}
