package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// Pk日志
type PaiThree struct {
	PaiThreeBIZ *biz.PaiThree
}

// @Tags PaiThreeAPI
// @Security ApiKeyAuth
// @Summary Query pai three list
// @Success 200 {object} util.ResponseResult{data=[]schema.PaiThree}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-threes [get]
func (a *PaiThree) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PaiThreeQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	var list []*schema.PaiThree
	a.PaiThreeBIZ.Trans.DB.Find(&list)
	if len(list) == 0 {
		var userList []*schema.Employ
		a.PaiThreeBIZ.Trans.DB.Find(&userList)
		for _, item := range userList {
			if !item.IsTeacher {
				row := schema.PaiThree{}
				row.ID = util.NewXID()
				row.EmployId = item.ID
				row.Score = 0
				a.PaiThreeBIZ.Trans.DB.Create(&row)
			}
		}
	}

	result, err := a.PaiThreeBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags PaiThreeAPI
// @Security ApiKeyAuth
// @Summary Get pai three record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.PaiThree}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-threes/{id} [get]
func (a *PaiThree) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.PaiThreeBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags PaiThreeAPI
// @Security ApiKeyAuth
// @Summary Create pai three record
// @Param body body schema.PaiThreeForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.PaiThree}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-threes [post]
func (a *PaiThree) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PaiThreeForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PaiThreeBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags PaiThreeAPI
// @Security ApiKeyAuth
// @Summary Update pai three record by ID
// @Param id path string true "unique id"
// @Param body body schema.PaiThreeForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-threes/{id} [put]
func (a *PaiThree) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.PaiThreeForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.PaiThreeBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags PaiThreeAPI
// @Security ApiKeyAuth
// @Summary Delete pai three record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/pai-threes/{id} [delete]
func (a *PaiThree) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.PaiThreeBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
