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
type Comment struct {
	Trans      *util.Trans
	CommentDAL *dal.Comment
}

// Query comments from the data access object based on the provided parameters and options.
func (a *Comment) Query(ctx context.Context, params schema.CommentQueryParam) (*schema.CommentQueryResult, error) {
	params.Pagination = false

	result, err := a.CommentDAL.Query(ctx, params, schema.CommentQueryOptions{
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

// Get the specified comment from the data access object.
func (a *Comment) Get(ctx context.Context, id string) (*schema.Comment, error) {
	comment, err := a.CommentDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if comment == nil {
		return nil, errors.NotFound("", "Comment not found")
	}
	return comment, nil
}

// Create a new comment in the data access object.
func (a *Comment) Create(ctx context.Context, formItem *schema.CommentForm) (*schema.Comment, error) {
	comment := &schema.Comment{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(comment); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CommentDAL.Create(ctx, comment); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return comment, nil
}

// Update the specified comment in the data access object.
func (a *Comment) Update(ctx context.Context, id string, formItem *schema.CommentForm) error {
	comment, err := a.CommentDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if comment == nil {
		return errors.NotFound("", "Comment not found")
	}

	if err := formItem.FillTo(comment); err != nil {
		return err
	}
	comment.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CommentDAL.Update(ctx, comment); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified comment from the data access object.
func (a *Comment) Delete(ctx context.Context, id string) error {
	exists, err := a.CommentDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Comment not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CommentDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
