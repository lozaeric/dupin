package validation

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator"
)

var validate = validator.New()

func Validate(o interface{}) error {
	if err := validate.Struct(o); err != nil {
		var desc string
		for _, err := range err.(validator.ValidationErrors) {
			desc += fmt.Sprintln(err.Field() + ": " + err.ActualTag())
		}
		return errors.New(desc)
	}
	return nil
}

func IsValid(object reflect.Type, fieldName string, value interface{}) bool {
	field, found := reflect.TypeOf(object).FieldByName(fieldName)
	if !found {
		return false
	}
	tag, found := field.Tag.Lookup("validate")
	return found && validate.Var(value, tag) == nil
}

func IsValidID(ID string) bool {
	return validate.Var(ID, "required,len=20,alphanum") == nil
}
