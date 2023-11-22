package utils

import (
	"github.com/go-playground/validator/v10"
	"log/slog"
	"regexp"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	if err := validate.RegisterValidation("password", validatePassword); err != nil {
		slog.Error("register password validation failed", err.Error())
	}
}

func validatePassword(f validator.FieldLevel) bool {
	value := f.Field().String()
	hasLetters, _ := regexp.MatchString(`[A-Za-z]`, value)
	hasDigit, _ := regexp.MatchString(`[0-9]`, value)
	hasSpecialChar, _ := regexp.MatchString(`[!@#$%^&*(),.?":{}|<>]`, value)
	hasSpace, _ := regexp.MatchString(`\s`, value)
	return hasLetters && hasDigit && !hasSpace || hasSpecialChar
}

func Verify(data any) error {
	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}
