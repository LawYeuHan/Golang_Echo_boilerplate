package handler

import (
	"ecpos/internal/service"
	"ecpos/pkg/helper/resp"
	"ecpos/pkg/log"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetUserByID(ctx echo.Context) error
	GetUserByIDWithError(ctx echo.Context) error
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(us service.UserService) UserHandler {
	return &userHandler{us}
}

func (uh *userHandler) GetUserByID(ctx echo.Context) error {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		//log.Logger().Error("user id should be a number")
		return resp.HandleError(ctx, http.StatusBadRequest, 1, "invalid id", nil, err)
	}

	user, err := uh.userService.GetUserByID(id)
	if err != nil {
		return resp.HandleError(ctx, http.StatusInternalServerError, log.GetErrCode(err), "user not found", nil, err)
	}
	return resp.HandleSuccess(ctx, user)
}

func (uh *userHandler) GetUserByIDWithError(ctx echo.Context) error {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		//log.Logger().Error("user id should be a number")
		return resp.HandleError(ctx, http.StatusBadRequest, log.GetErrCode(err), "invalid id", nil, err)
	}

	user, err := uh.userService.GetUserByIDWithError(id)
	if err != nil {
		if errors.Is(err, log.ErrNotFound) {
			return resp.HandleError(ctx, http.StatusBadRequest, log.GetErrCode(err), "user not found", nil, err)
		}
		return resp.HandleError(ctx, http.StatusInternalServerError, log.GetErrCode(err), "something went wrong", nil, err)
	}
	return resp.HandleSuccess(ctx, user)
}
