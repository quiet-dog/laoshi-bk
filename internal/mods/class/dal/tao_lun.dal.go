package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get tao lun storage instance
func GetTaoLunDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.TaoLun))
}

// 讨论
type TaoLun struct {
	DB *gorm.DB
}

// Query tao luns from the database based on the provided parameters and options.
func (a *TaoLun) Query(ctx context.Context, params schema.TaoLunQueryParam, opts ...schema.TaoLunQueryOptions) (*schema.TaoLunQueryResult, error) {
	var opt schema.TaoLunQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetTaoLunDB(ctx, a.DB)

	var list schema.TaoLuns
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.TaoLunQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified tao lun from the database.
func (a *TaoLun) Get(ctx context.Context, id string, opts ...schema.TaoLunQueryOptions) (*schema.TaoLun, error) {
	var opt schema.TaoLunQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.TaoLun)
	ok, err := util.FindOne(ctx, GetTaoLunDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified tao lun exists in the database.
func (a *TaoLun) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetTaoLunDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new tao lun.
func (a *TaoLun) Create(ctx context.Context, item *schema.TaoLun) error {
	result := GetTaoLunDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified tao lun in the database.
func (a *TaoLun) Update(ctx context.Context, item *schema.TaoLun) error {
	result := GetTaoLunDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified tao lun from the database.
func (a *TaoLun) Delete(ctx context.Context, id string) error {
	result := GetTaoLunDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.TaoLun))
	return errors.WithStack(result.Error)
}
