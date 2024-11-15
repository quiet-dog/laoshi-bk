package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/paione/biz"
	"github.com/xxx/testapp/internal/mods/paione/schema"
	"github.com/xxx/testapp/pkg/util"
)

// Pk日志
type PaiOne struct {
	PaiOneBIZ *biz.PaiOne
}

// @Tags PaiOneAPI
// @Security ApiKeyAuth
// @Summary Query pai one list
// @Success 200 {object} util.ResponseResult{data=[]schema.PaiOne}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/paione/pai-ones [get]
func (a *PaiOne) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PaiOneQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PaiOneBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags PaiOneAPI
// @Security ApiKeyAuth
// @Summary Get pai one record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.PaiOne}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/paione/pai-ones/{id} [get]
func (a *PaiOne) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.PaiOneBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags PaiOneAPI
// @Security ApiKeyAuth
// @Summary Create pai one record
// @Param body body schema.PaiOneForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.PaiOne}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/paione/pai-ones [post]
func (a *PaiOne) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PaiOneForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PaiOneBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags PaiOneAPI
// @Security ApiKeyAuth
// @Summary Update pai one record by ID
// @Param id path string true "unique id"
// @Param body body schema.PaiOneForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/paione/pai-ones/{id} [put]
func (a *PaiOne) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PaiOneForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.PaiOneBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags PaiOneAPI
// @Security ApiKeyAuth
// @Summary Delete pai one record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/paione/pai-ones/{id} [delete]
func (a *PaiOne) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.PaiOneBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
