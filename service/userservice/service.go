package userservice

import (
	"gameapp/entity"
)

type Storage interface {
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
	GetUserByID(uid uint) (entity.User, error)
}

type authService interface {
	CreateAccessToken(uid uint) (string, error)
	CreateRefreshToken(uid uint) (string, error)
}

type Service struct {
	storage Storage
	auth    authService
}

func New(stg Storage, authS authService) Service {
	return Service{storage: stg, auth: authS}
}
