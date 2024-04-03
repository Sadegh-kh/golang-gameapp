package httpserver

import (
	"fmt"
	"gameapp/config"
	"gameapp/delivery/httpserver/userhandler"
	"gameapp/service/authservice"
	"gameapp/service/userservice"
	"gameapp/validator/uservalidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config      config.Config
	userHandler userhandler.Handler
}

func New(cfg config.Config, authS authservice.Service, userS userservice.Service, uValidator uservalidator.Validator) Server {
	return Server{
		config:      cfg,
		userHandler: userhandler.New(authS, userS, uValidator),
	}
}

func (s Server) Serve() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", s.healthCheck)

	// user routers
	s.userHandler.UserRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf("127.0.0.1:%d", s.config.HttpConf.Port)))
}
