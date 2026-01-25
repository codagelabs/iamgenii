package configs

//PasswordPolicyConfiguration stores all password related configuration
type PasswordPolicyConfiguration interface {
	GetPasswordMaxLength() int64
	GetPasswordMinLength() int64
}

//NewPasswordPolicyConfiguration **
func NewPasswordPolicyConfiguration(passwordConfig PasswordConfiguration) PasswordPolicyConfiguration {
	return &passwordPolicyConfigurationImpl{passwordConfiguration: passwordConfig}
}

// passwordPolicyConfigurationImpl represent password configuration
type passwordPolicyConfigurationImpl struct {
	passwordConfiguration PasswordConfiguration
}

//GetPasswordMaxLength return password maximum length configuration
func (obj passwordPolicyConfigurationImpl) GetPasswordMaxLength() int64 {
	return obj.passwordConfiguration.PasswordMaxLength
}

//GetPasswordMinLength return passwords minimum length configuration
func (obj passwordPolicyConfigurationImpl) GetPasswordMinLength() int64 {
	return obj.passwordConfiguration.PasswordMinLength
}
