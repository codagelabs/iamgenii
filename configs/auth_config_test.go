package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthConfig(t *testing.T) {

	t.Run("EmptyAuthConfig", func(t *testing.T) {
		authConfig := NewAuthConfig(AuthSecrets{
			PublicKey:  "",
			PrivateKey: "",
		})
		assert.Equal(t, "", authConfig.GetPublicKey())
		assert.Equal(t, "", authConfig.GetPrivateKey())

	})

	t.Run("TestAuthConfig", func(t *testing.T) {
		authConfig := NewAuthConfig(AuthSecrets{
			PublicKey:  "some-piblic-lkeu",
			PrivateKey: "some-private-key",
		})
		assert.Equal(t, "some-piblic-lkeu", authConfig.GetPublicKey())
		assert.Equal(t, "some-private-key", authConfig.GetPrivateKey())

	})

}
