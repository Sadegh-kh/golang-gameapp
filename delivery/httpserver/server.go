package httpserver

import (
	"fmt"
	"gameapp/config"
	"gameapp/service/authservice"
	"gameapp/service/userservice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config      config.Config
	authService authservice.Service
	userService userservice.Service
}

func New(config config.Config, authService authservice.Service, userService userservice.Service) Server {
	return Server{
		config:      config,
		authService: authService,
		userService: userService,
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
	e.POST("/register", s.userRegister)
	e.POST("/login", s.userLogin)
	e.GET("/profile", s.userProfile)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HttpConf.Port)))
}
