package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
	"gorm.io/gorm"
)

// 活动
type Active struct {
	ID          string    `json:"id" gorm:"size:20;primaryKey;comment:Unique ID;"` // Unique ID
	SignId      string    `json:"sign_id" gorm:"size:512;comment:签到ID;"`           // 签到ID
	IsEnd       bool      `json:"is_end" gorm:"comment:是否结束;"`                     // 是否结束
	IsStart     bool      `json:"is_start" gorm:"comment:是否开始;"`                   // 是否开始
	CreatedAt   time.Time `json:"created_at" gorm:"index;comment:Create time;"`    // Create time
	UpdatedAt   time.Time `json:"updated_at" gorm:"index;comment:Update time;"`    // Update time
	SignModel   *Sign     `json:"sign_model" gorm:"foreignKey:SignId"`             // 签到
	PkId        string    `json:"pk_id" gorm:"size:20;comment:pkID;"`              // pkID
	PkModel     *Pk       `json:"pk_model" gorm:"foreignKey:PkId"`                 // pk
	TaoLunId    string    `json:"tao_lun_id" gorm:"size:20;comment:taoLunId;"`
	TaoLunModel *TaoLun   `json:"tao_lun_model" gorm:"foreignKey:TaoLunId"` // pk
	CanYu       []*Employ `json:"can_yu" gorm:"-"`
}

func (a *Active) AfterFind(tx *gorm.DB) error {
	if a.SignId != "" {
		tx.Where("id = ?", a.SignId).First(&a.SignModel)
		tx.Where("id in (?)", tx.Model(&SignLog{}).Where("active_id = ?", a.ID).Select("employ_id")).Find(&a.CanYu)
	}
	if a.PkId != "" {
		tx.Where("id = ?", a.PkId).First(&a.PkModel)
		tx.Where("id in (?)", tx.Model(&PkLog{}).Where("active_id = ?", a.ID).Select("employ_id")).Find(&a.CanYu)
	}
	if a.TaoLunId != "" {
		tx.Where("id = ?", a.TaoLunId).First(&a.TaoLunModel)
		tx.Where("id in (?)", tx.Model(&Comment{}).Group("employ_id").Distinct("employ_id").Where("active_id = ?", a.ID).Select("employ_id")).Find(&a.CanYu)
	}

	return nil
}

// Defining the query parameters for the `Active` struct.
type ActiveQueryParam struct {
	util.PaginationParam

	SignId  string `form:"-"` // 签到ID
	IsEnd   bool   `form:"-"` // 是否结束
	IsStart bool   `form:"-"` // 是否开始
}

// Defining the query options for the `Active` struct.
type ActiveQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Active` struct.
type ActiveQueryResult struct {
	Data       Actives
	PageResult *util.PaginationResult
}

// Defining the slice of `Active` struct.
type Actives []*Active

// Defining the data structure for creating a `Active` struct.
type ActiveForm struct {
	SignId   string `json:"sign_id"`
	IsEnd    bool   `json:"is_end"`
	IsStart  bool   `json:"is_start"`
	PkId     string `json:"pk_id"`
	TaoLunId string `json:"tao_lun_id"`
}

// A validation function for the `ActiveForm` struct.
func (a *ActiveForm) Validate() error {
	return nil
}

// Convert `ActiveForm` to `Active` object.
func (a *ActiveForm) FillTo(active *Active) error {
	active.SignId = a.SignId
	active.IsEnd = a.IsEnd
	active.IsStart = a.IsStart
	active.PkId = a.PkId
	active.TaoLunId = a.TaoLunId
	return nil
}
