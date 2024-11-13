package lib

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code39"
	"github.com/boombuler/barcode/qr"
)

// Convert image to base64 image string
func toBase64Image(m image.Image) (*string, error) {
	var imageBuffer bytes.Buffer
	err := png.Encode(&imageBuffer, m)

	if err != nil {
		return nil, err
	}

	base64Image := base64.StdEncoding.EncodeToString(imageBuffer.Bytes())
	return &base64Image, err
}

// Generate qrcode
func GenerateQR(data string, width int, height int) (*string, error) {
	// Create the barcode
	qrCode, err := qr.Encode(data, qr.M, qr.Auto)

	if err != nil {
		return nil, err
	}

	// Scale the qr code to defined size
	qrCode, _ = barcode.Scale(qrCode, width, height)

	return toBase64Image(qrCode)
}

// Generate barcode
func GenerateBarCode(data string) (*string, error) {
	// Create the barcode
	barcode_, err := code39.Encode(data, true, true)

	if err != nil {
		return nil, err
	}

	return toBase64Image(barcode_)
}
