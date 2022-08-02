package routes

import (
	_factory "news/app"
	_middlewares "news/app/middlewares"

	"github.com/labstack/echo/v4"
)

func New(presenter _factory.Presenter) *echo.Echo {

	e := echo.New()
	e.Pre(_middlewares.RemoveTrailingSlash())

	e.Use(_middlewares.CorsMiddleware())

	e.POST("/users", presenter.UserPresenter.Create)

	return e
}
