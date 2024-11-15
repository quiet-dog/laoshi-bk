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
type PaiTwo struct {
	Trans     *util.Trans
	PaiTwoDAL *dal.PaiTwo
}

// Query pai twos from the data access object based on the provided parameters and options.
func (a *PaiTwo) Query(ctx context.Context, params schema.PaiTwoQueryParam) (*schema.PaiTwoQueryResult, error) {
	params.Pagination = false

	result, err := a.PaiTwoDAL.Query(ctx, params, schema.PaiTwoQueryOptions{
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

// Get the specified pai two from the data access object.
func (a *PaiTwo) Get(ctx context.Context, id string) (*schema.PaiTwo, error) {
	paiTwo, err := a.PaiTwoDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if paiTwo == nil {
		return nil, errors.NotFound("", "Pai two not found")
	}
	return paiTwo, nil
}

// Create a new pai two in the data access object.
func (a *PaiTwo) Create(ctx context.Context, formItem *schema.PaiTwoForm) (*schema.PaiTwo, error) {
	paiTwo := &schema.PaiTwo{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(paiTwo); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiTwoDAL.Create(ctx, paiTwo); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paiTwo, nil
}

// Update the specified pai two in the data access object.
func (a *PaiTwo) Update(ctx context.Context, id string, formItem *schema.PaiTwoForm) error {
	paiTwo, err := a.PaiTwoDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if paiTwo == nil {
		return errors.NotFound("", "Pai two not found")
	}

	if err := formItem.FillTo(paiTwo); err != nil {
		return err
	}
	paiTwo.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiTwoDAL.Update(ctx, paiTwo); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified pai two from the data access object.
func (a *PaiTwo) Delete(ctx context.Context, id string) error {
	exists, err := a.PaiTwoDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Pai two not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiTwoDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
