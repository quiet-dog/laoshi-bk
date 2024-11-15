package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get pai four storage instance
func GetPaiFourDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.PaiFour))
}

// Pk日志
type PaiFour struct {
	DB *gorm.DB
}

// Query pai fours from the database based on the provided parameters and options.
func (a *PaiFour) Query(ctx context.Context, params schema.PaiFourQueryParam, opts ...schema.PaiFourQueryOptions) (*schema.PaiFourQueryResult, error) {
	var opt schema.PaiFourQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetPaiFourDB(ctx, a.DB)

	// 按找score由大到小排序
	db = db.Order("score desc")

	var list schema.PaiFours
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.PaiFourQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified pai four from the database.
func (a *PaiFour) Get(ctx context.Context, id string, opts ...schema.PaiFourQueryOptions) (*schema.PaiFour, error) {
	var opt schema.PaiFourQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.PaiFour)
	ok, err := util.FindOne(ctx, GetPaiFourDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified pai four exists in the database.
func (a *PaiFour) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetPaiFourDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new pai four.
func (a *PaiFour) Create(ctx context.Context, item *schema.PaiFour) error {
	result := GetPaiFourDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified pai four in the database.
func (a *PaiFour) Update(ctx context.Context, item *schema.PaiFour) error {
	result := GetPaiFourDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified pai four from the database.
func (a *PaiFour) Delete(ctx context.Context, id string) error {
	result := GetPaiFourDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.PaiFour))
	return errors.WithStack(result.Error)
}
