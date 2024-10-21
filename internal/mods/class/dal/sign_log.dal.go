package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get sign log storage instance
func GetSignLogDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.SignLog))
}

// 签到日志
type SignLog struct {
	DB *gorm.DB
}

// Query sign logs from the database based on the provided parameters and options.
func (a *SignLog) Query(ctx context.Context, params schema.SignLogQueryParam, opts ...schema.SignLogQueryOptions) (*schema.SignLogQueryResult, error) {
	var opt schema.SignLogQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetSignLogDB(ctx, a.DB)

	var list schema.SignLogs
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.SignLogQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified sign log from the database.
func (a *SignLog) Get(ctx context.Context, id string, opts ...schema.SignLogQueryOptions) (*schema.SignLog, error) {
	var opt schema.SignLogQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.SignLog)
	ok, err := util.FindOne(ctx, GetSignLogDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified sign log exists in the database.
func (a *SignLog) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetSignLogDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new sign log.
func (a *SignLog) Create(ctx context.Context, item *schema.SignLog) error {
	result := GetSignLogDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified sign log in the database.
func (a *SignLog) Update(ctx context.Context, item *schema.SignLog) error {
	result := GetSignLogDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified sign log from the database.
func (a *SignLog) Delete(ctx context.Context, id string) error {
	result := GetSignLogDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.SignLog))
	return errors.WithStack(result.Error)
}
