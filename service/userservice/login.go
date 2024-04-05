package userservice

import (
	"gameapp/param"
	"gameapp/pkg/hashpassword"
	"gameapp/pkg/richerror"
)

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
