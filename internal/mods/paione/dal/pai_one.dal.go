package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/paione/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get pai one storage instance
func GetPaiOneDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.PaiOne))
}

// Pk日志
type PaiOne struct {
	DB *gorm.DB
}

// Query pai ones from the database based on the provided parameters and options.
func (a *PaiOne) Query(ctx context.Context, params schema.PaiOneQueryParam, opts ...schema.PaiOneQueryOptions) (*schema.PaiOneQueryResult, error) {
	var opt schema.PaiOneQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetPaiOneDB(ctx, a.DB)

	var list schema.PaiOnes
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.PaiOneQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified pai one from the database.
func (a *PaiOne) Get(ctx context.Context, id string, opts ...schema.PaiOneQueryOptions) (*schema.PaiOne, error) {
	var opt schema.PaiOneQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.PaiOne)
	ok, err := util.FindOne(ctx, GetPaiOneDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified pai one exists in the database.
func (a *PaiOne) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetPaiOneDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new pai one.
func (a *PaiOne) Create(ctx context.Context, item *schema.PaiOne) error {
	result := GetPaiOneDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified pai one in the database.
func (a *PaiOne) Update(ctx context.Context, item *schema.PaiOne) error {
	result := GetPaiOneDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified pai one from the database.
func (a *PaiOne) Delete(ctx context.Context, id string) error {
	result := GetPaiOneDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.PaiOne))
	return errors.WithStack(result.Error)
}
