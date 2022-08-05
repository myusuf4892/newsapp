package delivery

import (
	_helper "news/app/helper"
	"news/features/categories"
	_request "news/features/categories/delivery/request"
	_response "news/features/categories/delivery/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryBusiness categories.Business
}

func NewCategoryHandler(dataPostType categories.Business) *CategoryHandler {
	return &CategoryHandler{
		categoryBusiness: dataPostType,
	}
}

func (h *CategoryHandler) Create(c echo.Context) error {
	dataReq := _request.PostType{}
	errBind := c.Bind(&dataReq)
	if errBind != nil {
		return c.JSON(_helper.ResponseBadRequest("check your input data"))
	}

	dataCore := _request.ToCore(dataReq)

	row, err := h.categoryBusiness.AddCategories(dataCore)
	if err != nil {
		return c.JSON(_helper.ResponseInternalServerError(err.Error()))
	}
	if row == 0 {
		return c.JSON(_helper.ResponseBadRequest("insert post type failed"))
	}
	return c.JSON(_helper.ResponseCreateSuccess("insert post type success"))
}

func (h *CategoryHandler) Get(c echo.Context) error {
	response, err := h.categoryBusiness.GetCategories()
	if err != nil {
		return c.JSON(_helper.ResponseBadRequest("failed get data post type"))
	}

	result := _response.FromCoreToList(response)
	return c.JSON(_helper.ResponseStatusOkWithData("success get data post type", result))
}

func (h *CategoryHandler) Update(c echo.Context) error {
	categoryID, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		return c.JSON(_helper.ResponseBadRequest("error parameter"))
	}

	dataReq := _request.PostType{}
	errBind := c.Bind(&dataReq)
	if errBind != nil {
		return c.JSON(_helper.ResponseBadRequest("check your input data"))
	}

	dataCore := _request.ToCore(dataReq)

	row, err := h.categoryBusiness.UpdateCategories(dataCore, categoryID)
	if err != nil {
		return c.JSON(_helper.ResponseInternalServerError(err.Error()))
	}
	if row == 0 {
		return c.JSON(_helper.ResponseBadRequest("no data, update failed"))
	}
	return c.JSON(_helper.ResponseStatusOkNoData("updated post type success"))
}

func (h *CategoryHandler) Destroy(c echo.Context) error {
	categoryID, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		return c.JSON(_helper.ResponseBadRequest("error parameter"))
	}

	row, err := h.categoryBusiness.DestroyCategories(categoryID)
	if err != nil {
		return c.JSON(_helper.ResponseInternalServerError(err.Error()))
	}
	if row == 0 {
		return c.JSON(_helper.ResponseBadRequest("no data, delete failed"))
	}
	return c.JSON(_helper.ResponseStatusOkNoData("delete post type success"))
}
