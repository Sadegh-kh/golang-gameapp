package uservalidator

import (
	"errors"
	"fmt"
	"gameapp/dto"
	"gameapp/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

type Storage interface {
	IsPhoneNumberUniq(phoneNumber string) (bool, error)
}
type Validator struct {
	storage Storage
}

func New(stg Storage) Validator {
	return Validator{storage: stg}
}

func (v Validator) Register(req dto.RegisterRequest) error {
	const op = "uservalidator.Register"

	// TODO - must add 11 to config
	if err := validation.ValidateStruct(&req,

		// Name cannot be empty, and the length must between 5 and 50
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),

		// Minimum eight characters, at least one uppercase letter, one lowercase letter and one number
		validation.Field(&req.Password, validation.Required, validation.Match(
			regexp.MustCompile(`^[A-Za-z0-9@#$%!*&]{8,}$`))),

		// Phone number cannot be empty,  and follow the regular expression like 09359354856
		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile(`^09[0-9]{9}$`)),
			validation.By(v.checkPhoneNumberUniq)),
	); err != nil {
		var vErrs validation.Errors
		errors.As(err, &vErrs)

		mapErrs := make(map[string]string)
		for key, value := range vErrs {
			mapErrs[key] = value.Error()
		}

		return richerror.RichError{
			Operation:        op,
			WrappedError:     err,
			Message:          err.Error(),
			Kind:             richerror.Invalid,
			Meta:             nil,
			ValidationErrors: mapErrs,
		}
	}

	return nil
}

func (v Validator) checkPhoneNumberUniq(value interface{}) error {
	phoneNum := value.(string)
	if isUniq, err := v.storage.IsPhoneNumberUniq(phoneNum); err != nil || !isUniq {
		if err != nil {
			return err
		}
		if !isUniq {
			return fmt.Errorf("phone number is not uniq")
		}
	}

	return nil

}
