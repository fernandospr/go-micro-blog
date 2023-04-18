package config

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

type Server struct {
	Port string `yaml:"port" env:"SRV_PORT" env-default:"8080" env-description:"Server port"`
}

type Database struct {
	Type     string `yaml:"type" env:"DB_TYPE" env-default:"in-memory" env-description:"Database type (in-memory, sqlite or mysql)"`
	Host     string `yaml:"host" env:"DB_HOST" env-description:"Database host"`
	Port     string `yaml:"port" env:"DB_PORT" env-description:"Database port"`
	Username string `yaml:"username" env:"DB_USER" env-description:"Database user name"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-description:"Database user password"`
	Name     string `yaml:"db-name" env:"DB_NAME" env-description:"Database name"`
}
