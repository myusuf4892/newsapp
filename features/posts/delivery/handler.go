package delivery

import (
	"fmt"
	_helper "news/app/helper"
	_middlewares "news/app/middlewares"
	"news/features/posts"
	_request "news/features/posts/delivery/request"
	_response "news/features/posts/delivery/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	postBusiness posts.Business
}

func NewPostHandler(dataPost posts.Business) *PostHandler {
	return &PostHandler{
		postBusiness: dataPost,
	}
}

func (h *PostHandler) Create(c echo.Context) error {
	tokenUserID, errToken := _middlewares.ExtractToken(c)
	if tokenUserID == 0 || errToken != nil {
		return c.JSON(_helper.ResponseBadRequest("failed to get user id"))
	}

	dataReq := _request.Post{}
	errBind := c.Bind(&dataReq)
	if errBind != nil {
		return c.JSON(_helper.ResponseBadRequest("check your input data"))
	}

	dataCore := _request.ToCore(dataReq)
	dataCore.UserID = tokenUserID

	fmt.Println(dataCore.PostType)

	row, err := h.postBusiness.AddPost(dataCore)
	if err != nil {
		return c.JSON(_helper.ResponseInternalServerError(err.Error()))
	}
	if row == 0 {
		return c.JSON(_helper.ResponseBadRequest("insert post failed"))
	}
	return c.JSON(_helper.ResponseCreateSuccess("success insert data post"))
}

func (h *PostHandler) Get(c echo.Context) error {
	response, err := h.postBusiness.GetPost()
	if err != nil {
		return c.JSON(_helper.ResponseBadRequest("failed get data post"))
	}

	result := _response.FromCoreToList(response)
	return c.JSON(_helper.ResponseStatusOkWithData("success get data post", result))
}

func (h *PostHandler) Update(c echo.Context) error {
	tokenUserID, errToken := _middlewares.ExtractToken(c)
	if tokenUserID == 0 || errToken != nil {
		return c.JSON(_helper.ResponseBadRequest("failed to get user id"))
	}

	categoryID, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		return c.JSON(_helper.ResponseBadRequest("error parameter"))
	}

	dataReq := _request.Post{}
	errBind := c.Bind(&dataReq)
	if errBind != nil {
		return c.JSON(_helper.ResponseBadRequest("check your input data"))
	}

	dataCore := _request.ToCore(dataReq)
	dataCore.UserID = tokenUserID

	row, err := h.postBusiness.UpdatePost(dataCore, categoryID)
	if err != nil {
		return c.JSON(_helper.ResponseInternalServerError(err.Error()))
	}
	if row == 0 {
		return c.JSON(_helper.ResponseBadRequest("no data, update post failed"))
	}
	return c.JSON(_helper.ResponseStatusOkNoData("success update data post"))
}

func (h *PostHandler) Destroy(c echo.Context) error {
	categoryID, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		return c.JSON(_helper.ResponseBadRequest("error parameter"))
	}

	row, err := h.postBusiness.DestroyPost(categoryID)
	if err != nil {
		return c.JSON(_helper.ResponseInternalServerError(err.Error()))
	}
	if row == 0 {
		return c.JSON(_helper.ResponseBadRequest("no data, delete post failed"))
	}
	return c.JSON(_helper.ResponseStatusOkNoData("success delete data post"))
}
