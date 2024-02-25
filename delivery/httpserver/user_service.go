package httpserver

import (
	"gameapp/service/userservice"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (s Server) userRegister(c echo.Context) error {
	var regReq userservice.RegisterRequest
	err := c.Bind(&regReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	newUser, err := s.userService.Register(regReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, newUser)
}

func (s Server) userLogin(c echo.Context) error {

	var req userservice.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	token, err := s.userService.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, token)
}

func (s Server) userProfile(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	authToken = strings.Replace(authToken, "Bearer ", "", 1)

	uid, err := s.authService.ParseToken(authToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	req := userservice.ProfileRequest{UserID: uid}

	rep, err := s.userService.Profile(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, rep)

}
