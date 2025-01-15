package api

import (
	"github.com/LyricTian/gin-admin/v10/internal/mods/account/entities"
	"github.com/LyricTian/gin-admin/v10/internal/mods/account/usecase"
	"github.com/LyricTian/gin-admin/v10/pkg/util"
	"github.com/gin-gonic/gin"
)

type Login struct {
	LoginUC *usecase.Login
}

// @Tags LoginAPI
// @Summary Get captcha ID
// @Success 200 {object} util.ResponseResult{data=entities.Captcha}
// @Router /api/v1/account/captcha/id [get]
func (a *Login) GetCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := a.LoginUC.GetCaptcha(ctx)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, data)
}

// @Tags LoginAPI
// @Summary Response captcha image
// @Param id query string true "Captcha ID"
// @Param reload query number false "Reload captcha image (reload=1)"
// @Produce image/png
// @Success 200 "Captcha image"
// @Failure 404 {object} util.ResponseResult
// @Router /api/v1/account/captcha/image [get]
func (a *Login) ResponseCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.LoginUC.ResponseCaptcha(ctx, c.Writer, c.Query("id"), c.Query("reload") == "1")
	if err != nil {
		util.ResError(c, err)
	}
}

// @Tags LoginAPI
// @Summary Login system with username and password
// @Param body body entities.LoginForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=entities.LoginToken}
// @Failure 400 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/account/login [post]
func (a *Login) Login(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(entities.LoginForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	}

	data, err := a.LoginUC.Login(ctx, item.Trim())
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, data)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Logout system
// @Success 200 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/account/current/logout [post]
func (a *Login) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.LoginUC.Logout(ctx)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Refresh current access token
// @Success 200 {object} util.ResponseResult{data=entities.LoginToken}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/account/current/refresh-token [post]
func (a *Login) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := a.LoginUC.RefreshToken(ctx)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, data)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Get current user info
// @Success 200 {object} util.ResponseResult{data=entities.User}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/account/current/user [get]
func (a *Login) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := a.LoginUC.GetUserInfo(ctx)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, data)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Change current user password
// @Param body body entities.UpdateLoginPassword true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/account/current/password [put]
func (a *Login) UpdatePassword(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(entities.UpdateLoginPassword)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.LoginUC.UpdatePassword(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Query current user menus based on the current user role
// @Success 200 {object} util.ResponseResult{data=[]entities.Menu}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/account/current/menus [get]
func (a *Login) QueryMenus(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := a.LoginUC.QueryMenus(ctx)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, data)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Update current user info
// @Param body body entities.UpdateCurrentUser true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/account/current/user [put]
func (a *Login) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(entities.UpdateCurrentUser)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.LoginUC.UpdateUser(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
