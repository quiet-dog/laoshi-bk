package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// 课件
type Class struct {
	ClassBIZ *biz.Class
}

// @Tags ClassAPI
// @Security ApiKeyAuth
// @Summary Query class list
// @Success 200 {object} util.ResponseResult{data=[]schema.Class}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/classes [get]
func (a *Class) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ClassQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ClassBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags ClassAPI
// @Security ApiKeyAuth
// @Summary Get class record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Class}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/classes/{id} [get]
func (a *Class) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ClassBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags ClassAPI
// @Security ApiKeyAuth
// @Summary Create class record
// @Param body body schema.ClassForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Class}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/classes [post]
func (a *Class) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ClassForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ClassBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags ClassAPI
// @Security ApiKeyAuth
// @Summary Update class record by ID
// @Param id path string true "unique id"
// @Param body body schema.ClassForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/classes/{id} [put]
func (a *Class) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ClassForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.ClassBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags ClassAPI
// @Security ApiKeyAuth
// @Summary Delete class record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/classes/{id} [delete]
func (a *Class) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ClassBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

func (a *Class) Tree(c *gin.Context) {
	ctx := c.Request.Context()
	list := a.ClassBIZ.Tree(ctx)

	util.ResSuccess(c, list)
}
