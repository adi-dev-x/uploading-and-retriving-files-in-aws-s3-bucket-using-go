package boot

import "github.com/labstack/echo"

type Server struct {
	Engine *echo.Echo
}

func (s *Server) Start() {
	s.Engine.Start("localhost:8080")
}

func NewHTTPServer() *Server {
	e := echo.New()
	return &Server{
		Engine: e,
	}
}
