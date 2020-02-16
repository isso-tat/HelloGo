package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tatsuki5820iso/HelloGo/core/auth"
)

func Auth() (g echo.MiddlewareFunc) {
	config := middleware.JWTConfig{
		Claims:     &auth.JwtClaims{},
		SigningKey: []byte("secret"),
	}
	return middleware.JWTWithConfig(config)
}
