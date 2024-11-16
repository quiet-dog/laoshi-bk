package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// Pk日志
type PaiFive struct {
	PaiFiveBIZ *biz.PaiFive
}

// @Tags PaiFiveAPI
// @Security ApiKeyAuth
// @Summary Query pai five list
// @Success 200 {object} util.ResponseResult{data=[]schema.PaiFive}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-fives [get]
func (a *PaiFive) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PaiFiveQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	var list []*schema.PaiFive
	a.PaiFiveBIZ.Trans.DB.Find(&list)
	if len(list) == 0 {
		var userList []*schema.Employ
		a.PaiFiveBIZ.Trans.DB.Find(&userList)
		for _, item := range userList {
			if !item.IsTeacher {
				row := schema.PaiFive{}
				row.ID = util.NewXID()
				row.EmployId = item.ID
				row.Score = 0
				a.PaiFiveBIZ.Trans.DB.Create(&row)
			}
		}
	}

	result, err := a.PaiFiveBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags PaiFiveAPI
// @Security ApiKeyAuth
// @Summary Get pai five record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.PaiFive}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-fives/{id} [get]
func (a *PaiFive) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.PaiFiveBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags PaiFiveAPI
// @Security ApiKeyAuth
// @Summary Create pai five record
// @Param body body schema.PaiFiveForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.PaiFive}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-fives [post]
func (a *PaiFive) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PaiFiveForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PaiFiveBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags PaiFiveAPI
// @Security ApiKeyAuth
// @Summary Update pai five record by ID
// @Param id path string true "unique id"
// @Param body body schema.PaiFiveForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-fives/{id} [put]
func (a *PaiFive) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PaiFiveForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.PaiFiveBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags PaiFiveAPI
// @Security ApiKeyAuth
// @Summary Delete pai five record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-fives/{id} [delete]
func (a *PaiFive) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.PaiFiveBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
