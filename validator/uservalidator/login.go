package uservalidator

import (
	"errors"
	"fmt"
	"gameapp/param"
	"gameapp/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

func (v Validator) Login(req param.LoginRequest) error {
	const op = "uservalidator.Register"

	richError.Kind = richerror.Invalid

	// TODO - must add 11 to config
	if err := validation.ValidateStruct(&req,

		// Minimum eight characters, or number
		validation.Field(&req.Password, validation.Required, validation.Match(
			regexp.MustCompile(PasswordRegex))),

		// Phone number cannot be empty,  and follow the regular expression like 09359354856
		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile(IRPhoneNumberRegex)),
			validation.By(v.checkUserExist)),
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
			Kind:             richError.Kind,
			Meta:             richError.Meta,
			ValidationErrors: mapErrs,
		}
	}

	return nil
}

func (v Validator) checkUserExist(value interface{}) error {
	pNum := value.(string)
	_, err := v.storage.GetUserByPhoneNumber(pNum)
	if err != nil {
		var rErr richerror.RichError
		errors.As(err, &rErr)

		richError.Kind = rErr.Kind
		richError.Meta = rErr.Meta

		if rErr.Kind == richerror.NotFound {
			return fmt.Errorf("phone number or password is incorrect")
		}

		return err
	}
	return nil
}
