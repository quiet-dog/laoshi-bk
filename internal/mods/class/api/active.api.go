package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// 活动
type Active struct {
	ActiveBIZ *biz.Active
}

// @Tags ActiveAPI
// @Security ApiKeyAuth
// @Summary Query active list
// @Success 200 {object} util.ResponseResult{data=[]schema.Active}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/actives [get]
func (a *Active) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ActiveQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ActiveBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags ActiveAPI
// @Security ApiKeyAuth
// @Summary Get active record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Active}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/actives/{id} [get]
func (a *Active) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ActiveBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags ActiveAPI
// @Security ApiKeyAuth
// @Summary Create active record
// @Param body body schema.ActiveForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Active}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/actives [post]
func (a *Active) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ActiveForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ActiveBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags ActiveAPI
// @Security ApiKeyAuth
// @Summary Update active record by ID
// @Param id path string true "unique id"
// @Param body body schema.ActiveForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/actives/{id} [put]
func (a *Active) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ActiveForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.ActiveBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags ActiveAPI
// @Security ApiKeyAuth
// @Summary Delete active record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/actives/{id} [delete]
func (a *Active) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ActiveBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
