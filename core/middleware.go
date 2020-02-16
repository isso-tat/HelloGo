package core

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tatsuki5820iso/HelloGo/core/middlewares"
)

const AuthPrefix = ""

func SetMiddleware(e *echo.Echo) (authGroup *echo.Group) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Authentication
	authGroup = e.Group(AuthPrefix)
	authGroup.Use(middlewares.Auth())

	return
}
