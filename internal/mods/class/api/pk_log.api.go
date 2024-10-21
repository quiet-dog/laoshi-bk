package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// Pk日志
type PkLog struct {
	PkLogBIZ *biz.PkLog
}

// @Tags PkLogAPI
// @Security ApiKeyAuth
// @Summary Query pk log list
// @Success 200 {object} util.ResponseResult{data=[]schema.PkLog}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pk-logs [get]
func (a *PkLog) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PkLogQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PkLogBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags PkLogAPI
// @Security ApiKeyAuth
// @Summary Get pk log record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.PkLog}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pk-logs/{id} [get]
func (a *PkLog) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.PkLogBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags PkLogAPI
// @Security ApiKeyAuth
// @Summary Create pk log record
// @Param body body schema.PkLogForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.PkLog}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pk-logs [post]
func (a *PkLog) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PkLogForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PkLogBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags PkLogAPI
// @Security ApiKeyAuth
// @Summary Update pk log record by ID
// @Param id path string true "unique id"
// @Param body body schema.PkLogForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pk-logs/{id} [put]
func (a *PkLog) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PkLogForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.PkLogBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags PkLogAPI
// @Security ApiKeyAuth
// @Summary Delete pk log record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pk-logs/{id} [delete]
func (a *PkLog) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.PkLogBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
