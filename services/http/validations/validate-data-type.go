package validations

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// ValidateShouldBeInteger ...
var ValidateShouldBeInteger = ShouldBeInteger{
	Key:      "should-be-integer",
	Function: ShouldBeIntegerFunc,
}

// ValidateShouldBeString ...
var ValidateShouldBeString = ShouldBeString{
	Key:      "should-be-integer",
	Function: ShouldBeIntegerFunc,
}

// ShouldBeInteger ..
type ShouldBeInteger struct {
	Key      string
	Function func(validator.FieldLevel) bool
}

// ShouldBeString ..
type ShouldBeString struct {
	Key      string
	Function func(validator.FieldLevel) bool
}

// ShouldBeIntegerFunc ..
func ShouldBeIntegerFunc(field validator.FieldLevel) bool {

	if field.Field().Kind().String() != reflect.Float64.String() {
		return false
	}

	return true
}

// ShouldBeStringFunc ..
func ShouldBeStringFunc(field validator.FieldLevel) bool {

	if field.Field().Kind().String() != reflect.String.String() {
		return false
	}

	return true
}
