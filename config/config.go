package config

type AppConfig struct {
	MySQL    MySQLConfig
	Midtrans MidtransConfig
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type MidtransConfig struct {
	ServerKey string
}