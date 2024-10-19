package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// 讨论
type TaoLun struct {
	TaoLunBIZ *biz.TaoLun
}

// @Tags TaoLunAPI
// @Security ApiKeyAuth
// @Summary Query tao lun list
// @Success 200 {object} util.ResponseResult{data=[]schema.TaoLun}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/tao-luns [get]
func (a *TaoLun) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.TaoLunQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.TaoLunBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags TaoLunAPI
// @Security ApiKeyAuth
// @Summary Get tao lun record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.TaoLun}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/tao-luns/{id} [get]
func (a *TaoLun) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.TaoLunBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags TaoLunAPI
// @Security ApiKeyAuth
// @Summary Create tao lun record
// @Param body body schema.TaoLunForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.TaoLun}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/tao-luns [post]
func (a *TaoLun) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.TaoLunForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.TaoLunBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags TaoLunAPI
// @Security ApiKeyAuth
// @Summary Update tao lun record by ID
// @Param id path string true "unique id"
// @Param body body schema.TaoLunForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/tao-luns/{id} [put]
func (a *TaoLun) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.TaoLunForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.TaoLunBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags TaoLunAPI
// @Security ApiKeyAuth
// @Summary Delete tao lun record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/tao-luns/{id} [delete]
func (a *TaoLun) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.TaoLunBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
