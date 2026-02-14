package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordPolicyConfiguration(t *testing.T) {
	t.Run("ValidConfiguration", func(t *testing.T) {
		config := PasswordConfiguration{
			PasswordMaxLength: 20,
			PasswordMinLength: 8,
		}

		policyConfig := NewPasswordPolicyConfiguration(config)

		assert.Equal(t, int64(20), policyConfig.GetPasswordMaxLength())
		assert.Equal(t, int64(8), policyConfig.GetPasswordMinLength())
	})

	t.Run("ZeroValues", func(t *testing.T) {
		config := PasswordConfiguration{
			PasswordMaxLength: 0,
			PasswordMinLength: 0,
		}

		policyConfig := NewPasswordPolicyConfiguration(config)

		assert.Equal(t, int64(0), policyConfig.GetPasswordMaxLength())
		assert.Equal(t, int64(0), policyConfig.GetPasswordMinLength())
	})
}
