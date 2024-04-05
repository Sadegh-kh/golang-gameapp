package userhandler

import (
	"gameapp/param"
	"gameapp/pkg/error_converter/httpconverter"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) userLogin(c echo.Context) error {

	var req param.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return httpconverter.RaiseError(err)
	}
	err = h.userValidator.Login(req)
	if err != nil {
		return httpconverter.RaiseValidationError(err)
	}
	token, err := h.userService.Login(req)
	if err != nil {
		return httpconverter.RaiseError(err)
	}

	return c.JSON(http.StatusOK, token)
}
