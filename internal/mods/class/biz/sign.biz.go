package biz

import (
	"context"
	"time"

	"github.com/xxx/testapp/internal/mods/class/dal"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
)

// 签到
type Sign struct {
	Trans   *util.Trans
	SignDAL *dal.Sign
}

// Query signs from the data access object based on the provided parameters and options.
func (a *Sign) Query(ctx context.Context, params schema.SignQueryParam) (*schema.SignQueryResult, error) {
	params.Pagination = false

	result, err := a.SignDAL.Query(ctx, params, schema.SignQueryOptions{
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

// Get the specified sign from the data access object.
func (a *Sign) Get(ctx context.Context, id string) (*schema.Sign, error) {
	sign, err := a.SignDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if sign == nil {
		return nil, errors.NotFound("", "Sign not found")
	}
	return sign, nil
}

// Create a new sign in the data access object.
func (a *Sign) Create(ctx context.Context, formItem *schema.SignForm) (*schema.Sign, error) {
	sign := &schema.Sign{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(sign); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.SignDAL.Create(ctx, sign); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return sign, nil
}

// Update the specified sign in the data access object.
func (a *Sign) Update(ctx context.Context, id string, formItem *schema.SignForm) error {
	sign, err := a.SignDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if sign == nil {
		return errors.NotFound("", "Sign not found")
	}

	if err := formItem.FillTo(sign); err != nil {
		return err
	}
	sign.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.SignDAL.Update(ctx, sign); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified sign from the data access object.
func (a *Sign) Delete(ctx context.Context, id string) error {
	exists, err := a.SignDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Sign not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.SignDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
