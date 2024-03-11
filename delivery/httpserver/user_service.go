package httpserver

import (
	"errors"
	"gameapp/dto"
	"gameapp/pkg/richerror"
	"gameapp/service/userservice"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (s Server) userRegister(c echo.Context) error {
	var regReq dto.RegisterRequest
	err := c.Bind(&regReq)
	if err != nil {
		return raiseError(err)
	}

	//validation layer
	err = s.userValidator.Register(regReq)
	if err != nil {
		var rErrs richerror.RichError
		errors.As(err, &rErrs)
		return echo.NewHTTPError(http.StatusBadRequest, rErrs.ValidationErrors)
	}

	newUser, err := s.userService.Register(regReq)
	if err != nil {
		return raiseError(err)
	}
	return c.JSON(http.StatusCreated, newUser)
}

func (s Server) userLogin(c echo.Context) error {

	var req dto.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return raiseError(err)
	}
	err = s.userValidator.Login(req)
	if err != nil {
		var rErr richerror.RichError
		errors.As(err, &rErr)
		return echo.NewHTTPError(http.StatusBadRequest, rErr.ValidationErrors)
	}
	token, err := s.userService.Login(req)
	if err != nil {
		return raiseError(err)
	}

	return c.JSON(http.StatusOK, token)
}

func (s Server) userProfile(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	authToken = strings.Replace(authToken, "Bearer ", "", 1)

	uid, err := s.authService.ParseToken(authToken)
	if err != nil {
		return raiseError(err)
	}
	req := userservice.ProfileRequest{UserID: uid}

	rep, err := s.userService.Profile(req)
	if err != nil {
		return raiseError(err)
	}

	return c.JSON(http.StatusOK, rep)

}
