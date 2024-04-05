package userhandler

import (
	"gameapp/param"
	"gameapp/pkg/error_converter/httpconverter"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) userRegister(c echo.Context) error {
	var regReq param.RegisterRequest
	err := c.Bind(&regReq)
	if err != nil {
		return httpconverter.RaiseError(err)
	}

	//validation layer
	err = h.userValidator.Register(regReq)
	if err != nil {
		return httpconverter.RaiseValidationError(err)
	}

	newUser, err := h.userService.Register(regReq)
	if err != nil {
		return httpconverter.RaiseError(err)
	}
	return c.JSON(http.StatusCreated, newUser)
}
