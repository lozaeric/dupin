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
	V := reflect.ValueOf(o).Elem()
	if err := Check(o, values, updatable); err != nil {
		return err
	}
	for k, v := range values {
		name, _ := toStructName(o, k)
		switch f := V.FieldByName(name); f.Type().String() {
		case "string":
			f.SetString(v.(string))
		case "bool":
			f.SetBool(v.(bool))
		default:
			return errors.New("unknown type")
		}
	}
	return nil
}

func Check(o interface{}, values map[string]interface{}, allowed map[string]bool) error {
	T := reflect.TypeOf(o).Elem()
	for k, v := range values {
		_, ok := allowed[k]
		name, found := toStructName(o, k)
		if !ok || !found {
			return errors.New("invalid field")
		}
		f, _ := T.FieldByName(name)
		actualT, valueT := f.Type.String(), reflect.TypeOf(v).String()
		if actualT != valueT {
			return errors.New("invalid value type")
		}
		if !isValid(f.Tag.Get("validate"), v) {
			return errors.New("invalid value")
		}
	}
	return nil
}

func isValid(validation string, value interface{}) bool {
	return validation == "" || validate.Var(value, validation) == nil
}

func toStructName(o interface{}, jsonTag string) (string, bool) {
	T := reflect.TypeOf(o).Elem()
	for i := 0; i < T.NumField(); i++ {
		if T.Field(i).Tag.Get("json") == jsonTag {
			return T.Field(i).Name, true
		}
	}
	return "", false
}
