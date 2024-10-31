package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// 用户
type Employ struct {
	ID          string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	Name        string    `json:"name" gorm:"size:512;comment:姓名;"`                // 姓名
	Username    string    `json:"username" gorm:"comment:用户名;"`                    // 用户名
	Password    string    `json:"password" gorm:"comment:密码;"`                     // 密码
	IsTeacher   bool      `json:"is_teacher" gorm:"comment:是否是老师;"`                // 是否是老师
	IsCommittee bool      `json:"is_committee" gorm:"comment:是否是委员;"`              // 是否是委员
	Committee   int       `json:"committee" gorm:"comment:组号;"`                    // 组号
	CreatedAt   time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt   time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
	Avatar      string    `json:"avatar" gorm:"comment:头像;"`                       // 头像
}

func (e *Employ) AfterFind(tx *gorm.DB) error {
	e.Avatar = "upload/user/ava/" + e.Name + ".png"
	return nil
}

// Defining the query parameters for the `Employ` struct.
type EmployQueryParam struct {
	util.PaginationParam

	Name        string `form:"name"`      // 姓名
	Username    string `form:"username"`  // 用户名
	Password    string `form:"password"`  // 密码
	IsTeacher   bool   `form:"-"`         // 是否是老师
	IsCommittee bool   `form:"-"`         // 是否是委员
	Committee   int    `form:"committee"` // 组号
}

// Defining the query options for the `Employ` struct.
type EmployQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Employ` struct.
type EmployQueryResult struct {
	Data       Employs
	PageResult *util.PaginationResult
}

// Defining the slice of `Employ` struct.
type Employs []*Employ

// Defining the data structure for creating a `Employ` struct.
type EmployForm struct {
	IsCommittee bool   `json:"is_committee"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	IsTeacher   bool   `json:"is_teacher"`
	Committee   int    `json:"committee"`
}

// A validation function for the `EmployForm` struct.
func (a *EmployForm) Validate() error {
	return nil
}

// Convert `EmployForm` to `Employ` object.
func (a *EmployForm) FillTo(employ *Employ) error {
	employ.IsCommittee = a.IsCommittee
	employ.Username = a.Username
	employ.Password = a.Password
	employ.Name = a.Name
	employ.IsTeacher = a.IsTeacher
	employ.Committee = a.Committee
	return nil
}
