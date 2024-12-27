// @author AlphaSnow

package schema

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessJson(c *gin.Context, data any, opts ...AntdProRespOption) {
	c.JSON(http.StatusOK, SuccessResp(data, opts...))
}
func ErrorJson(c *gin.Context, errMsg string, opts ...AntdProRespOption) {
	c.Abort()
	c.JSON(http.StatusOK, ErrorResp(errMsg, opts...))
}

func PageListJson(c *gin.Context, data any, total int64, opts ...AntdProRespOption) {
	resp := SuccessResp(data, opts...)
	listResp := PageListResp{AntdProResp: *resp, Total: total}
	c.JSON(http.StatusOK, listResp)
}
