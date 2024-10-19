package biz

import (
	"context"
	"time"

	"github.com/xxx/testapp/internal/mods/class/dal"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
)

// 小组pk
type Pk struct {
	Trans *util.Trans
	PkDAL *dal.Pk
}

// Query pks from the data access object based on the provided parameters and options.
func (a *Pk) Query(ctx context.Context, params schema.PkQueryParam) (*schema.PkQueryResult, error) {
	params.Pagination = false

	result, err := a.PkDAL.Query(ctx, params, schema.PkQueryOptions{
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

// Get the specified pk from the data access object.
func (a *Pk) Get(ctx context.Context, id string) (*schema.Pk, error) {
	pk, err := a.PkDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if pk == nil {
		return nil, errors.NotFound("", "Pk not found")
	}
	return pk, nil
}

// Create a new pk in the data access object.
func (a *Pk) Create(ctx context.Context, formItem *schema.PkForm) (*schema.Pk, error) {
	pk := &schema.Pk{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(pk); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PkDAL.Create(ctx, pk); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return pk, nil
}

// Update the specified pk in the data access object.
func (a *Pk) Update(ctx context.Context, id string, formItem *schema.PkForm) error {
	pk, err := a.PkDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if pk == nil {
		return errors.NotFound("", "Pk not found")
	}

	if err := formItem.FillTo(pk); err != nil {
		return err
	}
	pk.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PkDAL.Update(ctx, pk); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified pk from the data access object.
func (a *Pk) Delete(ctx context.Context, id string) error {
	exists, err := a.PkDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Pk not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PkDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
