package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Pk日志
type PaiThree struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	ActiveId  string    `json:"active_id" gorm:"size:512;comment:签到标题;"`         // 签到标题
	EmployId  string    `json:"employ_id" gorm:"comment:自动结束;"`                  // 自动结束
	Employ    *Employ   `json:"employ" gorm:"-"`                                 // 自动结束
	Score     int       `json:"score" gorm:"comment:自动结束;"`                      // 自动结束
	CreatedAt time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
}

func (p *PaiThree) AfterFind(tx *gorm.DB) error {
	tx.Where("id = ?", p.EmployId).First(&p.Employ)
	return nil
}

// Defining the query parameters for the `PaiThree` struct.
type PaiThreeQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `PaiThree` struct.
type PaiThreeQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PaiThree` struct.
type PaiThreeQueryResult struct {
	Data       PaiThrees
	PageResult *util.PaginationResult
}

// Defining the slice of `PaiThree` struct.
type PaiThrees []*PaiThree

// Defining the data structure for creating a `PaiThree` struct.
type PaiThreeForm struct {
	Score int `json:"score"` // 自动结束
}

// A validation function for the `PaiThreeForm` struct.
func (a *PaiThreeForm) Validate() error {
	return nil
}

// Convert `PaiThreeForm` to `PaiThree` object.
func (a *PaiThreeForm) FillTo(paiThree *PaiThree) error {
	paiThree.Score = a.Score
	return nil
}
