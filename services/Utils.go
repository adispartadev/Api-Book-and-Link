package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"reflect"
	"strings"
)

func RandSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func ValidatingRequest(c *fiber.Ctx, formData interface{}) (status bool, data any, message string) {
	if err := c.BodyParser(formData); err != nil {
		return false, err.Error(), "Error occured"
	}

	validate := validator.New()
	fieldErrors := make(map[string]string)

	if err := validate.Struct(formData); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			fieldName := getFieldJSONName(formData, field)
			fieldErrors[fieldName] = generateValidationMessage(err.Tag())
		}
		return false, fieldErrors, "Validation error"
	}

	return true, nil, ""
}

func getFieldJSONName(s interface{}, fieldName string) string {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	field, _ := t.FieldByName(fieldName)
	jsonTag := field.Tag.Get("json")
	jsonName := strings.Split(jsonTag, ",")[0]

	return jsonName
}

func generateValidationMessage(tag string) string {
	var message string
	switch tag {
	case "required":
		message = "The field is required."
	case "email":
		message = "Invalid email format."
	case "eqfield":
		message = "Value was not same."
	}
	return message
}
