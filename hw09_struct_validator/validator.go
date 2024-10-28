package hw09structvalidator

import (
	"errors"
	"reflect"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	panic("implement me")
}

func Validate(v interface{}) error {
	t := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	if t.Kind() != reflect.Struct {
		return errors.New("value must be a struct")
	}
	ve := make(ValidationErrors, 0)
	var ok ValidationError
	var err error
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("validate") != "" {
			switch t.Field(i).Type {
			case reflect.TypeOf(""):
				ok, err = validateString(t.Field(i), vv.Field(i))
			case reflect.TypeOf(0):
				ok, err = validateInt(t.Field(i), vv.Field(i))
			default:
				err = errors.New(t.Field(i).Name + ": " + t.Field(i).Type.String() + "(this type is not supported)")
			}
			if err != nil {
				return err
			}
			if !errors.Is(ok.Err, nil) {
				ve = append(ve, ok)
			}

		}
	}

	return nil
}

func validateInt(field reflect.StructField, field2 reflect.Value) (ValidationError, error) {
	return ValidationError{field.Name, nil}, nil
}

func validateString(field reflect.StructField, field2 reflect.Value) (ValidationError, error) {
	return ValidationError{field.Name, nil}, nil
}
