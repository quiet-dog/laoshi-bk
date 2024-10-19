package biz

import (
	"context"
	"time"

	"github.com/xxx/testapp/internal/mods/class/dal"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
)

// 讨论
type TaoLun struct {
	Trans     *util.Trans
	TaoLunDAL *dal.TaoLun
}

// Query tao luns from the data access object based on the provided parameters and options.
func (a *TaoLun) Query(ctx context.Context, params schema.TaoLunQueryParam) (*schema.TaoLunQueryResult, error) {
	params.Pagination = false

	result, err := a.TaoLunDAL.Query(ctx, params, schema.TaoLunQueryOptions{
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

// Get the specified tao lun from the data access object.
func (a *TaoLun) Get(ctx context.Context, id string) (*schema.TaoLun, error) {
	taoLun, err := a.TaoLunDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if taoLun == nil {
		return nil, errors.NotFound("", "Tao lun not found")
	}
	return taoLun, nil
}

// Create a new tao lun in the data access object.
func (a *TaoLun) Create(ctx context.Context, formItem *schema.TaoLunForm) (*schema.TaoLun, error) {
	taoLun := &schema.TaoLun{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(taoLun); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.TaoLunDAL.Create(ctx, taoLun); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return taoLun, nil
}

// Update the specified tao lun in the data access object.
func (a *TaoLun) Update(ctx context.Context, id string, formItem *schema.TaoLunForm) error {
	taoLun, err := a.TaoLunDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if taoLun == nil {
		return errors.NotFound("", "Tao lun not found")
	}

	if err := formItem.FillTo(taoLun); err != nil {
		return err
	}
	taoLun.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.TaoLunDAL.Update(ctx, taoLun); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified tao lun from the data access object.
func (a *TaoLun) Delete(ctx context.Context, id string) error {
	exists, err := a.TaoLunDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Tao lun not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.TaoLunDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
