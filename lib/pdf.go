package lib

import (
	"log"
	"os"
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// Convert html into PDF
func HtmlToPDF(outputPath string, width string, height string, htmls ...string) []byte {

	if os.Getenv("WKHTMLTOPDF_PATH") == "" {
		os.Setenv("WKHTMLTOPDF_PATH", os.Getenv("LAMBDA_TASK_ROOT"))
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Grayscale.Set(true)
	pdfg.MarginTop.Set(0)
	pdfg.MarginBottom.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.PageHeightUnit.Set(height)
	pdfg.PageWidthUnit.Set(width)

	// Create a new input page from an URL
	for _, html := range htmls {
		// create html page
		page := wkhtmltopdf.NewPageReader(strings.NewReader(html))
		// Zoom Page
		page.Zoom.Set(1.5)
		// Add to document
		pdfg.AddPage(page)
	}

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// for debugging purposes
	// err = pdfg.WriteFile(outputPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return pdfg.Bytes()
}
