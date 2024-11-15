package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Pk日志
type PaiFour struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	ActiveId  string    `json:"active_id" gorm:"size:512;comment:签到标题;"`         // 签到标题
	EmployId  string    `json:"employ_id" gorm:"comment:自动结束;"`                  // 自动结束
	Employ    *Employ   `json:"employ" gorm:"-"`                                 // 自动结束
	Score     int       `json:"score" gorm:"comment:自动结束;"`                      // 自动结束
	Level     string    `json:"level" gorm:"comment:自动结束;"`                      // 自动结束
	CreatedAt time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
}

func (p *PaiFour) AfterFind(tx *gorm.DB) error {
	tx.Where("id = ?", p.EmployId).First(&p.Employ)
	p.Level = "高级工"
	if p.Score < 80 {
		p.Level = "中级工"
	}
	if p.Score < 70 {
		p.Level = "初级工"
	}
	if p.Score < 60 {
		p.Level = "学徒工"
	}
	return nil
}

// Defining the query parameters for the `PaiFour` struct.
type PaiFourQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `PaiFour` struct.
type PaiFourQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PaiFour` struct.
type PaiFourQueryResult struct {
	Data       PaiFours
	PageResult *util.PaginationResult
}

// Defining the slice of `PaiFour` struct.
type PaiFours []*PaiFour

// Defining the data structure for creating a `PaiFour` struct.
type PaiFourForm struct {
	Score int `json:"score"` // 自动结束
}

// A validation function for the `PaiFourForm` struct.
func (a *PaiFourForm) Validate() error {
	return nil
}

// Convert `PaiFourForm` to `PaiFour` object.
func (a *PaiFourForm) FillTo(paiFour *PaiFour) error {
	paiFour.Score = a.Score
	return nil
}
