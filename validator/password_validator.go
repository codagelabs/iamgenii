package validator

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/iamgenii/configs"
)

//PasswordValidator holds password Validator Policies
type PasswordValidator interface {
	IsPassworPolicyCompliant(password string) (bool, error)
}

type passwordValidator struct {
	passPolicyConfig configs.PasswordPolicyConfiguration
}

//NewPasswordValidator hold dependancies
func NewPasswordValidator(
	passPolicyConfig configs.PasswordPolicyConfiguration,
) PasswordValidator {

	return passwordValidator{
		passPolicyConfig: passPolicyConfig,
	}
}

func (passValidate passwordValidator) IsPassworPolicyCompliant(password string) (bool, error) {
	isLower, err := regexp.MatchString(`.*[a-z].*`, password)
	if err != nil {

		return false, errors.New("password must have lowercase latters")
	}

	isUpper, err := regexp.MatchString(`.*[A-Z].*`, password)
	if err != nil {

		return false, errors.New("password must have uppercase latters")
	}

	isNumber, err := regexp.MatchString(`.*[0-9].*`, password)
	if err != nil {

		return false, errors.New("password must have digits ")
	}

	isSpecialChar, err := regexp.MatchString(`.*[!@#$%&*].*`, password)
	if err != nil {

		return false, errors.New("password must have special char `!@#$&*` ")
	}

	passMatchString := fmt.Sprintf("^[A-Za-z0-9!@#$&*]{%d,%d}",
		passValidate.passPolicyConfig.GetPasswordMinLength(),
		passValidate.passPolicyConfig.GetPasswordMaxLength())

	isValid, err := regexp.MatchString(passMatchString, password)

	if err != nil {

		return false, errors.New("password not in proper format")

	}

	isPolicyComplient := isLower && isUpper && isNumber && isSpecialChar && isValid
	if !isPolicyComplient {
		return false, errors.New("password not in proper format")

	}
	return isPolicyComplient, nil
}
