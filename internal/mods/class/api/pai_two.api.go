package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// Pk日志
type PaiTwo struct {
	PaiTwoBIZ *biz.PaiTwo
}

// @Tags PaiTwoAPI
// @Security ApiKeyAuth
// @Summary Query pai two list
// @Success 200 {object} util.ResponseResult{data=[]schema.PaiTwo}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-twos [get]
func (a *PaiTwo) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PaiTwoQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	var list []*schema.PaiTwo
	a.PaiTwoBIZ.Trans.DB.Find(&list)
	if len(list) == 0 {
		var userList []*schema.Employ
		a.PaiTwoBIZ.Trans.DB.Find(&userList)
		for _, item := range userList {
			if !item.IsTeacher {
				row := schema.PaiTwo{}
				row.ID = util.NewXID()
				row.EmployId = item.ID
				row.Score = 0
				a.PaiTwoBIZ.Trans.DB.Create(&row)
			}
		}
	}

	result, err := a.PaiTwoBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags PaiTwoAPI
// @Security ApiKeyAuth
// @Summary Get pai two record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.PaiTwo}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-twos/{id} [get]
func (a *PaiTwo) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.PaiTwoBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags PaiTwoAPI
// @Security ApiKeyAuth
// @Summary Create pai two record
// @Param body body schema.PaiTwoForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.PaiTwo}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-twos [post]
func (a *PaiTwo) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PaiTwoForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PaiTwoBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags PaiTwoAPI
// @Security ApiKeyAuth
// @Summary Update pai two record by ID
// @Param id path string true "unique id"
// @Param body body schema.PaiTwoForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-twos/{id} [put]
func (a *PaiTwo) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PaiTwoForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.PaiTwoBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags PaiTwoAPI
// @Security ApiKeyAuth
// @Summary Delete pai two record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-twos/{id} [delete]
func (a *PaiTwo) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.PaiTwoBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
