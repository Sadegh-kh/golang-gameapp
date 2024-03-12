package userservice

import (
	"gameapp/entity"
	"gameapp/param"
	"gameapp/pkg/hashpassword"
	"gameapp/pkg/richerror"
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
func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {

	passHash := hashpassword.EncodePasword(req.Password)

	newUser, err := s.storage.Register(entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    passHash,
	})
	if err != nil {
		return param.RegisterResponse{}, richerror.RichError{
			Operation:    "userservice.Register",
			WrappedError: err,
			Message:      "unexpected error",
			Kind:         richerror.Unexpected,
			Meta: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}

	return param.RegisterResponse{ID: newUser.ID, Name: newUser.Name, PhoneNumber: newUser.PhoneNumber}, nil
}

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	// TODO - is better use 2 method for check user exist and get user for SOLID (S single responsibility)

	user, err := s.storage.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return param.LoginResponse{}, richerror.RichError{
			Operation:    "userservice.Login",
			WrappedError: nil,
			Message:      "phone number or password is incorrect",
			Kind:         richerror.Invalid,
			Meta: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}

	// secure reason
	//TODO-add to login validator

	// check password
	if user.Password != hashpassword.EncodePasword(req.Password) {
		return param.LoginResponse{}, richerror.RichError{
			Operation:    "userservice.Login",
			WrappedError: nil,
			Message:      "phone number or password is incorrect",
			Kind:         richerror.Invalid,
			Meta:         nil,
		}
	}

	// method 1
	// we can generate session and send session ID to user
	// with that session ID we can authentication and authorization user

	// method 2
	// jwt token
	accessToken, err := s.auth.CreateAccessToken(user.ID)
	if err != nil {
		return param.LoginResponse{}, richerror.RichError{
			Operation:    "userservice.Register",
			WrappedError: err,
			Message:      "unexpected error",
			Kind:         richerror.Unexpected,
			Meta: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}

	refreshToken, err := s.auth.CreateRefreshToken(user.ID)
	if err != nil {
		return param.LoginResponse{}, richerror.RichError{
			Operation:    "userservice.Register",
			WrappedError: err,
			Message:      "unexpected error",
			Kind:         richerror.Unexpected,
			Meta: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}

	return param.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil

}

func (s Service) Profile(req param.ProfileRequest) (param.ProfileResponse, error) {
	user, err := s.storage.GetUserByID(req.UserID)
	if err != nil {
		return param.ProfileResponse{}, err
	}

	return param.ProfileResponse{Name: user.Name}, nil

}
