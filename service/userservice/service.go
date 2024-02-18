package userservice

import (
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/hashpassword"
)

type Storage interface {
	Validator
	Register(u entity.User) (entity.User, error)
	CheckUserExistAndGet(phoneNumber string) (entity.User, bool, error)
}
type Validator interface {
	IsPhoneNumberUniq(phoneNumber string) (bool, error)
}

type Service struct {
	storage Storage
}
type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User entity.User `json:"user"`
}

func New(storage Storage) Service {
	return Service{storage: storage}
}
func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {

	// validator
	// uniq phone number
	// 11 len of phon number
	// more than 6 len of password
	if isUniq, err := s.storage.IsPhoneNumberUniq(req.PhoneNumber); err != nil || !isUniq {
		if err != nil {
			return RegisterResponse{}, err
		}
		if !isUniq {
			return RegisterResponse{}, fmt.Errorf("phone number is not uniq")
		}
	}
	if len(req.PhoneNumber) != 11 {
		return RegisterResponse{}, fmt.Errorf("phone number must writed by 11 number")
	}

	if len(req.Password) < 6 {
		return RegisterResponse{}, fmt.Errorf("password should be more than 6 character")
	}

	// TODO - check regex pattern password

	// hash passwrd
	passHash := hashpassword.EncodePasword(req.Password)

	// save to storage
	newUser, err := s.storage.Register(entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    passHash,
	})
	if err != nil {
		return RegisterResponse{}, err
	}

	// return new user
	return RegisterResponse{User: newUser}, nil
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	User entity.User `json:"user"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	// TODO - is better use 2 method for check user exist and get user for SOLID (S single responsibility)
	// phone is exist and get user
	user, exist, err := s.storage.CheckUserExistAndGet(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, err
	}

	// secure reason
	if !exist {
		return LoginResponse{}, fmt.Errorf("phone number or password is incorrect")
	}

	// check password
	if user.Password != hashpassword.EncodePasword(req.Password) {
		return LoginResponse{}, fmt.Errorf("phone number or password is incorrect")
	}

	return LoginResponse{User: user}, nil

}
