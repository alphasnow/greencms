// @author AlphaSnow

package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gen/field"
	"server/internal/core/model/dao"
	"server/internal/web/ecode"
	"server/internal/web/schema"
	"strconv"
)

type TagHandler struct {
}

// Show
// @Summary 查看标签
// @Description 根据标签序号获取标签详细数据
// @Tags Web
// @Param id path int true "标签ID"
// @Success 200	{object} entity.ArticleTag "ok"
// @Router /api/web/tag/{id} [get]
func (h *TagHandler) Show(c *gin.Context) {

	idP := c.Param("id")
	id, err := strconv.ParseUint(idP, 10, 0)
	if err != nil {
		schema.Error(c, ecode.BadRequest, err)
		return
	}

	data, err := dao.ArticleTag.WithContext(c).
		Preload(field.Associations).
		Where(dao.ArticleTag.ID.Eq(uint(id))).
		Take()
	if err != nil {
		schema.Error(c, ecode.ModelNotFound, err)
		return
	}

	schema.Success(c, data)
}
