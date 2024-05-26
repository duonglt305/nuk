package utils

import (
	"bytes"
	"html/template"
)

func ParseHTML(name string, data map[string]any, files ...string) (*bytes.Buffer, error) {
	t, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}
	var body bytes.Buffer
	if err := t.ExecuteTemplate(&body, name, data); err != nil {
		return nil, err
	}
	return &body, nil
}
