package biz

import (
	"context"
	"time"

	"github.com/xxx/testapp/internal/mods/class/dal"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
)

// 用户
type Employ struct {
	Trans     *util.Trans
	EmployDAL *dal.Employ
}

// Query employs from the data access object based on the provided parameters and options.
func (a *Employ) Query(ctx context.Context, params schema.EmployQueryParam) (*schema.EmployQueryResult, error) {
	params.Pagination = false

	result, err := a.EmployDAL.Query(ctx, params, schema.EmployQueryOptions{
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

// Get the specified employ from the data access object.
func (a *Employ) Get(ctx context.Context, id string) (*schema.Employ, error) {
	employ, err := a.EmployDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if employ == nil {
		return nil, errors.NotFound("", "Employ not found")
	}
	return employ, nil
}

// Create a new employ in the data access object.
func (a *Employ) Create(ctx context.Context, formItem *schema.EmployForm) (*schema.Employ, error) {
	employ := &schema.Employ{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(employ); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.EmployDAL.Create(ctx, employ); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return employ, nil
}

// Update the specified employ in the data access object.
func (a *Employ) Update(ctx context.Context, id string, formItem *schema.EmployForm) error {
	employ, err := a.EmployDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if employ == nil {
		return errors.NotFound("", "Employ not found")
	}

	if err := formItem.FillTo(employ); err != nil {
		return err
	}
	employ.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.EmployDAL.Update(ctx, employ); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified employ from the data access object.
func (a *Employ) Delete(ctx context.Context, id string) error {
	exists, err := a.EmployDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Employ not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.EmployDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
