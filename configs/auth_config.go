package configs

type AuthConfig interface {
	GetPublicKey() string
	GetPrivateKey() string
}

type authConfig struct {
	authSecrets AuthSecrets
}

func (config authConfig) GetPublicKey() string {
	return config.authSecrets.PublicKey
}

func (config authConfig) GetPrivateKey() string {
	return config.authSecrets.PrivateKey
}

func NewAuthConfig(authSecrets AuthSecrets) AuthConfig {
	return &authConfig{authSecrets: authSecrets}
}
