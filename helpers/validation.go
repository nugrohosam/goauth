package helpers

import (
	"reflect"
	"strings"

	validator "github.com/go-playground/validator/v10"
	validations "github.com/nugrohosam/gosampleapi/services/http/validations"
)

// TransformValidations ...
func TransformValidations(validationErrors validator.ValidationErrors) []map[string]interface{} {
	fieldsErrors := make([]map[string]interface{}, len(validationErrors))
	for index, fieldErr := range validationErrors {
		fieldsErrors[index] = map[string]interface{}{
			"key":  fieldErr.Field(),
			"kind": fieldErr.ActualTag(),
		}
	}

	return fieldsErrors
}

// NewValidation ..
func NewValidation() *validator.Validate {
	validate := validator.New()

	// Load all custom validation
	validate.RegisterValidation(validations.ValidateShouldBeInteger.Key, validations.ValidateShouldBeInteger.Function)
	validate.RegisterTagNameFunc(
		func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})

	return validate
}
