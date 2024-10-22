package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// 小组pk打分
type PkScore struct {
	PkScoreBIZ *biz.PkScore
}

// @Tags PkScoreAPI
// @Security ApiKeyAuth
// @Summary Query pk score list
// @Success 200 {object} util.ResponseResult{data=[]schema.PkScore}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pk-scores [get]
func (a *PkScore) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PkScoreQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PkScoreBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags PkScoreAPI
// @Security ApiKeyAuth
// @Summary Get pk score record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.PkScore}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pk-scores/{id} [get]
func (a *PkScore) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.PkScoreBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags PkScoreAPI
// @Security ApiKeyAuth
// @Summary Create pk score record
// @Param body body schema.PkScoreForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.PkScore}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pk-scores [post]
func (a *PkScore) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PkScoreForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PkScoreBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags PkScoreAPI
// @Security ApiKeyAuth
// @Summary Update pk score record by ID
// @Param id path string true "unique id"
// @Param body body schema.PkScoreForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pk-scores/{id} [put]
func (a *PkScore) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PkScoreForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.PkScoreBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags PkScoreAPI
// @Security ApiKeyAuth
// @Summary Delete pk score record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pk-scores/{id} [delete]
func (a *PkScore) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.PkScoreBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
