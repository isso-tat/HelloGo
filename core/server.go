package core

import "github.com/labstack/echo"

type Server struct {
	Echo *echo.Echo
}

func NewServer () *Server {
	server := &Server{}

	server.SetRouting()

	return server
}