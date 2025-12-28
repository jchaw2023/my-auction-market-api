package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
	Message string `json:"message"`
}


func FormatValidationErrors(err error) []ValidationError {
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		return formatValidationErrors(validationErr)
	}
	return formatValidationErrors(err)
}

func formatValidationErrors(err error) []ValidationError {
	var validationErrors []ValidationError

	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErr {
			validationErrors = append(validationErrors, ValidationError{
				Field:   fieldErr.Field(),
				Tag:     fieldErr.Tag(),
				Value:   fmt.Sprintf("%v", fieldErr.Value()),
				Message: getErrorMessage(fieldErr),
			})
		}
	} else {
		validationErrors = append(validationErrors, ValidationError{
			Field:   "",
			Tag:     "",
			Value:   "",
			Message: err.Error(),
		})
	}

	return validationErrors
}

func getErrorMessage(fieldErr validator.FieldError) string {
	field := fieldErr.Field()
	tag := fieldErr.Tag()
	param := fieldErr.Param()

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, param)
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, param)
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", field, param)
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, param)
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, param)
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, param)
	case "lt":
		return fmt.Sprintf("%s must be less than %s", field, param)
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", field, param)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

