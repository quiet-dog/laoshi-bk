package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
)

// 小组pk
type Pk struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	Title     string    `json:"title" gorm:"size:512;comment:pk标题;"`             // pk标题
	Count     int       `json:"count" gorm:"comment:小组数量;"`                      // 小组数量
	CreatedAt time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
}

// Defining the query parameters for the `Pk` struct.
type PkQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `Pk` struct.
type PkQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Pk` struct.
type PkQueryResult struct {
	Data       Pks
	PageResult *util.PaginationResult
}

// Defining the slice of `Pk` struct.
type Pks []*Pk

// Defining the data structure for creating a `Pk` struct.
type PkForm struct {
	Count int    `json:"count"`
	Title string `json:"title"`
}

// A validation function for the `PkForm` struct.
func (a *PkForm) Validate() error {
	return nil
}

// Convert `PkForm` to `Pk` object.
func (a *PkForm) FillTo(pk *Pk) error {
	pk.Count = a.Count
	pk.Title = a.Title
	return nil
}
