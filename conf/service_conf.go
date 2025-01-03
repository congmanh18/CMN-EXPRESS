package conf

type ServiceConfig struct {
	ServiceName string `mapstructure:"SERVICE_NAME"`
	ServicePort int    `mapstructure:"SERVICE_PORT"`

	DBHost string `mapstructure:"POSTGRES_HOST"`
	DBPort int    `mapstructure:"POSTGRES_PORT"`
	DBUser string `mapstructure:"POSTGRES_USER"`
	DBPwd  string `mapstructure:"POSTGRES_PASSWORD"`
	DBName string `mapstructure:"POSTGRES_DB"`

	RedisAddr string `mapstructure:"REDIS_ADDR"`
	RedisPass string `mapstructure:"REDIS_PASSWORD"`
	RedisPort int    `mapstructure:"REDIS_PORT"`

	SecretKey    string `mapstructure:"SECRET_KEY"`
	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

// docker run --name db -e POSTGRES_DB=local -e POSTGRES_USER=local -e POSTGRES_PASSWORD=123456 -p 5432:5432 -d postgres:17-alpine
