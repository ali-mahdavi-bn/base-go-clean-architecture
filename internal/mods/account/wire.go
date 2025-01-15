package account

import (
	"github.com/LyricTian/gin-admin/v10/internal/mods/account/api"
	"github.com/LyricTian/gin-admin/v10/internal/mods/account/reposetory"
	"github.com/LyricTian/gin-admin/v10/internal/mods/account/usecase"
	"github.com/google/wire"
)

// Collection of wire providers
var Set = wire.NewSet(
	wire.Struct(new(AccountModules), "*"),
	wire.Struct(new(Casbinx), "*"),
	wire.Struct(new(reposetory.Menu), "*"),
	wire.Struct(new(usecase.Menu), "*"),
	wire.Struct(new(api.Menu), "*"),
	wire.Struct(new(reposetory.MenuResource), "*"),
	wire.Struct(new(reposetory.Role), "*"),
	wire.Struct(new(usecase.Role), "*"),
	wire.Struct(new(api.Role), "*"),
	wire.Struct(new(reposetory.RoleMenu), "*"),
	wire.Struct(new(reposetory.User), "*"),
	wire.Struct(new(usecase.User), "*"),
	wire.Struct(new(api.User), "*"),
	wire.Struct(new(reposetory.UserRole), "*"),
	wire.Struct(new(usecase.Login), "*"),
	wire.Struct(new(api.Login), "*"),
)
