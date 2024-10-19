package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xxx/testapp/internal/mods/class/biz"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/util"
)

// 文件
type File struct {
	FileBIZ *biz.File
}

// @Tags FileAPI
// @Security ApiKeyAuth
// @Summary Query file list
// @Success 200 {object} util.ResponseResult{data=[]schema.File}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/files [get]
func (a *File) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.FileQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.FileBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags FileAPI
// @Security ApiKeyAuth
// @Summary Get file record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.File}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/files/{id} [get]
func (a *File) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.FileBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags FileAPI
// @Security ApiKeyAuth
// @Summary Create file record
// @Param body body schema.FileForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.File}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/files [post]
func (a *File) Create(c *gin.Context) {

	// 获取上传文件
	file, err := c.FormFile("file")
	if err != nil {
		util.ResError(c, err)
		return
	}

	// 保存文件
	filePath := "upload/" + uuid.NewString() + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		util.ResError(c, err)
		return
	}

	ctx := c.Request.Context()
	item := new(schema.FileForm)
	item.OrignName = file.Filename
	item.Path = filePath

	result, err := a.FileBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags FileAPI
// @Security ApiKeyAuth
// @Summary Update file record by ID
// @Param id path string true "unique id"
// @Param body body schema.FileForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/files/{id} [put]
func (a *File) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.FileForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.FileBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags FileAPI
// @Security ApiKeyAuth
// @Summary Delete file record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/class/files/{id} [delete]
func (a *File) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.FileBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
