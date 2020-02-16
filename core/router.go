package core

import (
	"github.com/labstack/echo"
	"github.com/tatsuki5820iso/HelloGo/controllers"
	"net/http"
)

func (server *Server) SetRouting() {
	e := echo.New()

	// Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/test/:param", controllers.Test)

	// Users
	e.GET("/users/:id", controllers.GetUser)

	server.Echo = e
}