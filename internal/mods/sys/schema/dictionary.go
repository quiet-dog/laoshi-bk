package schema

import (
	"time"

	"github.com/xxx/testapp/pkg/util"
)

// Dictionaries management
type Dictionary struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;"` // Unique ID
	CreatedAt time.Time `json:"created_at" gorm:"index;"`      // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`      // Update time
}

// Defining the query parameters for the `Dictionary` struct.
type DictionaryQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `Dictionary` struct.
type DictionaryQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Dictionary` struct.
type DictionaryQueryResult struct {
	Data       Dictionaries
	PageResult *util.PaginationResult
}

// Defining the slice of `Dictionary` struct.
type Dictionaries []*Dictionary

// Defining the data structure for creating a `Dictionary` struct.
type DictionaryForm struct {
}

// A validation function for the `DictionaryForm` struct.
func (a *DictionaryForm) Validate() error {
	return nil
}

// Convert `DictionaryForm` to `Dictionary` object.
func (a *DictionaryForm) FillTo(dictionary *Dictionary) error {
	return nil
}
