package reposetory

import (
	"context"

	"github.com/LyricTian/gin-admin/v10/internal/mods/post/entities"
	"github.com/LyricTian/gin-admin/v10/pkg/errors"
	"github.com/LyricTian/gin-admin/v10/pkg/util"
	"gorm.io/gorm"
)

// Get user storage instance
func GetPostDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(entities.Post))
}

// Post management for AccountModules
type Post struct {
	DB *gorm.DB
}

// Query users from the database based on the provided parameters and options.
func (a *Post) Query(ctx context.Context, params entities.PostQueryParam, opts ...entities.PostQueryOptions) (*entities.PostQueryResult, error) {
	var opt entities.PostQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetPostDB(ctx, a.DB)
	if v := params.LikeName; len(v) > 0 {
		db = db.Where("username LIKE ?", "%"+v+"%")
	}

	var list entities.Posts
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &entities.PostQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified user from the database.
func (a *Post) Get(ctx context.Context, id string, opts ...entities.PostQueryOptions) (*entities.Post, error) {
	var opt entities.PostQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(entities.Post)
	ok, err := util.FindOne(ctx, GetPostDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

func (a *Post) GetByPostName(ctx context.Context, username string, opts ...entities.PostQueryOptions) (*entities.Post, error) {
	var opt entities.PostQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(entities.Post)
	ok, err := util.FindOne(ctx, GetPostDB(ctx, a.DB).Where("username=?", username), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exist checks if the specified user exists in the database.
func (a *Post) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetPostDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

func (a *Post) ExistsPostName(ctx context.Context, username string) (bool, error) {
	ok, err := util.Exists(ctx, GetPostDB(ctx, a.DB).Where("username=?", username))
	return ok, errors.WithStack(err)
}

// Create a new user.
func (a *Post) Create(ctx context.Context, item *entities.Post) error {
	result := GetPostDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified user in the database.
func (a *Post) Update(ctx context.Context, item *entities.Post, selectFields ...string) error {
	db := GetPostDB(ctx, a.DB).Where("id=?", item.ID)
	if len(selectFields) > 0 {
		db = db.Select(selectFields)
	} else {
		db = db.Select("*").Omit("created_at")
	}
	result := db.Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified user from the database.
func (a *Post) Delete(ctx context.Context, id string) error {
	result := GetPostDB(ctx, a.DB).Where("id=?", id).Delete(new(entities.Post))
	return errors.WithStack(result.Error)
}
