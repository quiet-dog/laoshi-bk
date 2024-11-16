package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Pk日志
type PaiFive struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	ActiveId  string    `json:"active_id" gorm:"size:512;comment:签到标题;"`         // 签到标题
	EmployId  string    `json:"employ_id" gorm:"comment:自动结束;"`                  // 自动结束
	Employ    *Employ   `json:"employ" gorm:"-"`
	Score     int       `json:"score" gorm:"comment:自动结束;"`                   // 自动结束
	CreatedAt time.Time `json:"created_at" gorm:"index;comment:Create time;"` // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;comment:Update time;"` // Update time
}

func (p *PaiFive) AfterFind(tx *gorm.DB) error {
	tx.Where("id = ?", p.EmployId).First(&p.Employ)
	return nil
}

// Defining the query parameters for the `PaiFive` struct.
type PaiFiveQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `PaiFive` struct.
type PaiFiveQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PaiFive` struct.
type PaiFiveQueryResult struct {
	Data       PaiFives
	PageResult *util.PaginationResult
}

// Defining the slice of `PaiFive` struct.
type PaiFives []*PaiFive

// Defining the data structure for creating a `PaiFive` struct.
type PaiFiveForm struct {
	Score int `json:"score" form:"score"`
}

// A validation function for the `PaiFiveForm` struct.
func (a *PaiFiveForm) Validate() error {
	return nil
}

// Convert `PaiFiveForm` to `PaiFive` object.
func (a *PaiFiveForm) FillTo(paiFive *PaiFive) error {
	paiFive.Score = a.Score
	return nil
}
