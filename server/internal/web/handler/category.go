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

type CategoryHandler struct {
}

// Show
// @Summary 查看分类
// @Description 根据分类序号获取文章分类的详细数据
// @Tags Web
// @Param id path int true "分类ID"
// @Success 200	{object} entity.ArticleCategory "ok"
// @Router /api/web/category/{id} [get]
func (h *CategoryHandler) Show(c *gin.Context) {

	idP := c.Param("id")
	id, err := strconv.ParseUint(idP, 10, 0)
	if err != nil {
		schema.Error(c, ecode.BadRequest, err)
		return
	}

	data, err := dao.ArticleCategory.WithContext(c).
		Preload(field.Associations).
		Where(dao.ArticleCategory.ID.Eq(uint(id))).
		Take()
	if err != nil {
		schema.Error(c, ecode.ModelNotFound, err)
		return
	}

	schema.Success(c, data)
}
