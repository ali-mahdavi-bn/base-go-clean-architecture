package entities

import (
	"github.com/LyricTian/gin-admin/v10/internal/config"
	"github.com/LyricTian/gin-admin/v10/pkg/util"
	"time"
)

type Post struct {
	ID          string `json:"id" gorm:"size:20;primarykey;"`
	Name        string `json:"name" gorm:"size:64;index"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	UserID      uint
	CreatedAt   time.Time `json:"created_at" gorm:"index;"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"index;"`
}

func (a *Post) TableName() string {
	return config.C.FormatTableName("post")
}

// Defining the query parameters for the `Post` struct.
type PostQueryParam struct {
	util.PaginationParam
	LikeName string `form:"name"`
}

// Defining the query options for the `Post` struct.
type PostQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Post` struct.
type PostQueryResult struct {
	Data       Posts
	PageResult *util.PaginationResult
}

// Defining the slice of `Post` struct.
type Posts []*Post

func (a Posts) ToIDs() []string {
	var ids []string
	for _, item := range a {
		ids = append(ids, item.ID)
	}
	return ids
}

// Defining the data structure for creating a `Post` struct.
type PostForm struct {
	ID          string `json:"id" gorm:"size:20;primarykey;"`
	Name        string `json:"name" gorm:"size:64;index"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	UserID      uint
}

// A validation function for the `UserForm` struct.
func (a *PostForm) Validate() error {
	return nil
}

// Convert `UserForm` to `Post` object.
func (a *PostForm) FillTo(user *Post) error {
	user.Name = a.Name
	return nil
}
