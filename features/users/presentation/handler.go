package presentation

import (
	_helper "news/app/helper"
	_request "news/features/presentation/request"
	"news/features/users"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(dataUser users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: dataUser,
	}
}

func (h *UserHandler) Create(c echo.Context) error {
	dataReq := _request.User{}
	bind := c.Bind(&dataReq)
	if bind != nil {
		c.JSON(_helper.ResponseBadRequest("check your input data"))
	}

	dataCore := _request.ToCore(dataReq)

	response := h.userBusiness.AddUser(dataCore)
	if response != nil {
		c.JSON(_helper.ResponseInternalServerError("failed insert data user"))
	}
	return c.JSON(_helper.ResponseCreateSuccess("success insert data user"))
}
