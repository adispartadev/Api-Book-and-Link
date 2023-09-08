package services

import (
	"github.com/go-playground/validator/v10"
)

type (
	XValidator struct {
		validator *validator.Validate
	}

	XValidatorErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}
)

func (v XValidator) Validate(data interface{}) []XValidatorErrorResponse {
	validationErrors := []XValidatorErrorResponse{}
	var validate = validator.New()

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem XValidatorErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
