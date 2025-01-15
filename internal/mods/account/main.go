package account

import (
	"context"
	"github.com/LyricTian/gin-admin/v10/internal/mods/account/api"
	"path/filepath"

	"github.com/LyricTian/gin-admin/v10/internal/config"
	"github.com/LyricTian/gin-admin/v10/internal/mods/account/entities"
	"github.com/LyricTian/gin-admin/v10/pkg/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AccountModules struct {
	DB       *gorm.DB
	MenuAPI  *api.Menu
	RoleAPI  *api.Role
	UserAPI  *api.User
	LoginAPI *api.Login
	Casbinx  *Casbinx
}

func (a *AccountModules) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(
		new(entities.Menu),
		new(entities.MenuResource),
		new(entities.Role),
		new(entities.RoleMenu),
		new(entities.User),
		new(entities.UserRole),
	)
}

func (a *AccountModules) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}

	if err := a.Casbinx.Load(ctx); err != nil {
		return err
	}

	if name := config.C.General.MenuFile; name != "" {
		fullPath := filepath.Join(config.C.General.WorkDir, name)
		if err := a.MenuAPI.MenuUC.InitFromFile(ctx, fullPath); err != nil {
			logging.Context(ctx).Error("failed to init menu data", zap.Error(err), zap.String("file", fullPath))
		}
	}

	return nil
}

func (a *AccountModules) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	public := v1.Group("p")
	admin := v1.Group("admin")
	account := v1.Group("account")
	captcha := public.Group("captcha")
	{
		captcha.GET("id", a.LoginAPI.GetCaptcha)
		captcha.GET("image", a.LoginAPI.ResponseCaptcha)
	}

	public.POST("login", a.LoginAPI.Login)

	current := v1.Group("current")
	{
		current.POST("refresh-token", a.LoginAPI.RefreshToken)
		current.GET("user", a.LoginAPI.GetUserInfo)
		current.GET("menus", a.LoginAPI.QueryMenus)
		current.PUT("password", a.LoginAPI.UpdatePassword)
		current.PUT("user", a.LoginAPI.UpdateUser)
		current.POST("logout", a.LoginAPI.Logout)
	}

	user := admin.Group("users")
	{
		user.GET("", a.UserAPI.Query)
		user.GET(":id", a.UserAPI.Get)
		user.POST("", a.UserAPI.Create)
		user.PUT(":id", a.UserAPI.Update)
		user.DELETE(":id", a.UserAPI.Delete)
		user.PATCH(":id/reset-pwd", a.UserAPI.ResetPassword)
	}

	menu := account.Group("menus")
	{
		menu.GET("", a.MenuAPI.Query)
		menu.GET(":id", a.MenuAPI.Get)
		menu.POST("", a.MenuAPI.Create)
		menu.PUT(":id", a.MenuAPI.Update)
		menu.DELETE(":id", a.MenuAPI.Delete)
	}
	pmenu := public.Group("menus")
	{
		pmenu.GET("", a.MenuAPI.PQuery)
	}

	role := admin.Group("roles")
	{
		role.GET("", a.RoleAPI.Query)
		role.GET(":id", a.RoleAPI.Get)
		role.POST("", a.RoleAPI.Create)
		role.PUT(":id", a.RoleAPI.Update)
		role.DELETE(":id", a.RoleAPI.Delete)
	}

	return nil
}

func (a *AccountModules) Release(ctx context.Context) error {
	if err := a.Casbinx.Release(ctx); err != nil {
		return err
	}
	return nil
}
