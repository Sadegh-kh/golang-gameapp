package userhandler

import (
	"gameapp/delivery/httpserver/middleware"
	"github.com/labstack/echo/v4"
)

func (h Handler) UserRoutes(e *echo.Echo) {
	group := e.Group("/users")

	group.POST("/register", h.userRegister)
	group.POST("/login", h.userLogin)
	group.GET("/profile", h.userProfile, middleware.Auth(h.authService, h.authCfg))
}
