package validation

import (
	"fmt"
	"profile-creation/types"
	"regexp"

	"github.com/go-playground/validator/v10"
)
var validate = validator.New()
func ValidateUserInput(user types.UserRequest) error {
	panRegex := `^[A-Z]{5}[0-9]{4}[A-Z]{1}$`
	validPan, _ := regexp.MatchString(panRegex, user.Pan); 
	if !validPan {
		return fmt.Errorf("invalid pan number format")
	}

	if err := validate.Struct(user); err != nil {
		return err
	}	
	return nil
}