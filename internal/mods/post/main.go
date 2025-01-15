package account

import (
	"context"
	"github.com/LyricTian/gin-admin/v10/internal/config"
	"github.com/LyricTian/gin-admin/v10/internal/mods/post/api"
	"github.com/LyricTian/gin-admin/v10/internal/mods/post/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Post struct {
	DB      *gorm.DB
	PostAPI *api.Post
}

func (a *Post) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(
		new(entities.Post),
	)
}

func (a *Post) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *Post) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	puPost := v1.Group("p/post")
	prPost := v1.Group("/post")
	{
		// public
		puPost.GET("", a.PostAPI.Query)
		puPost.GET(":id", a.PostAPI.Get)
		// private
		prPost.POST("", a.PostAPI.Create)
		prPost.PUT(":id", a.PostAPI.Update)
	}
	return nil
}
