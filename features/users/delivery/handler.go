package presentation

import (
	_helper "news/app/helper"
	"news/features/users"
	_request "news/features/users/delivery/request"
	_response "news/features/users/delivery/response"

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
		return c.JSON(_helper.ResponseBadRequest(response.Error()))
	}
	return c.JSON(_helper.ResponseCreateSuccess("success insert data user"))
}

func (h *UserHandler) Login(c echo.Context) error {
	dataReq := _request.User{}
	errBind := c.Bind(&dataReq)
	if errBind != nil {
		c.JSON(_helper.ResponseBadRequest("check your input data"))
	}

	dataCore := _request.ToCore(dataReq)

	token, fullName, userID, err := h.userBusiness.Auth(dataCore)
	if err != nil {
		return c.JSON(_helper.ResponseNotFound(err.Error()))
	}

	result := _response.ToResponse(token, fullName, userID)

	return c.JSON(_helper.ResponseStatusOkWithData("login success", result))

}
