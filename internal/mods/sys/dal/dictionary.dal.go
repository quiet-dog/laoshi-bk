package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/sys/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get dictionary storage instance
func GetDictionaryDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Dictionary))
}

// Dictionaries management
type Dictionary struct {
	DB *gorm.DB
}

// Query dictionaries from the database based on the provided parameters and options.
func (a *Dictionary) Query(ctx context.Context, params schema.DictionaryQueryParam, opts ...schema.DictionaryQueryOptions) (*schema.DictionaryQueryResult, error) {
	var opt schema.DictionaryQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetDictionaryDB(ctx, a.DB)

	var list schema.Dictionaries
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.DictionaryQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified dictionary from the database.
func (a *Dictionary) Get(ctx context.Context, id string, opts ...schema.DictionaryQueryOptions) (*schema.Dictionary, error) {
	var opt schema.DictionaryQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Dictionary)
	ok, err := util.FindOne(ctx, GetDictionaryDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified dictionary exists in the database.
func (a *Dictionary) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetDictionaryDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new dictionary.
func (a *Dictionary) Create(ctx context.Context, item *schema.Dictionary) error {
	result := GetDictionaryDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified dictionary in the database.
func (a *Dictionary) Update(ctx context.Context, item *schema.Dictionary) error {
	result := GetDictionaryDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified dictionary from the database.
func (a *Dictionary) Delete(ctx context.Context, id string) error {
	result := GetDictionaryDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Dictionary))
	return errors.WithStack(result.Error)
}
