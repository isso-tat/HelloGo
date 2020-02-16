package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/tatsuki5820iso/HelloGo/core/auth"
	"net/http"
)

func GetMe(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JwtClaims)
	return c.JSON(http.StatusOK, claims)
}
