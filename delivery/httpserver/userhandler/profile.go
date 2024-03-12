package userhandler

import (
	"gameapp/pkg/error_converter/httpconverter"
	"gameapp/service/userservice"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (h Handler) userProfile(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	authToken = strings.Replace(authToken, "Bearer ", "", 1)

	uid, err := h.authService.ParseToken(authToken)
	if err != nil {
		return httpconverter.RaiseError(err)
	}
	req := userservice.ProfileRequest{UserID: uid}

	rep, err := h.userService.Profile(req)
	if err != nil {
		return httpconverter.RaiseError(err)
	}

	return c.JSON(http.StatusOK, rep)

}
