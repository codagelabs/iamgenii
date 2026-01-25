package validator

import (
	"fmt"
	"regexp"
	"strings"
)

//MobileNumberValidator validates mobile number
type MobileNumberValidator interface {
	IsValidMobileNumber(string) bool
}

type mobileNumberValidator struct {
}

//NewMobileNumberValidator returns MobileNumberValidator
func NewMobileNumberValidator() MobileNumberValidator {
	return &mobileNumberValidator{}
}

func (validator mobileNumberValidator) IsValidMobileNumber(mobileNumber string) bool {

	mustNubers := regexp.MustCompile(`^[0-9]`).MatchString(mobileNumber)
	var isValid bool
	if strings.HasPrefix(mobileNumber, "91") {

		isValid = regexp.MustCompile(`^\+?91\d{10}$`).MatchString(mobileNumber)

	} else {

		isValid = regexp.MustCompile(`^[0-9]\d{8,14}$`).MatchString(mobileNumber)

	}
	fmt.Println(mustNubers && isValid, mustNubers, isValid)
	return mustNubers && isValid
}
