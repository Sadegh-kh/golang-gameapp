package userhandler

import (
	"gameapp/param"
	"gameapp/pkg/constant"
	"gameapp/pkg/error_converter/httpconverter"
	"gameapp/service/authservice"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (h Handler) userProfile(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	authToken = strings.Replace(authToken, "Bearer ", "", 1)
	user := c.Get(constant.MiddlewareAuthJWTContext).(authservice.Claims)

	req := param.ProfileRequest{UserID: user.UserID}

	rep, err := h.userService.Profile(req)
	if err != nil {
		return httpconverter.RaiseError(err)
	}

	return c.JSON(http.StatusOK, rep)

}
