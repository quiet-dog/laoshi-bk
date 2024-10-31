package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// Pk日志
type PkLog struct {
	ID          string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	ActiveId    string    `json:"active_id" gorm:"size:512;comment:签到标题;"`         // 签到标题
	EmployId    string    `json:"employ_id" gorm:"comment:自动结束;"`                  // 自动结束
	EmployModel *Employ   `json:"employ_model" gorm:"-"`                           // 自动结束
	GroupId     int       `json:"group_id" gorm:"comment:自动结束;"`                   // 自动结束
	Comment     string    `json:"comment" gorm:"comment:自动结束;"`                    // 自动结束
	CreatedAt   time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt   time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
}

func (p *PkLog) AfterFind(tx *gorm.DB) error {
	tx.Where("id = ?", p.EmployId).First(&p.EmployModel)
	return nil
}

// Defining the query parameters for the `PkLog` struct.
type PkLogQueryParam struct {
	util.PaginationParam
	ActiveId string `form:"active_id"` // 活动ID
}

// Defining the query options for the `PkLog` struct.
type PkLogQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `PkLog` struct.
type PkLogQueryResult struct {
	Data       PkLogs
	PageResult *util.PaginationResult
}

// Defining the slice of `PkLog` struct.
type PkLogs []*PkLog

// Defining the data structure for creating a `PkLog` struct.
type PkLogForm struct {
	ActiveId string `json:"active_id"`
	EmployId string `json:"employ_id"`
	GroupId  int    `json:"group_id"`
	Comment  string `json:"comment"`
}

// A validation function for the `PkLogForm` struct.
func (a *PkLogForm) Validate() error {
	return nil
}

// Convert `PkLogForm` to `PkLog` object.
func (a *PkLogForm) FillTo(pkLog *PkLog) error {
	pkLog.ActiveId = a.ActiveId
	pkLog.EmployId = a.EmployId
	pkLog.GroupId = a.GroupId
	pkLog.Comment = a.Comment
	return nil
}
