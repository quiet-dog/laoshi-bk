package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Pk日志
type PaiTwo struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	ActiveId  string    `json:"active_id" gorm:"size:512;comment:签到标题;"`         // 签到标题
	EmployId  string    `json:"employ_id" gorm:"comment:自动结束;"`                  // 自动结束
	Employ    *Employ   `json:"employ" gorm:"-"`                                 // 自动结束
	Score     int       `json:"score" gorm:"comment:自动结束;"`                      // 自动结束
	CreatedAt time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
}

func (p *PaiTwo) AfterFind(tx *gorm.DB) error {
	tx.Where("id = ?", p.EmployId).First(&p.Employ)
	return nil
}

// Defining the query parameters for the `PaiTwo` struct.
type PaiTwoQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `PaiTwo` struct.
type PaiTwoQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PaiTwo` struct.
type PaiTwoQueryResult struct {
	Data       PaiTwos
	PageResult *util.PaginationResult
}

// Defining the slice of `PaiTwo` struct.
type PaiTwos []*PaiTwo

// Defining the data structure for creating a `PaiTwo` struct.
type PaiTwoForm struct {
	Score int `json:"score"` // 自动结束
}

// A validation function for the `PaiTwoForm` struct.
func (a *PaiTwoForm) Validate() error {
	return nil
}

// Convert `PaiTwoForm` to `PaiTwo` object.
func (a *PaiTwoForm) FillTo(paiTwo *PaiTwo) error {
	paiTwo.Score = a.Score
	return nil
}
