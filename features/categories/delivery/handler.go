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
		c.JSON(_helper.ResponseBadRequest("check your input data"))
	}

	dataCore := _request.ToCore(dataReq)

	response := h.categoryBusiness.AddCategories(dataCore)
	if response != nil {
		c.JSON(_helper.ResponseBadRequest("failed insert data post type"))
	}
	return c.JSON(_helper.ResponseCreateSuccess("success insert data post type"))
}

func (h *CategoryHandler) Get(c echo.Context) error {
	response, err := h.categoryBusiness.GetCategories()
	if err != nil {
		c.JSON(_helper.ResponseBadRequest("failed get data post type"))
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
		c.JSON(_helper.ResponseBadRequest("check your input data"))
	}

	dataCore := _request.ToCore(dataReq)

	response := h.categoryBusiness.UpdateCategories(dataCore, categoryID)
	if response != nil {
		c.JSON(_helper.ResponseBadRequest("failed update data post type"))
	}
	return c.JSON(_helper.ResponseStatusOkNoData("success update data post type"))
}

func (h *CategoryHandler) Destroy(c echo.Context) error {
	categoryID, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		return c.JSON(_helper.ResponseBadRequest("error parameter"))
	}

	response := h.categoryBusiness.DestroyCategories(categoryID)
	if response != nil {
		c.JSON(_helper.ResponseBadRequest("failed delete data post type"))
	}
	return c.JSON(_helper.ResponseStatusOkNoData("success delete data post type"))
}
