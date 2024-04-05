package uservalidator

import (
	"errors"
	"fmt"
	"gameapp/param"
	"gameapp/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

func (v Validator) Register(req param.RegisterRequest) error {
	const op = "uservalidator.Register"

	// default is invalid ,but it can be Unexpected error from server or not found kind error
	richError.Kind = richerror.Invalid

	// TODO - must add 11 to config
	if err := validation.ValidateStruct(&req,

		// Name cannot be empty, and the length must between 5 and 50
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),

		// Minimum eight characters, or password
		validation.Field(&req.Password, validation.Required, validation.Match(
			regexp.MustCompile(PasswordRegex))),

		// Phone number cannot be empty,  and follow the regular expression like 09359354856
		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile(IRPhoneNumberRegex)),
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
			Kind:             richError.Kind,
			Meta:             richError.Meta,
			ValidationErrors: mapErrs,
		}
	}

	return nil
}

func (v Validator) checkPhoneNumberUniq(value interface{}) error {
	phoneNum := value.(string)
	if isUniq, err := v.storage.IsPhoneNumberUniq(phoneNum); err != nil || !isUniq {
		if err != nil {
			var rErr richerror.RichError
			errors.As(err, &rErr)

			richError.Meta = rErr.Meta
			richError.Kind = rErr.Kind

			return fmt.Errorf("unexpected error")
		}
		if !isUniq {
			return fmt.Errorf("phone number is not uniq")
		}
	}

	return nil

}
