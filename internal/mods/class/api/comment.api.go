package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// 讨论
type Comment struct {
	CommentBIZ *biz.Comment
}

// @Tags CommentAPI
// @Security ApiKeyAuth
// @Summary Query comment list
// @Success 200 {object} util.ResponseResult{data=[]schema.Comment}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/comments [get]
func (a *Comment) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.CommentQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.CommentBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags CommentAPI
// @Security ApiKeyAuth
// @Summary Get comment record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Comment}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/comments/{id} [get]
func (a *Comment) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.CommentBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags CommentAPI
// @Security ApiKeyAuth
// @Summary Create comment record
// @Param body body schema.CommentForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Comment}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/comments [post]
func (a *Comment) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.CommentForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.CommentBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags CommentAPI
// @Security ApiKeyAuth
// @Summary Update comment record by ID
// @Param id path string true "unique id"
// @Param body body schema.CommentForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/comments/{id} [put]
func (a *Comment) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.CommentForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.CommentBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags CommentAPI
// @Security ApiKeyAuth
// @Summary Delete comment record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/comments/{id} [delete]
func (a *Comment) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.CommentBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
