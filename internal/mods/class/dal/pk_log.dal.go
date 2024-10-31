package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get pk log storage instance
func GetPkLogDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.PkLog))
}

// Pk日志
type PkLog struct {
	DB *gorm.DB
}

// Query pk logs from the database based on the provided parameters and options.
func (a *PkLog) Query(ctx context.Context, params schema.PkLogQueryParam, opts ...schema.PkLogQueryOptions) (*schema.PkLogQueryResult, error) {
	var opt schema.PkLogQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetPkLogDB(ctx, a.DB)

	if v := params.ActiveId; v != "" {
		db = db.Where("active_id = ?", v)
	}

	var list schema.PkLogs
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.PkLogQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified pk log from the database.
func (a *PkLog) Get(ctx context.Context, id string, opts ...schema.PkLogQueryOptions) (*schema.PkLog, error) {
	var opt schema.PkLogQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.PkLog)
	ok, err := util.FindOne(ctx, GetPkLogDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified pk log exists in the database.
func (a *PkLog) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetPkLogDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new pk log.
func (a *PkLog) Create(ctx context.Context, item *schema.PkLog) error {
	result := GetPkLogDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified pk log in the database.
func (a *PkLog) Update(ctx context.Context, item *schema.PkLog) error {
	result := GetPkLogDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified pk log from the database.
func (a *PkLog) Delete(ctx context.Context, id string) error {
	result := GetPkLogDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.PkLog))
	return errors.WithStack(result.Error)
}
