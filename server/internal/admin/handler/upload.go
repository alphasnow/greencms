// @author AlphaSnow

package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"server/internal/admin/schema"
	"server/internal/admin/service"
	"server/pkg/g"
	"server/pkg/xjwt"
)

type UploadHandler struct {
	AuthService *service.AuthService
	UserToken   *xjwt.UserToken
}

func NewUploadHandler() *UploadHandler {
	h := new(UploadHandler)
	return h
}

func (h *UploadHandler) Routes(g gin.IRouter) {
	g.POST("/upload/form-file", h.FormFile)
}

func (h *UploadHandler) FormFile(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		schema.ErrorJson(c, "上传文件为空", schema.WithError(err))
		return
	}

	name, _ := service.NewFormFileService(f).GetHashName()

	p := c.DefaultPostForm("path", "article")
	// 获取本地路径
	storePath := g.Path(fmt.Sprintf("storage/upload/%s/%s", p, name))
	// 存储到本地
	if _, err = os.Stat(storePath); err != nil {
		err = c.SaveUploadedFile(f, storePath)
		if err != nil {
			schema.ErrorJson(c, "文件存储失败", schema.WithError(err))
			return
		}
	}

	// 获取访问链接
	filePath := fmt.Sprintf("upload/%s/%s", p, name)
	fileUrl := g.Url(filePath)
	schema.SuccessJson(c, gin.H{"url": fileUrl})
}
