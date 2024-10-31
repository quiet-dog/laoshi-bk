package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
)

// 小组pk打分
type PkScore struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	ActiveId  string    `json:"active_id" gorm:"size:512;comment:活动Id;"`         // 活动Id
	Score     string    `json:"score" gorm:"comment:小组分;"`                       // 小组分
	GroupId   int       `json:"group_id" gorm:"comment:小组Id;"`                   // 小组Id
	CreatedAt time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
}

// Defining the query parameters for the `PkScore` struct.
type PkScoreQueryParam struct {
	util.PaginationParam
	ActiveId string `form:"active_id"` // 活动ID
}

// Defining the query options for the `PkScore` struct.
type PkScoreQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PkScore` struct.
type PkScoreQueryResult struct {
	Data       PkScores
	PageResult *util.PaginationResult
}

// Defining the slice of `PkScore` struct.
type PkScores []*PkScore

// Defining the data structure for creating a `PkScore` struct.
type PkScoreForm struct {
	ActiveId string `json:"active_id"`
	GroupId  int    `json:"group_id"`
	Score    string `json:"score"`
}

// A validation function for the `PkScoreForm` struct.
func (a *PkScoreForm) Validate() error {
	return nil
}

// Convert `PkScoreForm` to `PkScore` object.
func (a *PkScoreForm) FillTo(pkScore *PkScore) error {
	pkScore.ActiveId = a.ActiveId
	pkScore.GroupId = a.GroupId
	pkScore.Score = a.Score
	return nil
}
