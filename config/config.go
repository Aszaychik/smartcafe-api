package config

type AppConfig struct {
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}