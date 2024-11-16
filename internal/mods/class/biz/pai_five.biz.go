package biz

import (
	"context"
	"time"

	"github.com/xxx/testapp/internal/mods/class/dal"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
)

// Pk日志
type PaiFive struct {
	Trans      *util.Trans
	PaiFiveDAL *dal.PaiFive
}

// Query pai fives from the data access object based on the provided parameters and options.
func (a *PaiFive) Query(ctx context.Context, params schema.PaiFiveQueryParam) (*schema.PaiFiveQueryResult, error) {
	params.Pagination = false

	result, err := a.PaiFiveDAL.Query(ctx, params, schema.PaiFiveQueryOptions{
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

// Get the specified pai five from the data access object.
func (a *PaiFive) Get(ctx context.Context, id string) (*schema.PaiFive, error) {
	paiFive, err := a.PaiFiveDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if paiFive == nil {
		return nil, errors.NotFound("", "Pai five not found")
	}
	return paiFive, nil
}

// Create a new pai five in the data access object.
func (a *PaiFive) Create(ctx context.Context, formItem *schema.PaiFiveForm) (*schema.PaiFive, error) {
	paiFive := &schema.PaiFive{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(paiFive); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiFiveDAL.Create(ctx, paiFive); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paiFive, nil
}

// Update the specified pai five in the data access object.
func (a *PaiFive) Update(ctx context.Context, id string, formItem *schema.PaiFiveForm) error {
	paiFive, err := a.PaiFiveDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if paiFive == nil {
		return errors.NotFound("", "Pai five not found")
	}

	if err := formItem.FillTo(paiFive); err != nil {
		return err
	}
	paiFive.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiFiveDAL.Update(ctx, paiFive); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified pai five from the data access object.
func (a *PaiFive) Delete(ctx context.Context, id string) error {
	exists, err := a.PaiFiveDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Pai five not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiFiveDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
