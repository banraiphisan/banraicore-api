package utils

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
)

var validate *validator.Validate
var trans ut.Translator

func init() {
	validate = validator.New()
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ = uni.GetTranslator("en")

	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			return field.Name
		}

		return jsonTag
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
