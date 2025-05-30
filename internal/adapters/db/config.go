package db

type DatabaseConfig struct {
	Host     string `envconfig:"GHS_DB_HOST" default:"localhost"`
	Port     string `envconfig:"GHS_DB_PORT" default:"5432"`
	Username string `envconfig:"GHS_DB_USERNAME" default:"postgres"`
	Password string `envconfig:"GHS_DB_PASSWORD" default:"password"`
	Database string `envconfig:"GHS_DB_DATABASE" default:"go_hexagonal_sandbox"`
}
