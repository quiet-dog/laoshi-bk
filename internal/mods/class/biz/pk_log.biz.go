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
type PkLog struct {
	Trans    *util.Trans
	PkLogDAL *dal.PkLog
}

// Query pk logs from the data access object based on the provided parameters and options.
func (a *PkLog) Query(ctx context.Context, params schema.PkLogQueryParam) (*schema.PkLogQueryResult, error) {
	params.Pagination = false

	result, err := a.PkLogDAL.Query(ctx, params, schema.PkLogQueryOptions{
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

// Get the specified pk log from the data access object.
func (a *PkLog) Get(ctx context.Context, id string) (*schema.PkLog, error) {
	pkLog, err := a.PkLogDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if pkLog == nil {
		return nil, errors.NotFound("", "Pk log not found")
	}
	return pkLog, nil
}

// Create a new pk log in the data access object.
func (a *PkLog) Create(ctx context.Context, formItem *schema.PkLogForm) (*schema.PkLog, error) {
	pkLog := &schema.PkLog{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(pkLog); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PkLogDAL.Create(ctx, pkLog); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return pkLog, nil
}

// Update the specified pk log in the data access object.
func (a *PkLog) Update(ctx context.Context, id string, formItem *schema.PkLogForm) error {
	pkLog, err := a.PkLogDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if pkLog == nil {
		return errors.NotFound("", "Pk log not found")
	}

	if err := formItem.FillTo(pkLog); err != nil {
		return err
	}
	pkLog.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PkLogDAL.Update(ctx, pkLog); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified pk log from the data access object.
func (a *PkLog) Delete(ctx context.Context, id string) error {
	exists, err := a.PkLogDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Pk log not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PkLogDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
