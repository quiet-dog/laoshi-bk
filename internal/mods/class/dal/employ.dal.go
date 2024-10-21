package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get employ storage instance
func GetEmployDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Employ))
}

// 用户
type Employ struct {
	DB *gorm.DB
}

// Query employs from the database based on the provided parameters and options.
func (a *Employ) Query(ctx context.Context, params schema.EmployQueryParam, opts ...schema.EmployQueryOptions) (*schema.EmployQueryResult, error) {
	var opt schema.EmployQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetEmployDB(ctx, a.DB)
	if v := params.Name; len(v) > 0 {
		db = db.Where("name = ?", v)
	}
	if v := params.Username; len(v) > 0 {
		db = db.Where("username = ?", v)
	}
	if v := params.Password; len(v) > 0 {
		db = db.Where("password = ?", v)
	}

	if v := params.Committee; v > 0 {
		db = db.Where("committee = ?", v)
	}

	// if v := params.IsTeacher; v != nil {
	// 	db = db.Where("is_teacher = ?", v)
	// }
	// if v := params.IsCommittee; v != nil {
	// 	db = db.Where("is_committee = ?", v)
	// }
	// if v := params.Committee; v != 0 {
	// 	db = db.Where("committee = ?", v)
	// }

	var list schema.Employs
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.EmployQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified employ from the database.
func (a *Employ) Get(ctx context.Context, id string, opts ...schema.EmployQueryOptions) (*schema.Employ, error) {
	var opt schema.EmployQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Employ)
	ok, err := util.FindOne(ctx, GetEmployDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified employ exists in the database.
func (a *Employ) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetEmployDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new employ.
func (a *Employ) Create(ctx context.Context, item *schema.Employ) error {
	result := GetEmployDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified employ in the database.
func (a *Employ) Update(ctx context.Context, item *schema.Employ) error {
	result := GetEmployDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified employ from the database.
func (a *Employ) Delete(ctx context.Context, id string) error {
	result := GetEmployDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Employ))
	return errors.WithStack(result.Error)
}
