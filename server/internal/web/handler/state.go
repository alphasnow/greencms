// @author AlphaSnow

package handler

import (
	"encoding/json"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/gin-gonic/gin"
	"server/internal/core/constant"
	"server/internal/core/model/dao"
	"server/internal/web/ecode"
	"server/internal/web/schema"
	"server/pkg/g"
	"time"
)

type StateHandler struct{}

// Data
// @Summary 全局初始数据
// @Description 获取初始数据,如标题/统计/LOGO等
// @Tags Web
// @Success 200	{object} schema.StateDataResp "ok"
// @Router /api/web/state [get]
func (h *StateHandler) Data(c *gin.Context) {
	resp, err := h.getCacheState(c)
	if err != nil {
		schema.Error(c, ecode.InternalServerError, err)
		return
	}

	schema.Success(c, resp)
}

func (h *StateHandler) getDatabaseState(c *gin.Context) (*schema.StateDataResp, error) {
	metas, err := dao.WebMeta.WithContext(c).Where(dao.WebMeta.MetaGroup.In(constant.WebMetaDefaultGroup)).Find()
	if err != nil {
		return nil, err
	}
	categories, err := dao.ArticleCategory.WithContext(c).Order(dao.ArticleCategory.Sort.Asc(), dao.ArticleCategory.ID.Desc()).Limit(6).Find()
	if err != nil {
		return nil, err
	}

	metaMap := map[string]string{}
	for _, meta := range metas {
		metaMap[meta.MetaKey] = meta.MetaValue
	}

	resp := &schema.StateDataResp{Metas: metaMap, Categories: categories}
	return resp, nil
}

func (h *StateHandler) getCacheState(c *gin.Context) (*schema.StateDataResp, error) {
	resJson, err := g.Cache().Get(c, constant.WebGlobalStateCacheKey)
	if err == nil && resJson != "" {
		resp := &schema.StateDataResp{}
		if err = json.Unmarshal([]byte(resJson), resp); err != nil {
			return nil, err
		}
		return resp, nil
	}

	resp, err := h.getDatabaseState(c)
	if err != nil {
		return nil, err
	}

	respJsonByte, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	resJson = string(respJsonByte)
	if err = g.Cache().Set(c, constant.WebGlobalStateCacheKey, resJson, store.WithExpiration(1*time.Hour)); err != nil {
		return nil, err
	}
	return resp, nil
}
