package delivery

import (
	_helper "news/app/helper"
	"news/features/posts"
	_request "news/features/posts/delivery/request"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	articleBusiness posts.Business
}

func NewArticleHandler(dataPost posts.Business) *ArticleHandler {
	return &ArticleHandler{
		articleBusiness: dataPost,
	}
}

func (h *ArticleHandler) Create(c echo.Context) error {
	dataReq := _request.Post{}
	bind := c.Bind(&dataReq)
	if bind != nil {
		c.JSON(_helper.ResponseBadRequest("check your input data"))
	}

	dataCore := _request.ToCore(dataReq)

	response := h.articleBusiness.AddArticles(dataCore)
	if response != nil {
		c.JSON(_helper.ResponseInternalServerError("failed insert data article"))
	}
	return c.JSON(_helper.ResponseCreateSuccess("success insert data article"))
}
