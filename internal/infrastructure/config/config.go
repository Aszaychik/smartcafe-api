package config

import (
	"aszaychik/smartcafe-api/config"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (*config.AppConfig, error) {
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("failed to load environment variables from .env file: %w", err)
		}
	}

	return &config.AppConfig{
			MySQL: config.MySQLConfig{
					Username: os.Getenv("DB_USERNAME"),
					Password: os.Getenv("DB_PASSWORD"),
					Host:     os.Getenv("DB_HOST"),
					Port:     os.Getenv("DB_PORT"),
					Database: os.Getenv("DB_NAME"),
			},
			Midtrans: config.MidtransConfig{
				ServerKey: os.Getenv("MIDTRANS_SERVER_KEY"),
			},
			Barcode: config.BarcodeConfig{
				WifiKey: os.Getenv("WIFI_SECRET_KEY"),
			},
			AWS: config.AWSConfig{
				Region: os.Getenv("AWS_REGION"),
				AccessKeyId: os.Getenv("AWS_ACCESS_KEY_ID"),
				SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
				BucketName: os.Getenv("AWS_BUCKET_NAME"),
			},
	}, nil
}