package biz

import (
	"context"
	"time"

	"github.com/xxx/testapp/internal/mods/class/dal"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
)

// 签到日志
type SignLog struct {
	Trans      *util.Trans
	SignLogDAL *dal.SignLog
}

// Query sign logs from the data access object based on the provided parameters and options.
func (a *SignLog) Query(ctx context.Context, params schema.SignLogQueryParam) (*schema.SignLogQueryResult, error) {
	params.Pagination = false

	result, err := a.SignLogDAL.Query(ctx, params, schema.SignLogQueryOptions{
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

// Get the specified sign log from the data access object.
func (a *SignLog) Get(ctx context.Context, id string) (*schema.SignLog, error) {
	signLog, err := a.SignLogDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if signLog == nil {
		return nil, errors.NotFound("", "Sign log not found")
	}
	return signLog, nil
}

// Create a new sign log in the data access object.
func (a *SignLog) Create(ctx context.Context, formItem *schema.SignLogForm) (*schema.SignLog, error) {

	var haveSignLog schema.SignLog
	if a.SignLogDAL.DB.Where("active_id = ? and employ_id = ?", formItem.ActiveId, formItem.EmployId).First(&haveSignLog).Error == nil {
		return &haveSignLog, nil
	}

	signLog := &schema.SignLog{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(signLog); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.SignLogDAL.Create(ctx, signLog); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return signLog, nil
}

// Update the specified sign log in the data access object.
func (a *SignLog) Update(ctx context.Context, id string, formItem *schema.SignLogForm) error {
	signLog, err := a.SignLogDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if signLog == nil {
		return errors.NotFound("", "Sign log not found")
	}

	if err := formItem.FillTo(signLog); err != nil {
		return err
	}
	signLog.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.SignLogDAL.Update(ctx, signLog); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified sign log from the data access object.
func (a *SignLog) Delete(ctx context.Context, id string) error {
	exists, err := a.SignLogDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Sign log not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.SignLogDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
