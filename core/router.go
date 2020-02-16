package core

import (
	"github.com/labstack/echo"
	"github.com/tatsuki5820iso/HelloGo/controllers"
	"net/http"
)

func (server *Server) SetRouting() {
	e := echo.New()

	// SetMiddleware middleware.
	// Also get authenticate require route group
	authGroup := SetMiddleware(e)

	// Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/test/:param", controllers.TestGet)
	e.POST("/test", controllers.TestPost)

	// Authentication
	e.POST("/login", controllers.Login)

	// Needs authentication
	// The authenticated user.
	authGroup.GET("/me", controllers.GetMe)

	// Users
	authGroup.GET("/users/:id", controllers.GetUser)

	server.Echo = e
}


