package barcode

import (
	"aszaychik/smartcafe-api/config"
	"aszaychik/smartcafe-api/domain"
	"context"
	"fmt"
	"image/png"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)


type BarcodeGenerator interface {
	GenerateBarcodeWifiAccess(order *domain.Order) error
}

type BarcodeGeneratorImpl struct {
	BarcodeConfig *config.BarcodeConfig
	AWSConfig     *config.AWSConfig
	S3Client      *s3.Client
}

func NewBarcodeGenerator(barcodeConfig *config.BarcodeConfig, awsConfig *config.AWSConfig, s3client *s3.Client) BarcodeGenerator {
	return &BarcodeGeneratorImpl{
		BarcodeConfig: barcodeConfig,
		AWSConfig:     awsConfig,
		S3Client:      s3client,
	}
}

func(bg *BarcodeGeneratorImpl) GenerateBarcodeWifiAccess(order *domain.Order) error {
	qrCode, err := qr.Encode(bg.BarcodeConfig.WifiKey, qr.M, qr.Auto)
	if err != nil {
		return fmt.Errorf("failed encode qr: %w", err)
	}

	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		return fmt.Errorf("failed scale qr: %w", err)
	}

	qrFileName := fmt.Sprintf("%s-wifi-access-%d.png", order.Customer.CustomerName, order.ID)

	qrFile, _ := os.Create(fmt.Sprintf("uploads/%s", qrFileName))

	png.Encode(qrFile, qrCode)

	uploader := manager.NewUploader(bg.S3Client)
	
	file, err := os.Open(fmt.Sprintf("uploads/%s", qrFileName))
	if err != nil {
		return fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bg.AWSConfig.BucketName),
		Key: aws.String(fmt.Sprintf("wifi-accesses/%s", qrFileName)),
		Body: file,
		ACL: "public-read",
	})

	if err != nil {
		return fmt.Errorf("failed upload qr: %w", err)
	}

	fmt.Println(result.Location)

	return nil
}