package biz

import (
	"context"
	"time"

	"github.com/xxx/testapp/internal/mods/paione/dal"
	"github.com/xxx/testapp/internal/mods/paione/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
)

// Pk日志
type PaiOne struct {
	Trans     *util.Trans
	PaiOneDAL *dal.PaiOne
}

// Query pai ones from the data access object based on the provided parameters and options.
func (a *PaiOne) Query(ctx context.Context, params schema.PaiOneQueryParam) (*schema.PaiOneQueryResult, error) {
	params.Pagination = false

	result, err := a.PaiOneDAL.Query(ctx, params, schema.PaiOneQueryOptions{
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

// Get the specified pai one from the data access object.
func (a *PaiOne) Get(ctx context.Context, id string) (*schema.PaiOne, error) {
	paiOne, err := a.PaiOneDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if paiOne == nil {
		return nil, errors.NotFound("", "Pai one not found")
	}
	return paiOne, nil
}

// Create a new pai one in the data access object.
func (a *PaiOne) Create(ctx context.Context, formItem *schema.PaiOneForm) (*schema.PaiOne, error) {
	paiOne := &schema.PaiOne{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(paiOne); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiOneDAL.Create(ctx, paiOne); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paiOne, nil
}

// Update the specified pai one in the data access object.
func (a *PaiOne) Update(ctx context.Context, id string, formItem *schema.PaiOneForm) error {
	paiOne, err := a.PaiOneDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if paiOne == nil {
		return errors.NotFound("", "Pai one not found")
	}

	if err := formItem.FillTo(paiOne); err != nil {
		return err
	}
	paiOne.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiOneDAL.Update(ctx, paiOne); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified pai one from the data access object.
func (a *PaiOne) Delete(ctx context.Context, id string) error {
	exists, err := a.PaiOneDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Pai one not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiOneDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
