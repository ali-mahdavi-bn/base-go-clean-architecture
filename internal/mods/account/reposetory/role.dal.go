package reposetory

import (
	"context"

	"github.com/LyricTian/gin-admin/v10/internal/mods/account/entities"
	"github.com/LyricTian/gin-admin/v10/pkg/errors"
	"github.com/LyricTian/gin-admin/v10/pkg/util"
	"gorm.io/gorm"
)

// Get role storage instance
func GetRoleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(entities.Role))
}

// Role management for AccountModules
type Role struct {
	DB *gorm.DB
}

// Query roles from the database based on the provided parameters and options.
func (a *Role) Query(ctx context.Context, params entities.RoleQueryParam, opts ...entities.RoleQueryOptions) (*entities.RoleQueryResult, error) {
	var opt entities.RoleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetRoleDB(ctx, a.DB)
	if v := params.InIDs; len(v) > 0 {
		db = db.Where("id IN (?)", v)
	}
	if v := params.LikeName; len(v) > 0 {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.Status; len(v) > 0 {
		db = db.Where("status = ?", v)
	}
	if v := params.GtUpdatedAt; v != nil {
		db = db.Where("updated_at > ?", v)
	}

	var list entities.Roles
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &entities.RoleQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified role from the database.
func (a *Role) Get(ctx context.Context, id string, opts ...entities.RoleQueryOptions) (*entities.Role, error) {
	var opt entities.RoleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(entities.Role)
	ok, err := util.FindOne(ctx, GetRoleDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exist checks if the specified role exists in the database.
func (a *Role) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetRoleDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

func (a *Role) ExistsCode(ctx context.Context, code string) (bool, error) {
	ok, err := util.Exists(ctx, GetRoleDB(ctx, a.DB).Where("code=?", code))
	return ok, errors.WithStack(err)
}

// Create a new role.
func (a *Role) Create(ctx context.Context, item *entities.Role) error {
	result := GetRoleDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified role in the database.
func (a *Role) Update(ctx context.Context, item *entities.Role) error {
	result := GetRoleDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified role from the database.
func (a *Role) Delete(ctx context.Context, id string) error {
	result := GetRoleDB(ctx, a.DB).Where("id=?", id).Delete(new(entities.Role))
	return errors.WithStack(result.Error)
}
