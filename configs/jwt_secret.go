package configs

type JwtConfig interface {
	GetJwtSecretKey() string
}

type jwtConfig struct {
	jwtSecret JwtSecret
}

func NewJwtConfig(jwtSecret JwtSecret) JwtConfig {
	return &jwtConfig{jwtSecret: jwtSecret}
}

func (jwtConf jwtConfig) GetJwtSecretKey() string {
	return jwtConf.jwtSecret.JwtSecretKey
}
