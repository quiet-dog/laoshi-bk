package biz

import (
	"context"
	"time"

	"github.com/xxx/testapp/internal/mods/class/dal"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
)

// 小组pk打分
type PkScore struct {
	Trans      *util.Trans
	PkScoreDAL *dal.PkScore
}

// Query pk scores from the data access object based on the provided parameters and options.
func (a *PkScore) Query(ctx context.Context, params schema.PkScoreQueryParam) (*schema.PkScoreQueryResult, error) {
	params.Pagination = false

	result, err := a.PkScoreDAL.Query(ctx, params, schema.PkScoreQueryOptions{
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

// Get the specified pk score from the data access object.
func (a *PkScore) Get(ctx context.Context, id string) (*schema.PkScore, error) {
	pkScore, err := a.PkScoreDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if pkScore == nil {
		return nil, errors.NotFound("", "Pk score not found")
	}
	return pkScore, nil
}

// Create a new pk score in the data access object.
func (a *PkScore) Create(ctx context.Context, formItem *schema.PkScoreForm) (*schema.PkScore, error) {
	pkScore := &schema.PkScore{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(pkScore); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PkScoreDAL.Create(ctx, pkScore); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return pkScore, nil
}

// Update the specified pk score in the data access object.
func (a *PkScore) Update(ctx context.Context, id string, formItem *schema.PkScoreForm) error {
	pkScore, err := a.PkScoreDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if pkScore == nil {
		return errors.NotFound("", "Pk score not found")
	}

	if err := formItem.FillTo(pkScore); err != nil {
		return err
	}
	pkScore.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PkScoreDAL.Update(ctx, pkScore); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified pk score from the data access object.
func (a *PkScore) Delete(ctx context.Context, id string) error {
	exists, err := a.PkScoreDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Pk score not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.PkScoreDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
