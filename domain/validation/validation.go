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

func IsValidID(ID string) bool {
	return validate.Var(ID, "required,len=20,alphanum") == nil
}

func Update(o interface{}, values map[string]interface{}, updatable map[string]bool) error {
	V, T := reflect.ValueOf(o).Elem(), reflect.TypeOf(o).Elem()
	for k, v := range values {
		if _, ok := updatable[k]; !ok {
			return errors.New("invalid field")
		}
		for i := 0; i < T.NumField(); i++ {
			f := T.Field(i)
			if f.Tag.Get("json") == k {
				actualT, valueT := f.Type.String(), reflect.TypeOf(v).String()
				if actualT != valueT {
					return errors.New("invalid value type")
				}
				if !isValid(f.Tag.Get("validate"), v) {
					return errors.New("invalid value")
				}
				switch actualT {
				case "string":
					V.Field(i).SetString(v.(string))
				case "bool":
					V.Field(i).SetBool(v.(bool))
				}
			}
		}
	}
	return nil
}

func isValid(validation string, value interface{}) bool {
	return validation == "" || validate.Var(value, validation) == nil
}
