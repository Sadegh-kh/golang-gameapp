package userservice

import (
	"gameapp/dto"
	"gameapp/entity"
	"gameapp/pkg/hashpassword"
	"gameapp/pkg/richerror"
)

type Storage interface {
	Validator
	Register(u entity.User) (entity.User, error)
	CheckUserExistAndGet(phoneNumber string) (entity.User, bool, error)
	GetUserByID(uid uint) (entity.User, error)
}

type Validator interface {
	IsPhoneNumberUniq(phoneNumber string) (bool, error)
}

type authService interface {
	CreateAccessToken(uid uint) (string, error)
	CreateRefreshToken(uid uint) (string, error)
}

type Service struct {
	storage Storage
	auth    authService
}

func New(storage Storage, authService authService) Service {
	return Service{storage: storage, auth: authService}
}
func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {

	// validator
	// uniq phone number
	// 11 len of phone number
	// more than 6 len of password
	if isUniq, err := s.storage.IsPhoneNumberUniq(req.PhoneNumber); err != nil || !isUniq {
		if err != nil {
			return dto.RegisterResponse{}, richerror.RichError{
				Operation:    "userservice.Register",
				WrappedError: err,
				Message:      "unexpected error",
				Kind:         richerror.Unexpected,
				Meta: map[string]interface{}{
					"message": err.Error(),
				},
			}
		}
		if !isUniq {
			return dto.RegisterResponse{}, richerror.RichError{
				Operation:    "userservice.Register",
				WrappedError: nil,
				Message:      "phone number is not uniq",
				Kind:         richerror.Invalid,
				Meta:         nil,
			}
		}
	}

	if len(req.PhoneNumber) != 11 {
		return dto.RegisterResponse{}, richerror.RichError{
			Operation:    "userservice.Register",
			WrappedError: nil,
			Message:      "phone number must written by 11 number",
			Kind:         richerror.Invalid,
			Meta:         nil,
		}
	}

	if len(req.Password) < 6 {
		return dto.RegisterResponse{}, richerror.RichError{
			Operation:    "userservice.Register",
			WrappedError: nil,
			Message:      "password should be more than 6 character",
			Kind:         richerror.Invalid,
			Meta:         nil,
		}

	}

	// TODO - check regex pattern password

	// hash password
	passHash := hashpassword.EncodePasword(req.Password)

	// save to storage
	newUser, err := s.storage.Register(entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    passHash,
	})
	if err != nil {
		return dto.RegisterResponse{}, richerror.RichError{
			Operation:    "userservice.Register",
			WrappedError: err,
			Message:      "unexpected error",
			Kind:         richerror.Unexpected,
			Meta: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}

	// return new user
	return dto.RegisterResponse{ID: newUser.ID, Name: newUser.Name, PhoneNumber: newUser.PhoneNumber}, nil
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	// TODO - is better use 2 method for check user exist and get user for SOLID (S single responsibility)
	// phone is exist and get user
	user, exist, err := s.storage.CheckUserExistAndGet(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, richerror.RichError{
			Operation:    "userservice.Login",
			WrappedError: err,
			Message:      "unexpected error",
			Kind:         richerror.Unexpected,
			Meta: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}

	// secure reason
	if !exist {
		return LoginResponse{}, richerror.RichError{
			Operation:    "userservice.Login",
			WrappedError: nil,
			Message:      "phone number or password is incorrect",
			Kind:         richerror.Invalid,
			Meta:         nil,
		}
	}

	// check password
	if user.Password != hashpassword.EncodePasword(req.Password) {
		return LoginResponse{}, richerror.RichError{
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
		return LoginResponse{}, richerror.RichError{
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
		return LoginResponse{}, richerror.RichError{
			Operation:    "userservice.Register",
			WrappedError: err,
			Message:      "unexpected error",
			Kind:         richerror.Unexpected,
			Meta: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}

	return LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil

}

type ProfileRequest struct {
	UserID uint `json:"user_id"`
}
type ProfileResponse struct {
	Name string `json:"name"`
}

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	user, err := s.storage.GetUserByID(req.UserID)
	if err != nil {
		return ProfileResponse{}, err
	}

	return ProfileResponse{Name: user.Name}, nil

}
