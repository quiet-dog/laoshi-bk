package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
)

// 文件
type File struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	OrignName string    `json:"orign_name" gorm:"size:512;comment:原始名字;"`        // 原始名字
	Path      string    `json:"path" gorm:"comment:路径;"`                         // 路径
	CreatedAt time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
}

// Defining the query parameters for the `File` struct.
type FileQueryParam struct {
	util.PaginationParam

	OrignName string `form:"-"` // 原始名字
	Path      string `form:"-"` // 路径
}

// Defining the query options for the `File` struct.
type FileQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `File` struct.
type FileQueryResult struct {
	Data       Files
	PageResult *util.PaginationResult
}

// Defining the slice of `File` struct.
type Files []*File

// Defining the data structure for creating a `File` struct.
type FileForm struct {
	OrignName string `json:"orign_name"` // 原始名字
	Path      string `json:"path"`       // 路径
}

// A validation function for the `FileForm` struct.
func (a *FileForm) Validate() error {
	return nil
}

// Convert `FileForm` to `File` object.
func (a *FileForm) FillTo(file *File) error {
	file.OrignName = a.OrignName
	file.Path = a.Path
	return nil
}
