	package configs

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestJwtConfig(t *testing.T) {

	t.Run("EmptyJwtConfig", func(t *testing.T) {
		jwtConfig := NewJwtConfig(JwtSecret{
			JwtSecretKey: "",
		})
		assert.Equal(t, "", jwtConfig.GetJwtSecretKey())

	})

	t.Run("TestJwtConfig", func(t *testing.T) {
		jwtConfig := NewJwtConfig(JwtSecret{
			JwtSecretKey: "some-jwt-secret-key",
		})
		assert.Equal(t, "some-jwt-secret-key", jwtConfig.GetJwtSecretKey())

	})

}