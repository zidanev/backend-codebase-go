package handlers

import (
	"codebase-go/bin/middlewares"
	"codebase-go/bin/modules/user"
	"codebase-go/bin/modules/user/models"
	"codebase-go/bin/pkg/errors"
	"codebase-go/bin/pkg/helpers"

	"github.com/labstack/echo"
)

type userHttpHandler struct {
	userUsecaseQuery   user.UsecaseQuery
	userUseCaseCommand user.UsecaseCommand
}

func InituserHttpHandler(e *echo.Echo, uq user.UsecaseQuery, uc user.UsecaseCommand) {

	handler := &userHttpHandler{
		userUsecaseQuery:   uq,
		userUseCaseCommand: uc,
	}

	route := e.Group("/codebase-go")

	route.GET("/users/v1", handler.Getuser, middlewares.VerifyBearer)
	route.POST("/users/v1/register", handler.RegisterUser, middlewares.VerifyBasicAuth)
	route.POST("/users/v1/login", handler.LoginUser, middlewares.VerifyBasicAuth)
	route.PUT("/users/v1/update/:id", handler.UpdateUser, middlewares.VerifyBearer)
	route.DELETE("/users/v1/delete/:id", handler.DeleteUser, middlewares.VerifyBearer)
}

func (u userHttpHandler) Getuser(c echo.Context) error {
	userId := c.Get("userId").(string)
	result := u.userUsecaseQuery.GetUser(c.Request().Context(), userId)

	if result.Error != nil {
		return helpers.RespError(c, result.Error)
	}

	return helpers.RespSuccess(c, result.Data, "Get user success")
}

func (u userHttpHandler) RegisterUser(c echo.Context) error {
	req := new(models.User)

	if err := c.Bind(req); err != nil {
		return helpers.RespError(c, errors.BadRequest("bad request."))
	}
	if err := c.Validate(req); err != nil {
		return helpers.RespError(c, err)
	}

	result := u.userUseCaseCommand.RegisterUser(c.Request().Context(), *req)
	if result.Error != nil {
		return helpers.RespError(c, result.Error)
	}

	return helpers.RespSuccess(c, result.Data, "Register user success")
}

func (u userHttpHandler) LoginUser(c echo.Context) error {
	req := new(models.LoginRequest)

	if err := c.Bind(req); err != nil {
		return helpers.RespError(c, errors.BadRequest("bad request."))
	}
	if err := c.Validate(req); err != nil {
		return helpers.RespError(c, err)
	}

	result := u.userUseCaseCommand.LoginUser(c.Request().Context(), *req)
	if result.Error != nil {
		return helpers.RespError(c, result.Error)
	}

	return helpers.RespSuccess(c, result.Data, "Register user success")
}

func (u userHttpHandler) UpdateUser(c echo.Context) error {
	req := new(models.User)
	req.Id = c.Param("id")

	if err := c.Bind(req); err != nil {
		return helpers.RespError(c, errors.BadRequest("bad request."))
	}
	if err := c.Validate(req); err != nil {
		return helpers.RespError(c, err)
	}
	if req.Id == "" {
		return helpers.RespError(c, errors.BadRequest("id can't empty."))
	}

	result := u.userUseCaseCommand.UpdateUser(c.Request().Context(), *req)
	if result.Error != nil {
		return helpers.RespError(c, result.Error)
	}

	return helpers.RespSuccess(c, result.Data, "Register user success")
}

func (u userHttpHandler) DeleteUser(c echo.Context) error {
	userId := c.Param("id")
	if userId == "" {
		return helpers.RespError(c, errors.BadRequest("id can't empty."))
	}

	result := u.userUseCaseCommand.DeleteUser(c.Request().Context(), userId)
	if result.Error != nil {
		return helpers.RespError(c, result.Error)
	}

	return helpers.RespSuccess(c, result.Data, "Delete user success")
}
