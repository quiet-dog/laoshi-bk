package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// 签到
type Sign struct {
	SignBIZ *biz.Sign
}

// @Tags SignAPI
// @Security ApiKeyAuth
// @Summary Query sign list
// @Success 200 {object} util.ResponseResult{data=[]schema.Sign}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/signs [get]
func (a *Sign) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.SignQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.SignBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags SignAPI
// @Security ApiKeyAuth
// @Summary Get sign record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Sign}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/signs/{id} [get]
func (a *Sign) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SignBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags SignAPI
// @Security ApiKeyAuth
// @Summary Create sign record
// @Param body body schema.SignForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Sign}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/signs [post]
func (a *Sign) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.SignForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.SignBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}

	util.ResSuccess(c, result)
}

// @Tags SignAPI
// @Security ApiKeyAuth
// @Summary Update sign record by ID
// @Param id path string true "unique id"
// @Param body body schema.SignForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/signs/{id} [put]
func (a *Sign) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.SignForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.SignBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags SignAPI
// @Security ApiKeyAuth
// @Summary Delete sign record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/signs/{id} [delete]
func (a *Sign) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SignBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
