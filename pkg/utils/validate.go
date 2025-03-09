package utils

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/google/uuid"
	"reflect"
)

var validate *validator.Validate
var trans ut.Translator

func PasswordRequired(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Get the "ID" field from the struct
	idField := fl.Parent().FieldByName("ID")

	// If the ID field exists and is NOT the zero value, it's an update → allow empty password
	if idField.IsValid() && idField.CanInterface() {
		idValue, ok := idField.Interface().(uuid.UUID) // Convert to uuid.UUID
		if ok && idValue != uuid.Nil {
			return true // Allow empty password for updates
		}
	}

	// If ID is missing (new user), password is required
	return password != ""
}

func init() {
	validate = validator.New()
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ = uni.GetTranslator("en")
	_ = validate.RegisterValidation("password_required", PasswordRequired)

	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			return field.Name
		}

		return jsonTag
	})

	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field()

		if field.Kind() == reflect.String { // Case when it's a string
			_, err := uuid.Parse(field.String())
			return err == nil
		}

		if field.Kind() == reflect.Struct { // Case when it's uuid.UUID
			_, ok := field.Interface().(uuid.UUID)
			return ok
		}

		return false
	})
}

func ValidateStruct(input interface{}) error {
	if err := validate.Struct(input); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("validation failed for %s", err.Translate(trans))
		}
	}

	return nil
}
