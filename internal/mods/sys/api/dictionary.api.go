package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/sys/biz"
	"github.com/xxx/testapp/internal/mods/sys/schema"
	"github.com/xxx/testapp/pkg/util"
)

// Dictionaries management
type Dictionary struct {
	DictionaryBIZ *biz.Dictionary
}

// @Tags DictionaryAPI
// @Security ApiKeyAuth
// @Summary Query dictionary list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Success 200 {object} util.ResponseResult{data=[]schema.Dictionary}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/sys/dictionaries [get]
func (a *Dictionary) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.DictionaryQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.DictionaryBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags DictionaryAPI
// @Security ApiKeyAuth
// @Summary Get dictionary record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Dictionary}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/sys/dictionaries/{id} [get]
func (a *Dictionary) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.DictionaryBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags DictionaryAPI
// @Security ApiKeyAuth
// @Summary Create dictionary record
// @Param body body schema.DictionaryForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Dictionary}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/sys/dictionaries [post]
func (a *Dictionary) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.DictionaryForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.DictionaryBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags DictionaryAPI
// @Security ApiKeyAuth
// @Summary Update dictionary record by ID
// @Param id path string true "unique id"
// @Param body body schema.DictionaryForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/sys/dictionaries/{id} [put]
func (a *Dictionary) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.DictionaryForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.DictionaryBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags DictionaryAPI
// @Security ApiKeyAuth
// @Summary Delete dictionary record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/sys/dictionaries/{id} [delete]
func (a *Dictionary) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DictionaryBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
