package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get pai five storage instance
func GetPaiFiveDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.PaiFive))
}

// Pk日志
type PaiFive struct {
	DB *gorm.DB
}

// Query pai fives from the database based on the provided parameters and options.
func (a *PaiFive) Query(ctx context.Context, params schema.PaiFiveQueryParam, opts ...schema.PaiFiveQueryOptions) (*schema.PaiFiveQueryResult, error) {
	var opt schema.PaiFiveQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetPaiFiveDB(ctx, a.DB)

	var list schema.PaiFives
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.PaiFiveQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified pai five from the database.
func (a *PaiFive) Get(ctx context.Context, id string, opts ...schema.PaiFiveQueryOptions) (*schema.PaiFive, error) {
	var opt schema.PaiFiveQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.PaiFive)
	ok, err := util.FindOne(ctx, GetPaiFiveDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified pai five exists in the database.
func (a *PaiFive) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetPaiFiveDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new pai five.
func (a *PaiFive) Create(ctx context.Context, item *schema.PaiFive) error {
	result := GetPaiFiveDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified pai five in the database.
func (a *PaiFive) Update(ctx context.Context, item *schema.PaiFive) error {
	result := GetPaiFiveDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified pai five from the database.
func (a *PaiFive) Delete(ctx context.Context, id string) error {
	result := GetPaiFiveDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.PaiFive))
	return errors.WithStack(result.Error)
}
