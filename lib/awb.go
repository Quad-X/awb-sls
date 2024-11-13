package lib

import (
	m "github.com/quad-x/awb-sls/models"
	tmp "github.com/quad-x/awb-sls/templates"
	util "github.com/quad-x/awb-sls/utils"
)

// Generate consolidated Airwaybill
func GenerateConsolidatedAWB(data *m.Consolidated) ([]byte, error) {
	var err error

	// Generate QR code
	data.QRCode, err = GenerateQR(data.TrackingNumber, 100, 100)

	if err != nil {
		return nil, err
	}

	// Generate Bar code
	data.BarCode, err = GenerateBarCode(data.TrackingNumber)

	if err != nil {
		return nil, err
	}

	html, err := tmp.ParseTemplate("consolidated", data)

	if err != nil {
		return nil, err
	}

	// Set pdf output path
	filename := data.TrackingNumber + ".pdf"

	outPath := util.Join(util.GetTempDir(), filename)

	// Generate pdf
	output := HtmlToPDF(outPath, "153.2mm", "102.4mm", html)

	return output, nil
}
