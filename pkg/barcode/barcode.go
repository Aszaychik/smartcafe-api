package barcode

import (
	"aszaychik/smartcafe-api/config"
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/pkg/uploader"
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)


type BarcodeGenerator interface {
	GenerateBarcodeWifiAccess(order *domain.Order) error
}

type BarcodeGeneratorImpl struct {
	BarcodeConfig *config.BarcodeConfig
	AWSUploader   uploader.AWSUploader
}

func NewBarcodeGenerator(barcodeConfig *config.BarcodeConfig, awsUploader uploader.AWSUploader) BarcodeGenerator {
	return &BarcodeGeneratorImpl{
		BarcodeConfig: barcodeConfig,
		AWSUploader: awsUploader,
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
	qrFolder := fmt.Sprint("wifi-accesses")

	qrFile, _ := os.Create(fmt.Sprintf("uploads/%s", qrFileName))

	png.Encode(qrFile, qrCode)

	_, err = bg.AWSUploader.UploadFile(qrFileName, qrFolder)
	if err != nil {
		return err
	}

	return nil
}