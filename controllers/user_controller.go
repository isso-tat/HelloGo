package controllers

import (
	"github.com/labstack/echo"
	"github.com/tatsuki5820iso/HelloGo/models"
	"net/http"
)

func GetUser (c echo.Context) error {
	var user = models.User{
		Id: c.Param("id"),
		Name: "Mr.TestUser",
	}
	return c.JSON(http.StatusOK, user)
}