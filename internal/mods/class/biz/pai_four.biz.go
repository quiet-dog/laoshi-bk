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
type PaiFour struct {
	Trans      *util.Trans
	PaiFourDAL *dal.PaiFour
}

// Query pai fours from the data access object based on the provided parameters and options.
func (a *PaiFour) Query(ctx context.Context, params schema.PaiFourQueryParam) (*schema.PaiFourQueryResult, error) {
	params.Pagination = false

	result, err := a.PaiFourDAL.Query(ctx, params, schema.PaiFourQueryOptions{
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

// Get the specified pai four from the data access object.
func (a *PaiFour) Get(ctx context.Context, id string) (*schema.PaiFour, error) {
	paiFour, err := a.PaiFourDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if paiFour == nil {
		return nil, errors.NotFound("", "Pai four not found")
	}
	return paiFour, nil
}

// Create a new pai four in the data access object.
func (a *PaiFour) Create(ctx context.Context, formItem *schema.PaiFourForm) (*schema.PaiFour, error) {
	paiFour := &schema.PaiFour{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(paiFour); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiFourDAL.Create(ctx, paiFour); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paiFour, nil
}

// Update the specified pai four in the data access object.
func (a *PaiFour) Update(ctx context.Context, id string, formItem *schema.PaiFourForm) error {
	paiFour, err := a.PaiFourDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if paiFour == nil {
		return errors.NotFound("", "Pai four not found")
	}

	if err := formItem.FillTo(paiFour); err != nil {
		return err
	}
	paiFour.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiFourDAL.Update(ctx, paiFour); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified pai four from the data access object.
func (a *PaiFour) Delete(ctx context.Context, id string) error {
	exists, err := a.PaiFourDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Pai four not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiFourDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
