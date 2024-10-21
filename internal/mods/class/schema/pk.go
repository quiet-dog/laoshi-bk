package schema

import (
	"encoding/json"
	"time"

	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// 小组pk
type Pk struct {
	ID          string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	Title       string    `json:"title" gorm:"size:512;comment:pk标题;"`             // pk标题
	Count       int       `json:"count" gorm:"comment:小组数量;"`                      // 小组数量
	PkModel     []*Pk     `json:"pk_model" gorm:"-"`                               // 小组
	Paths       string    `json:"paths" gorm:"comment:文件地址;"`                      // 文件地址
	UploadFiles []*File   `json:"upload_fils" gorm:"-"`                            // 文件
	CreatedAt   time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt   time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
}

func (p *Pk) AfterFind(tx *gorm.DB) error {
	tx.Where("active_id = ?", tx.Model(&Active{}).Where("pk_id = ?", p.ID).Select("id")).Find(&p.PkModel)
	var paths []string
	json.Unmarshal([]byte(p.Paths), &paths)
	if len(paths) > 0 {
		tx.Where("id in (?)", paths).Find(&p.UploadFiles)
	}
	return nil
}

// Defining the query parameters for the `Pk` struct.
type PkQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `Pk` struct.
type PkQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Pk` struct.
type PkQueryResult struct {
	Data       Pks
	PageResult *util.PaginationResult
}

// Defining the slice of `Pk` struct.
type Pks []*Pk

// Defining the data structure for creating a `Pk` struct.
type PkForm struct {
	Count int    `json:"count"`
	Title string `json:"title"`
	Paths string `json:"paths"`
}

// A validation function for the `PkForm` struct.
func (a *PkForm) Validate() error {
	return nil
}

// Convert `PkForm` to `Pk` object.
func (a *PkForm) FillTo(pk *Pk) error {
	pk.Count = a.Count
	pk.Title = a.Title
	pk.Paths = a.Paths
	return nil
}
