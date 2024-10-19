package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// 用户
type Employ struct {
	EmployBIZ *biz.Employ
}

// @Tags EmployAPI
// @Security ApiKeyAuth
// @Summary Query employ list
// @Success 200 {object} util.ResponseResult{data=[]schema.Employ}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/employs [get]
func (a *Employ) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.EmployQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.EmployBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags EmployAPI
// @Security ApiKeyAuth
// @Summary Get employ record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Employ}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/employs/{id} [get]
func (a *Employ) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.EmployBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags EmployAPI
// @Security ApiKeyAuth
// @Summary Create employ record
// @Param body body schema.EmployForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Employ}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/employs [post]
func (a *Employ) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.EmployForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.EmployBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags EmployAPI
// @Security ApiKeyAuth
// @Summary Update employ record by ID
// @Param id path string true "unique id"
// @Param body body schema.EmployForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/employs/{id} [put]
func (a *Employ) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.EmployForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.EmployBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags EmployAPI
// @Security ApiKeyAuth
// @Summary Delete employ record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/employs/{id} [delete]
func (a *Employ) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.EmployBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
