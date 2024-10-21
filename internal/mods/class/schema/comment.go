package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// 讨论
type Comment struct {
	ID          string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	ActiveId    string    `json:"active_id" gorm:"size:512;comment:签到ID;"`         // 签到ID
	EmployId    string    `json:"employ_id" gorm:"comment:是否结束;"`                  // 是否结束
	EmployModel *Employ   `json:"employ_model" gorm:"-"`
	Comment     string    `json:"comment" gorm:"comment:是否开始;"`                 // 是否开始
	CreatedAt   time.Time `json:"created_at" gorm:"index;comment:Create time;"` // Create time
	UpdatedAt   time.Time `json:"updated_at" gorm:"index;comment:Update time;"` // Update time
}

func (c *Comment) AfterFind(tx *gorm.DB) error {
	tx.Where("id = ?", c.EmployId).First(&c.EmployModel)
	return nil
}

// Defining the query parameters for the `Comment` struct.
type CommentQueryParam struct {
	util.PaginationParam

	ActiveId string `form:"-"` // 签到ID
	EmployId string `form:"-"` // 是否结束
	Comment  string `form:"-"` // 是否开始
}

// Defining the query options for the `Comment` struct.
type CommentQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Comment` struct.
type CommentQueryResult struct {
	Data       Comments
	PageResult *util.PaginationResult
}

// Defining the slice of `Comment` struct.
type Comments []*Comment

// Defining the data structure for creating a `Comment` struct.
type CommentForm struct {
	ActiveId string `json:"active_id"`
	EmployId string `json:"employ_id"`
	Comment  string `json:"comment"`
}

// A validation function for the `CommentForm` struct.
func (a *CommentForm) Validate() error {
	return nil
}

// Convert `CommentForm` to `Comment` object.
func (a *CommentForm) FillTo(comment *Comment) error {
	comment.ActiveId = a.ActiveId
	comment.EmployId = a.EmployId
	comment.Comment = a.Comment
	return nil
}
