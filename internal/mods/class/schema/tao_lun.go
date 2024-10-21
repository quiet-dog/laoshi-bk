package schema

import (
	"encoding/json"
	"time"

	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// 讨论
type TaoLun struct {
	ID           string     `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	Title        string     `json:"title" gorm:"size:512;comment:签到标题;"`             // 签到标题
	Comment      string     `json:"comment" gorm:"comment:讨论内容;"`                    // 讨论内容
	Time         string     `json:"time" gorm:"comment:讨论时长;"`                       // 讨论时长
	Path         string     `json:"path" gorm:"comment:文件地址;"`                       // 文件地址
	CommentModel []*Comment `json:"comment_model" gorm:"-"`
	CreatedAt    time.Time  `json:"created_at" gorm:"index;comment:Create time;"` // Create time
	UpdatedAt    time.Time  `json:"updated_at" gorm:"index;comment:Update time;"` // Update time
	UploadFile   []*File    `json:"upload_file" gorm:"-"`                         // 文件
}

func (t *TaoLun) AfterFind(tx *gorm.DB) error {
	var paths []string
	json.Unmarshal([]byte(t.Path), &paths)
	if len(paths) > 0 {
		tx.Where("id in (?)", paths).Find(&t.UploadFile)
	}
	tx.Where("active_id in (?)", tx.Model(&Active{}).Where("tao_lun_id = ?", t.ID).Select("id")).Find(&t.CommentModel)
	return nil
}

// Defining the query parameters for the `TaoLun` struct.
type TaoLunQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `TaoLun` struct.
type TaoLunQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `TaoLun` struct.
type TaoLunQueryResult struct {
	Data       TaoLuns
	PageResult *util.PaginationResult
}

// Defining the slice of `TaoLun` struct.
type TaoLuns []*TaoLun

// Defining the data structure for creating a `TaoLun` struct.
type TaoLunForm struct {
	Title   string `json:"title"`   // 签到标题
	Comment string `json:"comment"` // 讨论内容
	Time    string `json:"time"`    // 讨论时长
	Path    string `json:"path"`    // 文件地址
}

// A validation function for the `TaoLunForm` struct.
func (a *TaoLunForm) Validate() error {
	return nil
}

// Convert `TaoLunForm` to `TaoLun` object.
func (a *TaoLunForm) FillTo(taoLun *TaoLun) error {
	taoLun.Comment = a.Comment
	taoLun.Path = a.Path
	taoLun.Time = a.Time
	taoLun.Title = a.Title
	taoLun.Path = a.Path
	return nil
}
