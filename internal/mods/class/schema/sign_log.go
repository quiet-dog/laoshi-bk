package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// 签到日志
type SignLog struct {
	ID          string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	ActiveId    string    `json:"active_id" gorm:"size:512;comment:签到标题;"`         // 签到标题
	EmployId    string    `json:"employ_id" gorm:"comment:自动结束;"`                  // 自动结束
	EmployModel *Employ   `json:"employ_model" gorm:"-"`                           // 自动结束
	CreatedAt   time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt   time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
}

func (s *SignLog) AfterFind(tx *gorm.DB) error {
	tx.Where("id = ?", s.EmployId).First(&s.EmployModel)
	return nil
}

// Defining the query parameters for the `SignLog` struct.
type SignLogQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `SignLog` struct.
type SignLogQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `SignLog` struct.
type SignLogQueryResult struct {
	Data       SignLogs
	PageResult *util.PaginationResult
}

// Defining the slice of `SignLog` struct.
type SignLogs []*SignLog

// Defining the data structure for creating a `SignLog` struct.
type SignLogForm struct {
	ActiveId string `json:"active_id"`
	EmployId string `json:"employ_id"`
}

// A validation function for the `SignLogForm` struct.
func (a *SignLogForm) Validate() error {
	return nil
}

// Convert `SignLogForm` to `SignLog` object.
func (a *SignLogForm) FillTo(signLog *SignLog) error {
	signLog.ActiveId = a.ActiveId
	signLog.EmployId = a.EmployId
	return nil
}
