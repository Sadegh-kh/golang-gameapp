package uservalidator

import (
	"gameapp/entity"
	"gameapp/pkg/richerror"
)

const (
	IRPhoneNumberRegex = `^09[0-9]{9}$`
	PasswordRegex      = `^[A-Za-z0-9@#$%!*&]{8,}$`
)

var (
	// richError for wrapped error from other layer
	richError = new(richerror.RichError)
)

type Storage interface {
	IsPhoneNumberUniq(phoneNumber string) (bool, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
}
type Validator struct {
	storage Storage
}

func New(stg Storage) Validator {
	return Validator{storage: stg}
}
