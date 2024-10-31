package dal

import (
	"context"

	"github.com/xxx/testapp/internal/mods/class/schema"
	"github.com/xxx/testapp/pkg/errors"
	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Get pk score storage instance
func GetPkScoreDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.PkScore))
}

// 小组pk打分
type PkScore struct {
	DB *gorm.DB
}

// Query pk scores from the database based on the provided parameters and options.
func (a *PkScore) Query(ctx context.Context, params schema.PkScoreQueryParam, opts ...schema.PkScoreQueryOptions) (*schema.PkScoreQueryResult, error) {
	var opt schema.PkScoreQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetPkScoreDB(ctx, a.DB)

	if v := params.ActiveId; v != "" {
		db = db.Where("active_id = ?", v)
	}

	var list schema.PkScores
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.PkScoreQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified pk score from the database.
func (a *PkScore) Get(ctx context.Context, id string, opts ...schema.PkScoreQueryOptions) (*schema.PkScore, error) {
	var opt schema.PkScoreQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.PkScore)
	ok, err := util.FindOne(ctx, GetPkScoreDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified pk score exists in the database.
func (a *PkScore) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetPkScoreDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new pk score.
func (a *PkScore) Create(ctx context.Context, item *schema.PkScore) error {
	result := GetPkScoreDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified pk score in the database.
func (a *PkScore) Update(ctx context.Context, item *schema.PkScore) error {
	result := GetPkScoreDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified pk score from the database.
func (a *PkScore) Delete(ctx context.Context, id string) error {
	result := GetPkScoreDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.PkScore))
	return errors.WithStack(result.Error)
}
