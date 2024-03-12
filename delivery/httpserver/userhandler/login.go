package userhandler

import (
	"errors"
	"gameapp/param"
	"gameapp/pkg/error_converter/httpconverter"
	"gameapp/pkg/richerror"
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
		var rErr richerror.RichError
		errors.As(err, &rErr)
		return echo.NewHTTPError(http.StatusBadRequest, rErr.ValidationErrors)
	}
	token, err := h.userService.Login(req)
	if err != nil {
		return httpconverter.RaiseError(err)
	}

	return c.JSON(http.StatusOK, token)
}
