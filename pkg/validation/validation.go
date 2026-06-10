package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func Validate[Type any](toValidate *Type) error {
	validate := validator.New()

	err := validate.Struct(toValidate)

	if err != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			return validateErrs
		}

		return err
	}

	return nil
}
