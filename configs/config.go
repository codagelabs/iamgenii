package configs

//Config store Project Configuration
type config struct {
	JwtSecret             JwtSecret             `json:"jwt_secret"`
	DBConfig              DBConfig              `json:"db_config"`
	PasswordConfiguration PasswordConfiguration `json:"password_config"`
	AuthSecrets           AuthSecrets           `json:"auth_config"`
}

//JwtSecret store jwt encryption key
type JwtSecret struct {
	JwtSecretKey string `json:"jwt_secret_key"`
}

//PasswordConfiguration store Password Config Details
type PasswordConfiguration struct {
	PasswordMaxLength int64 `json:"password_max_length"`
	PasswordMinLength int64 `json:"password_min_length"`
}
type DbConnectionPool struct {
	MaxOpenConnection              int `json:"max_open_connection"`
	MaxIdealConnection             int `json:"max_ideal_connection"`
	MaxConnectionLifetimeInMinutes int `json:"max_connection_lifetime_in_minutes"`
}

type DBConfig struct {
	DBLogMode        bool             `json:"db_log_mode"`
	DbConnectionPool DbConnectionPool `json:"db_connection_pool"`
}

type SendInBlueEmailConfign struct {
}

//AuthSecrets store jwt encryption key
type AuthSecrets struct {
	PublicKey  string `json:"jwt_encryption_public_key"`
	PrivateKey string `json:"jwt_encryption_private_key"`
}
