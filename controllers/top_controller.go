package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func Test (c echo.Context) error {
	// Return user param.
	param := c.Param("param")
	foo := c.QueryParam("foo")
	res := fmt.Sprintf("param: %s, foo: %s", param, foo)
	return c.String(http.StatusOK, res)
}
