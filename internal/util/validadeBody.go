package util

import (
	"github.com/go-playground/validator"
)

func ValidateBody(body interface{}) error {
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return err
	}

	return nil
}
