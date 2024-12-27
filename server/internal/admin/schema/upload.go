// @author AlphaSnow

package schema

import "mime/multipart"

type UploadFormFile struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
	Path string                `form:"path" binding:"required"`
}
