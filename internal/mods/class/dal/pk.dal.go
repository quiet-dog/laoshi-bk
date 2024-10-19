package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get pk storage instance
func GetPkDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Pk))
}

// 小组pk
type Pk struct {
	DB *gorm.DB
}

// Query pks from the database based on the provided parameters and options.
func (a *Pk) Query(ctx context.Context, params schema.PkQueryParam, opts ...schema.PkQueryOptions) (*schema.PkQueryResult, error) {
	var opt schema.PkQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetPkDB(ctx, a.DB)

	var list schema.Pks
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.PkQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified pk from the database.
func (a *Pk) Get(ctx context.Context, id string, opts ...schema.PkQueryOptions) (*schema.Pk, error) {
	var opt schema.PkQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Pk)
	ok, err := util.FindOne(ctx, GetPkDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified pk exists in the database.
func (a *Pk) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetPkDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new pk.
func (a *Pk) Create(ctx context.Context, item *schema.Pk) error {
	result := GetPkDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified pk in the database.
func (a *Pk) Update(ctx context.Context, item *schema.Pk) error {
	result := GetPkDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified pk from the database.
func (a *Pk) Delete(ctx context.Context, id string) error {
	result := GetPkDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Pk))
	return errors.WithStack(result.Error)
}
