package templates

import (
	"bytes"
	"embed"
	"html/template"
	"log"
)

//go:embed *.html
var templates embed.FS

// Parse Template
func ParseTemplate(templateName string, data interface{}) (string, error) {

	buff, err := templates.ReadFile(templateName + ".html")

	if err != nil {
		return "", err
	}

	html := string(buff)

	tmpl, err := template.New(templateName).Parse(html)

	if err != nil {
		log.Fatal("Error parsing template: ", err)
	}

	var htmlBuffer bytes.Buffer
	err = tmpl.Execute(&htmlBuffer, data)

	if err != nil {
		log.Fatal("Error executing template: ", err)
	}

	return htmlBuffer.String(), nil
}
