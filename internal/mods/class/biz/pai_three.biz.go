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
type PaiThree struct {
	Trans       *util.Trans
	PaiThreeDAL *dal.PaiThree
}

// Query pai threes from the data access object based on the provided parameters and options.
func (a *PaiThree) Query(ctx context.Context, params schema.PaiThreeQueryParam) (*schema.PaiThreeQueryResult, error) {
	params.Pagination = false

	result, err := a.PaiThreeDAL.Query(ctx, params, schema.PaiThreeQueryOptions{
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

// Get the specified pai three from the data access object.
func (a *PaiThree) Get(ctx context.Context, id string) (*schema.PaiThree, error) {
	paiThree, err := a.PaiThreeDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if paiThree == nil {
		return nil, errors.NotFound("", "Pai three not found")
	}
	return paiThree, nil
}

// Create a new pai three in the data access object.
func (a *PaiThree) Create(ctx context.Context, formItem *schema.PaiThreeForm) (*schema.PaiThree, error) {
	paiThree := &schema.PaiThree{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(paiThree); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiThreeDAL.Create(ctx, paiThree); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paiThree, nil
}

// Update the specified pai three in the data access object.
func (a *PaiThree) Update(ctx context.Context, id string, formItem *schema.PaiThreeForm) error {
	paiThree, err := a.PaiThreeDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if paiThree == nil {
		return errors.NotFound("", "Pai three not found")
	}

	if err := formItem.FillTo(paiThree); err != nil {
		return err
	}
	paiThree.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiThreeDAL.Update(ctx, paiThree); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified pai three from the data access object.
func (a *PaiThree) Delete(ctx context.Context, id string) error {
	exists, err := a.PaiThreeDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Pai three not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PaiThreeDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
