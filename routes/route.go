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

	e.POST("/register", presenter.UserPresenter.Create)
	e.POST("/login", presenter.UserPresenter.Login)

	// post type
	e.POST("/posts/categories", presenter.CategoryPresenter.Create, _middlewares.JWTMiddleware())
	e.GET("/posts/categories", presenter.CategoryPresenter.Get, _middlewares.JWTMiddleware())
	e.PUT("/posts/categories/:id", presenter.CategoryPresenter.Update, _middlewares.JWTMiddleware())
	e.DELETE("/posts/categories/:id", presenter.CategoryPresenter.Destroy, _middlewares.JWTMiddleware())
	// posting
	e.POST("/posts", presenter.PostPresenter.Create, _middlewares.JWTMiddleware())
	e.GET("/posts", presenter.PostPresenter.Get, _middlewares.JWTMiddleware())
	e.PUT("/posts/:id", presenter.PostPresenter.Update, _middlewares.JWTMiddleware())
	e.DELETE("/posts/:id", presenter.PostPresenter.Destroy, _middlewares.JWTMiddleware())

	return e
}
