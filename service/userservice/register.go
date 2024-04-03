package userservice

import (
	"gameapp/entity"
	"gameapp/param"
	"gameapp/pkg/hashpassword"
	"gameapp/pkg/richerror"
)

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
