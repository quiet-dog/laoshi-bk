package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get pai three storage instance
func GetPaiThreeDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.PaiThree))
}

// Pk日志
type PaiThree struct {
	DB *gorm.DB
}

// Query pai threes from the database based on the provided parameters and options.
func (a *PaiThree) Query(ctx context.Context, params schema.PaiThreeQueryParam, opts ...schema.PaiThreeQueryOptions) (*schema.PaiThreeQueryResult, error) {
	var opt schema.PaiThreeQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetPaiThreeDB(ctx, a.DB)
	db = db.Order("score desc")

	var list schema.PaiThrees
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.PaiThreeQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified pai three from the database.
func (a *PaiThree) Get(ctx context.Context, id string, opts ...schema.PaiThreeQueryOptions) (*schema.PaiThree, error) {
	var opt schema.PaiThreeQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.PaiThree)
	ok, err := util.FindOne(ctx, GetPaiThreeDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified pai three exists in the database.
func (a *PaiThree) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetPaiThreeDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new pai three.
func (a *PaiThree) Create(ctx context.Context, item *schema.PaiThree) error {
	result := GetPaiThreeDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified pai three in the database.
func (a *PaiThree) Update(ctx context.Context, item *schema.PaiThree) error {
	result := GetPaiThreeDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified pai three from the database.
func (a *PaiThree) Delete(ctx context.Context, id string) error {
	result := GetPaiThreeDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.PaiThree))
	return errors.WithStack(result.Error)
}
