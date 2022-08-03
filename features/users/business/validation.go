package business

import (
	"errors"
	"regexp"
)

func emailFormatValidation(email string) error {
	//	Check syntax email
	pattern := `^\w+@\w+\.\w+$`
	matched, _ := regexp.Match(pattern, []byte(email))
	if !matched {
		return errors.New("failed syntax email")
	}
	return nil
}
