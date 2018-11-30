package utils

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

func IsValidID(ID string) bool {
	return validate.Var(ID, "required,len=20,alphanum") == nil
}
