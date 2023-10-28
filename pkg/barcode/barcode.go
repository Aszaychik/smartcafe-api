package barcode

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func GenerateBarcode() error {
	qrCode, err := qr.Encode("WIFI:T:WPA;S:AsZaychik;P:sudo apt-get update;H:;;", qr.M, qr.Auto)
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