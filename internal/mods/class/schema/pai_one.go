package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Pk日志
type PaiOne struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	EmployId  string    `json:"employ_id" gorm:"comment:自动结束;"`                  // 自动结束
	Employ    *Employ   `json:"employ" gorm:"-"`                                 // 自动结束
	Score     int       `json:"score" gorm:"comment:自动结束;"`                      // 自动结束
	CreatedAt time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
}

func (p *PaiOne) AfterFind(tx *gorm.DB) error {
	tx.Where("id = ?", p.EmployId).First(&p.Employ)
	return nil
}

// Defining the query parameters for the `PaiOne` struct.
type PaiOneQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `PaiOne` struct.
type PaiOneQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PaiOne` struct.
type PaiOneQueryResult struct {
	Data       PaiOnes
	PageResult *util.PaginationResult
}

// Defining the slice of `PaiOne` struct.
type PaiOnes []*PaiOne

// Defining the data structure for creating a `PaiOne` struct.
type PaiOneForm struct {
	Score int `json:"score"` // 自动结束
}

// A validation function for the `PaiOneForm` struct.
func (a *PaiOneForm) Validate() error {
	return nil
}

// Convert `PaiOneForm` to `PaiOne` object.
func (a *PaiOneForm) FillTo(paiOne *PaiOne) error {
	paiOne.Score = a.Score
	return nil
}
