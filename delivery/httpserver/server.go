package httpserver

import (
	"fmt"
	"gameapp/config"
	"gameapp/service/authservice"
	"gameapp/service/userservice"
	"gameapp/validator/uservalidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config        config.Config
	authService   authservice.Service
	userService   userservice.Service
	userValidator uservalidator.Validator
}

func New(cfg config.Config, authS authservice.Service, userS userservice.Service, uValidator uservalidator.Validator) Server {
	return Server{
		config:        cfg,
		authService:   authS,
		userService:   userS,
		userValidator: uValidator,
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
	e.Logger.Fatal(e.Start(fmt.Sprintf("127.0.0.1:%d", s.config.HttpConf.Port)))
}
