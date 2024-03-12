package handler

import (
	"fmt"

	"github.com/Schierke/personal-site/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type serverCfg struct {
	Host string
	Port int
}

type SrvConfigFunc func(*serverCfg)

func WithHost(name string) func(*serverCfg) {
	return func(cfg *serverCfg) {
		cfg.Host = name
	}
}

func WithPort(port int) func(*serverCfg) {
	return func(cfg *serverCfg) {
		cfg.Port = port
	}
}

func defaultCfg() serverCfg {
	return serverCfg{
		Host: "localhost",
		Port: 3000,
	}
}

func (s serverCfg) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type Server struct {
	Router *echo.Echo
	Config serverCfg
}

func NewServer(opts ...SrvConfigFunc) *Server {
	server := &Server{
		Router: echo.New(),
		Config: defaultCfg(),
	}

	for _, opt := range opts {
		opt(&server.Config)
	}
	server.RegisterRoutes()

	return server
}

func (s Server) RegisterRoutes() {
	articleRepo := repository.NewArticleRepository()
	indexHandler := NewIndexHandler(articleRepo)
	articleHandlder := NewArticleHandler(articleRepo)
	aboutHandler := NewAboutHandler()

	indexHandler.RegisterRoutes(s.Router)
	aboutHandler.RegisterRoutes(s.Router)
	articleHandlder.RegisterRoutes(s.Router)
}

func (s Server) Start() error {
	s.Router.Use(middleware.Logger())
	return s.Router.Start(s.Config.Addr())
}
