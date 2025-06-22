package constants

import (
	"github.com/ArcherChu123/plugin-daemon/pkg/validators"
	"github.com/go-playground/validator/v10"
)

type Language string

const (
	Python Language = "python"
	Go     Language = "go" // not supported yet
)

func isAvailableLanguage(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	switch value {
	case string(Python):
		return true
	}
	return false
}

func init() {
	validators.GlobalEntitiesValidator.RegisterValidation("is_available_language", isAvailableLanguage)
}
