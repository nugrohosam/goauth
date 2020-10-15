package helpers

import (
	validator "github.com/go-playground/validator/v10"
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
