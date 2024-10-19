package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
)

// 签到
type Sign struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	Title     string    `json:"title" gorm:"size:512;comment:签到标题;"`             // 签到标题
	IsAuto    bool      `json:"is_auto" gorm:"comment:自动结束;"`                    // 自动结束
	Type      string    `json:"type" gorm:"size:1024;comment:签到类型;"`             // 签到类型
	CreatedAt time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
}

// Defining the query parameters for the `Sign` struct.
type SignQueryParam struct {
	util.PaginationParam

	Type string `form:"-"` // 签到类型
}

// Defining the query options for the `Sign` struct.
type SignQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Sign` struct.
type SignQueryResult struct {
	Data       Signs
	PageResult *util.PaginationResult
}

// Defining the slice of `Sign` struct.
type Signs []*Sign

// Defining the data structure for creating a `Sign` struct.
type SignForm struct {
	Title string `json:"title" binding:"required,max=512"`        // 签到标题
	Type  string `json:"type" binding:"required,oneof=一键签到 扫码签到"` // 签到类型
}

// A validation function for the `SignForm` struct.
func (a *SignForm) Validate() error {
	return nil
}

// Convert `SignForm` to `Sign` object.
func (a *SignForm) FillTo(sign *Sign) error {
	sign.Title = a.Title
	sign.Type = a.Type
	return nil
}
