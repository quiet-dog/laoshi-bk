package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get sign storage instance
func GetSignDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Sign))
}

// 签到
type Sign struct {
	DB *gorm.DB
}

// Query signs from the database based on the provided parameters and options.
func (a *Sign) Query(ctx context.Context, params schema.SignQueryParam, opts ...schema.SignQueryOptions) (*schema.SignQueryResult, error) {
	var opt schema.SignQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetSignDB(ctx, a.DB)
	if v := params.Type; len(v) > 0 {
		db = db.Where("type = ?", v)
	}

	var list schema.Signs
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.SignQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified sign from the database.
func (a *Sign) Get(ctx context.Context, id string, opts ...schema.SignQueryOptions) (*schema.Sign, error) {
	var opt schema.SignQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Sign)
	ok, err := util.FindOne(ctx, GetSignDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified sign exists in the database.
func (a *Sign) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetSignDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new sign.
func (a *Sign) Create(ctx context.Context, item *schema.Sign) error {
	result := GetSignDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified sign in the database.
func (a *Sign) Update(ctx context.Context, item *schema.Sign) error {
	result := GetSignDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified sign from the database.
func (a *Sign) Delete(ctx context.Context, id string) error {
	result := GetSignDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Sign))
	return errors.WithStack(result.Error)
}
