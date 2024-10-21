package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get comment storage instance
func GetCommentDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Comment))
}

// 讨论
type Comment struct {
	DB *gorm.DB
}

// Query comments from the database based on the provided parameters and options.
func (a *Comment) Query(ctx context.Context, params schema.CommentQueryParam, opts ...schema.CommentQueryOptions) (*schema.CommentQueryResult, error) {
	var opt schema.CommentQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetCommentDB(ctx, a.DB)
	if v := params.ActiveId; len(v) > 0 {
		db = db.Where("active_id = ?", v)
	}
	if v := params.EmployId; len(v) > 0 {
		db = db.Where("employ_id = ?", v)
	}
	if v := params.Comment; len(v) > 0 {
		db = db.Where("comment = ?", v)
	}

	var list schema.Comments
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.CommentQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified comment from the database.
func (a *Comment) Get(ctx context.Context, id string, opts ...schema.CommentQueryOptions) (*schema.Comment, error) {
	var opt schema.CommentQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Comment)
	ok, err := util.FindOne(ctx, GetCommentDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified comment exists in the database.
func (a *Comment) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetCommentDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new comment.
func (a *Comment) Create(ctx context.Context, item *schema.Comment) error {
	result := GetCommentDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified comment in the database.
func (a *Comment) Update(ctx context.Context, item *schema.Comment) error {
	result := GetCommentDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified comment from the database.
func (a *Comment) Delete(ctx context.Context, id string) error {
	result := GetCommentDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Comment))
	return errors.WithStack(result.Error)
}
