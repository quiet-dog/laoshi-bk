package biz

import (
	"context"

	"github.com/xxx/testapp/internal/mods/sys/dal"
	"github.com/xxx/testapp/internal/mods/sys/schema"
	"github.com/xxx/testapp/pkg/util"
)

// Logger management
type Logger struct {
	LoggerDAL *dal.Logger
}

// Query loggers from the data access object based on the provided parameters and options.
func (a *Logger) Query(ctx context.Context, params schema.LoggerQueryParam) (*schema.LoggerQueryResult, error) {
	params.Pagination = true

	result, err := a.LoggerDAL.Query(ctx, params, schema.LoggerQueryOptions{
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