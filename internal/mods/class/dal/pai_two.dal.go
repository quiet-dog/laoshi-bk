package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get pai two storage instance
func GetPaiTwoDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.PaiTwo))
}

// Pk日志
type PaiTwo struct {
	DB *gorm.DB
}

// Query pai twos from the database based on the provided parameters and options.
func (a *PaiTwo) Query(ctx context.Context, params schema.PaiTwoQueryParam, opts ...schema.PaiTwoQueryOptions) (*schema.PaiTwoQueryResult, error) {
	var opt schema.PaiTwoQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetPaiTwoDB(ctx, a.DB)
	db = db.Order("score desc")

	var list schema.PaiTwos
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.PaiTwoQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified pai two from the database.
func (a *PaiTwo) Get(ctx context.Context, id string, opts ...schema.PaiTwoQueryOptions) (*schema.PaiTwo, error) {
	var opt schema.PaiTwoQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.PaiTwo)
	ok, err := util.FindOne(ctx, GetPaiTwoDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified pai two exists in the database.
func (a *PaiTwo) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetPaiTwoDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new pai two.
func (a *PaiTwo) Create(ctx context.Context, item *schema.PaiTwo) error {
	result := GetPaiTwoDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified pai two in the database.
func (a *PaiTwo) Update(ctx context.Context, item *schema.PaiTwo) error {
	result := GetPaiTwoDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified pai two from the database.
func (a *PaiTwo) Delete(ctx context.Context, id string) error {
	result := GetPaiTwoDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.PaiTwo))
	return errors.WithStack(result.Error)
}
