package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// 签到日志
type SignLog struct {
	SignLogBIZ *biz.SignLog
}

// @Tags SignLogAPI
// @Security ApiKeyAuth
// @Summary Query sign log list
// @Success 200 {object} util.ResponseResult{data=[]schema.SignLog}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/sign-logs [get]
func (a *SignLog) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.SignLogQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.SignLogBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags SignLogAPI
// @Security ApiKeyAuth
// @Summary Get sign log record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.SignLog}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/sign-logs/{id} [get]
func (a *SignLog) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SignLogBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags SignLogAPI
// @Security ApiKeyAuth
// @Summary Create sign log record
// @Param body body schema.SignLogForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.SignLog}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/sign-logs [post]
func (a *SignLog) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.SignLogForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.SignLogBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags SignLogAPI
// @Security ApiKeyAuth
// @Summary Update sign log record by ID
// @Param id path string true "unique id"
// @Param body body schema.SignLogForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/sign-logs/{id} [put]
func (a *SignLog) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.SignLogForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.SignLogBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags SignLogAPI
// @Security ApiKeyAuth
// @Summary Delete sign log record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/sign-logs/{id} [delete]
func (a *SignLog) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SignLogBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
