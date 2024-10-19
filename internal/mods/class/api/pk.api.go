package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// 小组pk
type Pk struct {
	PkBIZ *biz.Pk
}

// @Tags PkAPI
// @Security ApiKeyAuth
// @Summary Query pk list
// @Success 200 {object} util.ResponseResult{data=[]schema.Pk}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pks [get]
func (a *Pk) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PkQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PkBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags PkAPI
// @Security ApiKeyAuth
// @Summary Get pk record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Pk}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pks/{id} [get]
func (a *Pk) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.PkBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags PkAPI
// @Security ApiKeyAuth
// @Summary Create pk record
// @Param body body schema.PkForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Pk}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pks [post]
func (a *Pk) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PkForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PkBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags PkAPI
// @Security ApiKeyAuth
// @Summary Update pk record by ID
// @Param id path string true "unique id"
// @Param body body schema.PkForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pks/{id} [put]
func (a *Pk) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PkForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.PkBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags PkAPI
// @Security ApiKeyAuth
// @Summary Delete pk record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pks/{id} [delete]
func (a *Pk) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.PkBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
