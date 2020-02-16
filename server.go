package main

import (
	"github.com/labstack/echo"
	"github.com/tatsuki5820iso/HelloGo/controllers"
	"net/http"
)

func main() {
	e := echo.New()

	// Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/test/:param", controllers.Test)
	e.GET("/users/:id", controllers.GetUser)

	e.Logger.Fatal(e.Start(":8000"))
}