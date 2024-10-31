package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get active storage instance
func GetActiveDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Active))
}

// 活动
type Active struct {
	DB *gorm.DB
}

// Query actives from the database based on the provided parameters and options.
func (a *Active) Query(ctx context.Context, params schema.ActiveQueryParam, opts ...schema.ActiveQueryOptions) (*schema.ActiveQueryResult, error) {
	var opt schema.ActiveQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetActiveDB(ctx, a.DB)

	if v := params.Type; v != "" {
		db = db.Where("type = ?", v)
	}

	// if v := params.Signid; len(v) > 0 {
	// 	db = db.Where("signid = ?", v)
	// }
	// if v := params.IsEnd; v != nil {
	// 	db = db.Where("is_end = ?", v)
	// }
	// if v := params.IsStart; v != nil {
	// 	db = db.Where("is_start = ?", v)
	// }
	var list schema.Actives
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.ActiveQueryResult{
		PageResult: pageResult,
		Data:       list,
	}

	return queryResult, nil
}

// Get the specified active from the database.
func (a *Active) Get(ctx context.Context, id string, opts ...schema.ActiveQueryOptions) (*schema.Active, error) {
	var opt schema.ActiveQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Active)
	ok, err := util.FindOne(ctx, GetActiveDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified active exists in the database.
func (a *Active) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetActiveDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new active.
func (a *Active) Create(ctx context.Context, item *schema.Active) error {
	result := GetActiveDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified active in the database.
func (a *Active) Update(ctx context.Context, item *schema.Active) error {
	result := GetActiveDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified active from the database.
func (a *Active) Delete(ctx context.Context, id string) error {
	result := GetActiveDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Active))
	return errors.WithStack(result.Error)
}
