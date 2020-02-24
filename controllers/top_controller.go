package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func TestGet(c echo.Context) error {
	// Return user param.
	param := c.Param("param")
	foo := c.QueryParam("foo")
	res := fmt.Sprintf("params: %s, foo: %s", param, foo)
	return c.String(http.StatusOK, res)
}

func TestPost(c echo.Context) error {
	param := c.FormValue("param")
	res := fmt.Sprintf("param: %s", param)
	return c.String(http.StatusOK, res)
}
