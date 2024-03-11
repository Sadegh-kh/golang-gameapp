package uservalidator

import "gameapp/entity"

const (
	IRPhoneNumberRegex = `^09[0-9]{9}$`
	PasswordRegex      = `^[A-Za-z0-9@#$%!*&]{8,}$`
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
