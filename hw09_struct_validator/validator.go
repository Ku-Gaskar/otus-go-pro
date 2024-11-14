package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var errs []string
	for _, e := range v {
		errs = append(errs, fmt.Sprintf("Validation error: field %s -> %s", e.Field, e.Err))
	}
	return strings.Join(errs, "\n")
}

func Validate(v interface{}) error {
	t := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	if t.Kind() != reflect.Struct {
		return errors.New("value must be a struct")
	}

	var ve ValidationErrors

	for i := 0; i < t.NumField(); i++ {
		if vv.Field(i).CanInterface() && t.Field(i).Tag.Get("validate") != "" {
			for _, rule := range strings.Split(t.Field(i).Tag.Get("validate"), "|") {
				ve = applyOneRule(rule, t.Field(i).Name, vv.Field(i), ve)
			}
		}

	}
	return ve
}

func applyOneRule(rule string, nameField string, dataForValidate reflect.Value, ves ValidationErrors) ValidationErrors {
	rl := strings.SplitN(rule, ":", 2)
	if len(rl) != 2 {
		return append(ves, ValidationError{nameField, errors.New("invalid rule")})
	}

	switch dataForValidate.Kind() {
	case reflect.String:
		err := validateString(rl[0], rl[1], nameField, dataForValidate.String())
		if !errors.Is(err.Err, nil) {
			ves = append(ves, err)
		}

	case reflect.Int:
		err := validateInt(rl[0], rl[1], nameField, dataForValidate.Int())
		if !errors.Is(err.Err, nil) {
			ves = append(ves, err)
		}
		return ves

	case reflect.Slice:
		for y := 0; y < dataForValidate.Len(); y++ {
			ves = applyOneRule(rule, nameField, dataForValidate.Index(y), ves)
		}
		return ves

	default:
		return append(ves, ValidationError{nameField,
			errors.New(nameField + ": " + "(Field type is not supported)")})
	}
	return ves
}

func validateString(ruleName string, ruleValue string, nameField string, dataStringForValidator string) ValidationError {
	switch ruleName {
	case "len":
		rv, err := strconv.Atoi(ruleValue)
		if err != nil {
			return ValidationError{nameField, err}
		}
		if len(dataStringForValidator) != rv {
			return ValidationError{nameField, errors.New("length must be " + ruleValue)}
		}

	case "regexp":
		re, err := regexp.Compile(ruleValue)
		if err != nil {
			return ValidationError{nameField, errors.New("invalid regexp pattern")}
		}
		if !re.MatchString(dataStringForValidator) {
			return ValidationError{nameField, errors.New("regexp does not match " + dataStringForValidator)}
			//return fmt.Errorf("must match regexp %s", ruleValue)
		}
	case "in":
		values := strings.Split(ruleValue, ",")
		for _, v := range values {
			if dataStringForValidator == v {
				return ValidationError{nameField, nil}
			}
		}
		return ValidationError{nameField, errors.New("must be one of " + ruleValue)}

	default:
		return ValidationError{nameField, errors.New("unknown rule " + ruleName)}

	}
	return ValidationError{nameField, nil}
}

func validateInt(ruleName string, ruleValue string, nameField string, dataIntForValidator int64) ValidationError {
	switch ruleName {
	case "min":
		rv, err := strconv.ParseInt(ruleValue, 10, 64)
		if err != nil {
			return ValidationError{nameField, err}
		}
		if dataIntForValidator < rv {
			return ValidationError{nameField, errors.New("must be >= " + ruleValue)}
		}
	case "max":
		rv, err := strconv.ParseInt(ruleValue, 10, 64)
		if err != nil {
			return ValidationError{nameField, err}
		}
		if dataIntForValidator > rv {
			return ValidationError{nameField, errors.New("max is " + ruleValue)}
		}
	case "in":
		values := strings.Split(ruleValue, ",")
		for _, v := range values {
			vInt, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return ValidationError{nameField, err}
			}
			if dataIntForValidator == vInt {
				return ValidationError{nameField, nil}
			}
		}
		return ValidationError{nameField, errors.New("must be one of " + ruleValue)}

	default:
		return ValidationError{nameField, errors.New("unknown rule " + ruleName)}
	}

	return ValidationError{nameField, nil}
}
