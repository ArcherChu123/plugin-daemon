package constants

import (
	"github.com/ArcherChu123/plugin-daemon/pkg/validators"
	"github.com/go-playground/validator/v10"
)

type Arch string

const (
	AMD64 Arch = "amd64"
	ARM64 Arch = "arm64"
)

func isAvailableArch(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return value == string(AMD64) || value == string(ARM64)
}

func init() {
	validators.GlobalEntitiesValidator.RegisterValidation("is_available_arch", isAvailableArch)
}
