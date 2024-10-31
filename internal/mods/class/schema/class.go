package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// 课件
type Class struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	Name      string    `json:"name" gorm:"size:512;comment:签到ID;"`              // 签到ID
	Path      string    `json:"path" gorm:"comment:是否结束;"`                       // 是否结束
	Pid       string    `json:"pid" gorm:"comment:是否开始;"`                        // 是否开始
	CreatedAt time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
	Children  []*Class  `json:"children" gorm:"-"`                               // 子课件
	Label     string    `json:"label" gorm:"-"`                                  // 标签
	Key       string    `json:"key" gorm:"-"`                                    // 键
	Value     string    `json:"value" gorm:"-"`                                  // 值
}

func (c *Class) AfterFind(tx *gorm.DB) error {
	c.Label = c.Name
	c.Key = c.ID
	c.Value = c.ID
	return nil
}

// Defining the query parameters for the `Class` struct.
type ClassQueryParam struct {
	util.PaginationParam

	Name string `form:"-"` // 签到ID
	Path string `form:"-"` // 是否结束
	Pid  string `form:"-"` // 是否开始
}

// Defining the query options for the `Class` struct.
type ClassQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Class` struct.
type ClassQueryResult struct {
	Data       Classes
	PageResult *util.PaginationResult
}

// Defining the slice of `Class` struct.
type Classes []*Class

// Defining the data structure for creating a `Class` struct.
type ClassForm struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Pid  string `json:"pid"`
}

// A validation function for the `ClassForm` struct.
func (a *ClassForm) Validate() error {
	return nil
}

// Convert `ClassForm` to `Class` object.
func (a *ClassForm) FillTo(class *Class) error {
	class.Name = a.Name
	class.Path = a.Path
	class.Pid = a.Pid
	return nil
}
