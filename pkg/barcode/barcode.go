package barcode

import (
	"aszaychik/smartcafe-api/config"
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func GenerateBarcode(config *config.BarcodeConfig) error {
	qrCode, err := qr.Encode(config.WifiKey, qr.M, qr.Auto)
	if err != nil {
		return fmt.Errorf("failed encode qr: %w", err)
	}

	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		return fmt.Errorf("failed scale qr: %w", err)
	}

	file, _ := os.Create("qrcode.png")

	png.Encode(file, qrCode)

	return nil
}