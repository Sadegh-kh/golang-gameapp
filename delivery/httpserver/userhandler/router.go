package userhandler

import "github.com/labstack/echo/v4"

func (h Handler) UserRouter(e *echo.Echo) {
	group := e.Group("/users")

	group.POST("/register", h.userRegister)
	group.POST("/login", h.userLogin)
	group.GET("/profile", h.userProfile)
}
