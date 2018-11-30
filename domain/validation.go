package domain

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

var validate = validator.New()

func Validate(o interface{}) error {
	err := validate.Struct(o)

	if err != nil {
		var desc string
		for _, err := range err.(validator.ValidationErrors) {
			desc += fmt.Sprintln(err.Field() + ": " + err.ActualTag())
		}
		return errors.New(desc)
	}
	return nil
}
