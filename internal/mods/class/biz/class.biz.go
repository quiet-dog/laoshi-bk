package biz

import (
	"context"
	"time"

	"github.com/xxx/testapp/internal/mods/class/dal"
	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
)

// 课件
type Class struct {
	Trans    *util.Trans
	ClassDAL *dal.Class
}

// Query classes from the data access object based on the provided parameters and options.
func (a *Class) Query(ctx context.Context, params schema.ClassQueryParam) (*schema.ClassQueryResult, error) {
	params.Pagination = false

	result, err := a.ClassDAL.Query(ctx, params, schema.ClassQueryOptions{
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

// Get the specified class from the data access object.
func (a *Class) Get(ctx context.Context, id string) (*schema.Class, error) {
	class, err := a.ClassDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if class == nil {
		return nil, errors.NotFound("", "Class not found")
	}
	return class, nil
}

// Create a new class in the data access object.
func (a *Class) Create(ctx context.Context, formItem *schema.ClassForm) (*schema.Class, error) {
	class := &schema.Class{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(class); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ClassDAL.Create(ctx, class); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return class, nil
}

// Update the specified class in the data access object.
func (a *Class) Update(ctx context.Context, id string, formItem *schema.ClassForm) error {
	class, err := a.ClassDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if class == nil {
		return errors.NotFound("", "Class not found")
	}

	if err := formItem.FillTo(class); err != nil {
		return err
	}
	class.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ClassDAL.Update(ctx, class); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified class from the data access object.
func (a *Class) Delete(ctx context.Context, id string) error {
	exists, err := a.ClassDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Class not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ClassDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified class from the data access object.
func (a *Class) Tree(ctx context.Context) []*schema.Class {
	// exists, err := a.ClassDAL.Exists(ctx, id)
	// if err != nil {
	// 	return err
	// } else if !exists {
	// 	return errors.NotFound("", "Class not found")
	// }

	var classes []*schema.Class
	a.Trans.DB.Where("pid = ?", "").Find(&classes)
	if len(classes) > 0 {
		a.get(&classes)
	}

	return classes
}

func (a *Class) get(c *[]*schema.Class) {
	for _, class := range *c {
		var children []*schema.Class
		a.Trans.DB.Where("pid = ?", class.ID).Find(&children)
		if len(children) == 0 {
			continue
		}
		class.Children = children
		a.get(&children)
	}
}
