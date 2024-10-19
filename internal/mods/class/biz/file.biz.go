package biz

import (
	"context"
	"time"

	"github.com/xxx/testapp/internal/mods/class/dal"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
)

// 文件
type File struct {
	Trans   *util.Trans
	FileDAL *dal.File
}

// Query files from the data access object based on the provided parameters and options.
func (a *File) Query(ctx context.Context, params schema.FileQueryParam) (*schema.FileQueryResult, error) {
	params.Pagination = false

	result, err := a.FileDAL.Query(ctx, params, schema.FileQueryOptions{
		QueryOptions: util.QueryOptions{
			OrderFields: []util.OrderByParam{
				{Field: "created_at", Direction: util.DESC},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Get the specified file from the data access object.
func (a *File) Get(ctx context.Context, id string) (*schema.File, error) {
	file, err := a.FileDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if file == nil {
		return nil, errors.NotFound("", "File not found")
	}
	return file, nil
}

// Create a new file in the data access object.
func (a *File) Create(ctx context.Context, formItem *schema.FileForm) (*schema.File, error) {
	file := &schema.File{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(file); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.FileDAL.Create(ctx, file); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Update the specified file in the data access object.
func (a *File) Update(ctx context.Context, id string, formItem *schema.FileForm) error {
	file, err := a.FileDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if file == nil {
		return errors.NotFound("", "File not found")
	}

	if err := formItem.FillTo(file); err != nil {
		return err
	}
	file.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.FileDAL.Update(ctx, file); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified file from the data access object.
func (a *File) Delete(ctx context.Context, id string) error {
	exists, err := a.FileDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "File not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.FileDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
