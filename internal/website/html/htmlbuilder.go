package page

import (
	"bytes"
	"html/template"
	"log"
)

func Build(templates []string, data any) ([]byte, error) {
	templ := template.Must(template.ParseFiles(templates...))
	var buffer bytes.Buffer
	err := templ.Execute(&buffer, data)
	if err != nil {
		log.Printf("failed to build html page template: %v", err.Error())
		return nil, err
	}
	return buffer.Bytes(), nil
}
