package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"math/rand"
)

func RandSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

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

func ValidateRequest(data interface{}) []XValidatorErrorResponse {
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

	fmt.Println(validationErrors)

	return validationErrors
}
