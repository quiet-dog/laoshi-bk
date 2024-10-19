package biz

import (
	"context"
	"encoding/json"
	"time"

	"github.com/xxx/testapp/internal/mods/class/dal"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
)

// 活动
type Active struct {
	Trans     *util.Trans
	ActiveDAL *dal.Active
}

// Query actives from the data access object based on the provided parameters and options.
func (a *Active) Query(ctx context.Context, params schema.ActiveQueryParam) (*schema.ActiveQueryResult, error) {
	params.Pagination = false

	result, err := a.ActiveDAL.Query(ctx, params, schema.ActiveQueryOptions{
		QueryOptions: util.QueryOptions{
			OrderFields: []util.OrderByParam{
				{Field: "created_at", Direction: util.DESC},
			},
			SelectFields: []string{"id", "sign_id", "is_end", "pk_id", "is_start", "tao_lun_id", "created_at", "updated_at"},
		},
	})

	for _, v := range result.Data {
		if v.SignId != "" {
			dal.GetSignDB(ctx, a.ActiveDAL.DB).Where("id = ?", v.SignId).First(&v.SignModel)
		}
		if v.PkId != "" {
			dal.GetPkDB(ctx, a.ActiveDAL.DB).Where("id = ?", v.PkId).First(&v.PkModel)
		}
		if v.TaoLunId != "" {
			dal.GetTaoLunDB(ctx, a.ActiveDAL.DB).Where("id = ?", v.TaoLunId).First(&v.TaoLunModel)
			if v.TaoLunModel != nil && v.TaoLunModel.Path != "" {
				var paths []string
				json.Unmarshal([]byte(v.TaoLunModel.Path), &paths)
				if len(paths) > 0 {
					dal.GetTaoLunDB(ctx, a.ActiveDAL.DB).Where("id in (?)", paths).Find(&v.TaoLunModel.UploadFile)
				}
			}
		}
	}

	if err != nil {
		return nil, err
	}
	return result, nil
}

// Get the specified active from the data access object.
func (a *Active) Get(ctx context.Context, id string) (*schema.Active, error) {
	active, err := a.ActiveDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if active == nil {
		return nil, errors.NotFound("", "Active not found")
	}
	return active, nil
}

// Create a new active in the data access object.
func (a *Active) Create(ctx context.Context, formItem *schema.ActiveForm) (*schema.Active, error) {
	active := &schema.Active{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(active); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ActiveDAL.Create(ctx, active); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return active, nil
}

// Update the specified active in the data access object.
func (a *Active) Update(ctx context.Context, id string, formItem *schema.ActiveForm) error {
	active, err := a.ActiveDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if active == nil {
		return errors.NotFound("", "Active not found")
	}

	if err := formItem.FillTo(active); err != nil {
		return err
	}
	active.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ActiveDAL.Update(ctx, active); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified active from the data access object.
func (a *Active) Delete(ctx context.Context, id string) error {
	exists, err := a.ActiveDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Active not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ActiveDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
