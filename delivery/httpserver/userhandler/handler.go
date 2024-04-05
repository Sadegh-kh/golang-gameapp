package userhandler

import (
	"gameapp/service/authservice"
	"gameapp/service/userservice"
	"gameapp/validator/uservalidator"
)

type Handler struct {
	authService   authservice.Service
	userService   userservice.Service
	userValidator uservalidator.Validator
	authCfg       authservice.Config
}

func New(authSvc authservice.Service, userSvc userservice.Service, userValid uservalidator.Validator, authCfg authservice.Config) Handler {
	return Handler{authService: authSvc, userService: userSvc, userValidator: userValid, authCfg: authCfg}
}
