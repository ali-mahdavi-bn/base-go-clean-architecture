package usecase

import (
	"context"
	"github.com/LyricTian/gin-admin/v10/internal/mods/post/reposetory"
	"time"

	"github.com/LyricTian/gin-admin/v10/internal/config"
	"github.com/LyricTian/gin-admin/v10/internal/mods/post/entities"
	"github.com/LyricTian/gin-admin/v10/pkg/cachex"
	"github.com/LyricTian/gin-admin/v10/pkg/errors"
	"github.com/LyricTian/gin-admin/v10/pkg/util"
)

// Post management for AccountModules
type Post struct {
	Cache   cachex.Cacher
	Trans   *util.Trans
	PostDAL *reposetory.Post
}

// Query users from the data access object based on the provided parameters and options.
func (a *Post) Query(ctx context.Context, params entities.PostQueryParam) (*entities.PostQueryResult, error) {
	params.Pagination = true

	result, err := a.PostDAL.Query(ctx, params, entities.PostQueryOptions{
		QueryOptions: util.QueryOptions{
			OrderFields: []util.OrderByParam{
				{Field: "created_at", Direction: util.DESC},
			},
			OmitFields: []string{"password"},
		},
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get the specified user from the data access object.
func (a *Post) Get(ctx context.Context, id string) (*entities.Post, error) {
	user, err := a.PostDAL.Get(ctx, id, entities.PostQueryOptions{
		QueryOptions: util.QueryOptions{
			OmitFields: []string{"password"},
		},
	})
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, errors.NotFound("", "Post not found")
	}

	return user, nil
}

// Create a new user in the data access object.
func (a *Post) Create(ctx context.Context, formItem *entities.PostForm) (*entities.Post, error) {
	existsPostname, err := a.PostDAL.ExistsPostName(ctx, formItem.Name)
	if err != nil {
		return nil, err
	} else if existsPostname {
		return nil, errors.BadRequest("", "Postname already exists")
	}

	user := &entities.Post{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if formItem.Password == "" {
		formItem.Password = config.C.General.DefaultLoginPwd
	}

	if err := formItem.FillTo(user); err != nil {
		return nil, err
	}
	return user, nil
}

// Update the specified user in the data access object.
func (a *Post) Update(ctx context.Context, id string, formItem *entities.PostForm) error {
	user, err := a.PostDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if user == nil {
		return errors.NotFound("", "Post not found")
	} else if user.Name != formItem.Name {
		existsPostname, err := a.PostDAL.ExistsPostName(ctx, formItem.Name)
		if err != nil {
			return err
		} else if existsPostname {
			return errors.BadRequest("", "Postname already exists")
		}
	}

	if err := formItem.FillTo(user); err != nil {
		return err
	}
	user.UpdatedAt = time.Now()
	return nil
}
