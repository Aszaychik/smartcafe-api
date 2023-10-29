package config

type AppConfig struct {
	MySQL    MySQLConfig
	Midtrans MidtransConfig
	Barcode  BarcodeConfig
	AWS      AWSConfig
	Auth     AuthConfig
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

type BarcodeConfig struct {
	WifiKey string
}

type AWSConfig struct {
	Region          string
	AccessKeyId     string
	SecretAccessKey string
	BucketName      string
}

type AuthConfig struct {
	XAPIKey   string
	JWTSecret string
}