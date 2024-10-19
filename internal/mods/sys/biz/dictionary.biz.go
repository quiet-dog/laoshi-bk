package biz

import (
	"context"
	"time"

	"github.com/xxx/testapp/internal/mods/sys/dal"
	"github.com/xxx/testapp/internal/mods/sys/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
)

// Dictionaries management
type Dictionary struct {
	Trans         *util.Trans
	DictionaryDAL *dal.Dictionary
}

// Query dictionaries from the data access object based on the provided parameters and options.
func (a *Dictionary) Query(ctx context.Context, params schema.DictionaryQueryParam) (*schema.DictionaryQueryResult, error) {
	params.Pagination = true

	result, err := a.DictionaryDAL.Query(ctx, params, schema.DictionaryQueryOptions{
		QueryOptions: util.QueryOptions{
			OrderFields: []util.OrderByParam{
				{Field: "created_at", Direction: util.DESC},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Get the specified dictionary from the data access object.
func (a *Dictionary) Get(ctx context.Context, id string) (*schema.Dictionary, error) {
	dictionary, err := a.DictionaryDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if dictionary == nil {
		return nil, errors.NotFound("", "Dictionary not found")
	}
	return dictionary, nil
}

// Create a new dictionary in the data access object.
func (a *Dictionary) Create(ctx context.Context, formItem *schema.DictionaryForm) (*schema.Dictionary, error) {
	dictionary := &schema.Dictionary{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(dictionary); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.DictionaryDAL.Create(ctx, dictionary); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dictionary, nil
}

// Update the specified dictionary in the data access object.
func (a *Dictionary) Update(ctx context.Context, id string, formItem *schema.DictionaryForm) error {
	dictionary, err := a.DictionaryDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if dictionary == nil {
		return errors.NotFound("", "Dictionary not found")
	}

	if err := formItem.FillTo(dictionary); err != nil {
		return err
	}
	dictionary.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.DictionaryDAL.Update(ctx, dictionary); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified dictionary from the data access object.
func (a *Dictionary) Delete(ctx context.Context, id string) error {
	exists, err := a.DictionaryDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Dictionary not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.DictionaryDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
