package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/tatsuki5820iso/HelloGo/core/auth"
	"net/http"
	"time"
)

const ExpiringHours = 72

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// TODO: Set logic to validate user.
	if username == "" || password == "" {
		return echo.ErrUnauthorized
	}

	claims := &auth.JwtClaims{
		Name: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * ExpiringHours).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
