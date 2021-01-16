package validations

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// ShouldBeType ..
type ShouldBeType struct {
	Key      string
	Function func(validator.FieldLevel) bool
}

// ValidateShouldBeInteger ...
var ValidateShouldBeInteger = ShouldBeType{
	Key:      "should-be-integer",
	Function: ShouldBeIntegerFunc,
}

// ValidateShouldBeString ...
var ValidateShouldBeString = ShouldBeType{
	Key:      "should-be-string",
	Function: ShouldBeStringFunc,
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
