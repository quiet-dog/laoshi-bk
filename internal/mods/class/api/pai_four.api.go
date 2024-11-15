package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// Pk日志
type PaiFour struct {
	PaiFourBIZ *biz.PaiFour
}

// @Tags PaiFourAPI
// @Security ApiKeyAuth
// @Summary Query pai four list
// @Success 200 {object} util.ResponseResult{data=[]schema.PaiFour}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-fours [get]
func (a *PaiFour) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PaiFourQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	var list []*schema.PaiFour
	a.PaiFourBIZ.Trans.DB.Find(&list)
	if len(list) == 0 {
		var userList []*schema.Employ
		a.PaiFourBIZ.Trans.DB.Find(&userList)
		for _, item := range userList {
			if !item.IsTeacher {
				row := schema.PaiFour{}
				row.ID = util.NewXID()
				row.EmployId = item.ID
				row.Score = 0
				a.PaiFourBIZ.Trans.DB.Create(&row)
			}
		}
	}

	result, err := a.PaiFourBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags PaiFourAPI
// @Security ApiKeyAuth
// @Summary Get pai four record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.PaiFour}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-fours/{id} [get]
func (a *PaiFour) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.PaiFourBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags PaiFourAPI
// @Security ApiKeyAuth
// @Summary Create pai four record
// @Param body body schema.PaiFourForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.PaiFour}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-fours [post]
func (a *PaiFour) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PaiFourForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PaiFourBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags PaiFourAPI
// @Security ApiKeyAuth
// @Summary Update pai four record by ID
// @Param id path string true "unique id"
// @Param body body schema.PaiFourForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-fours/{id} [put]
func (a *PaiFour) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PaiFourForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.PaiFourBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags PaiFourAPI
// @Security ApiKeyAuth
// @Summary Delete pai four record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-fours/{id} [delete]
func (a *PaiFour) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.PaiFourBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
