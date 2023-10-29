package customvalidator

import (
	"strings"

	"github.com/go-playground/validator"
)

func MapValidationErrors(err error) map[string]string {
	ret := map[string]string{}

	for _, err := range err.(validator.ValidationErrors) {
		key := strings.ToLower(err.Field())
		ret[key] = "Type error: " + err.Tag()
	}

	return ret
}
