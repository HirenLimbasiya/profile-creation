package validation

import (
	"fmt"
	"profile-creation/types"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	validate.RegisterValidation("pan", validatePAN)
}

func ValidateUserInput(user types.UserRequest) error {
	if err := validate.Struct(user); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return formatValidationError(err)
		}
	}

	if len(strconv.FormatInt(user.Mobile, 10)) != 10 {
		return fmt.Errorf("field 'Mobile' failed validation: must be a 10-digit number")
	}

	return nil
}

func validatePAN(fl validator.FieldLevel) bool {
	panRegex := `^[A-Z]{5}[0-9]{4}[A-Z]{1}$`
	matched, _ := regexp.MatchString(panRegex, fl.Field().String())
	return matched
}

func formatValidationError(err validator.FieldError) error {
	return fmt.Errorf("field '%s' failed validation: %s", err.Field(), err.Tag())
}
