package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get class storage instance
func GetClassDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Class))
}

// 课件
type Class struct {
	DB *gorm.DB
}

// Query classes from the database based on the provided parameters and options.
func (a *Class) Query(ctx context.Context, params schema.ClassQueryParam, opts ...schema.ClassQueryOptions) (*schema.ClassQueryResult, error) {
	var opt schema.ClassQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetClassDB(ctx, a.DB)
	if v := params.Name; len(v) > 0 {
		db = db.Where("name = ?", v)
	}
	if v := params.Path; len(v) > 0 {
		db = db.Where("path = ?", v)
	}
	if v := params.Pid; len(v) > 0 {
		db = db.Where("pid = ?", v)
	}

	var list schema.Classes
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.ClassQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified class from the database.
func (a *Class) Get(ctx context.Context, id string, opts ...schema.ClassQueryOptions) (*schema.Class, error) {
	var opt schema.ClassQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Class)
	ok, err := util.FindOne(ctx, GetClassDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified class exists in the database.
func (a *Class) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetClassDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new class.
func (a *Class) Create(ctx context.Context, item *schema.Class) error {
	result := GetClassDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified class in the database.
func (a *Class) Update(ctx context.Context, item *schema.Class) error {
	result := GetClassDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified class from the database.
func (a *Class) Delete(ctx context.Context, id string) error {
	result := GetClassDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Class))
	return errors.WithStack(result.Error)
}
