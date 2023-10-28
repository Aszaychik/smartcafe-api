package config

type AppConfig struct {
	MySQL    MySQLConfig
	Midtrans MidtransConfig
	Barcode  BarcodeConfig
	AWS      AWSConfig
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