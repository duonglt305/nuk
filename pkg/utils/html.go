package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type HtmlParser struct {
	path      string
	templates map[string]string
}

func NewHtmlParser(path string) HtmlParser {
	return HtmlParser{
		path:      path,
		templates: make(map[string]string),
	}
}

func (p HtmlParser) Parse() {
	fs.WalkDir(os.DirFS(p.path), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			if strings.HasSuffix(path, ".html") {
				dashed := strings.Split(path, ".")[0]
				dotted := strings.Join(
					strings.Split(strings.ReplaceAll(dashed, "\\", "/"), "/"),
					".",
				)
				p.templates[dotted] = filepath.Join(p.path, path)
			}
		}
		return nil
	})
}

func (p HtmlParser) Render(name string, data any) (bytes.Buffer, error) {
	var body bytes.Buffer
	path, ok := p.templates[name]
	if !ok {
		return body, fmt.Errorf("template %s not found", name)
	}
	t, err := template.New(name).ParseFiles(path, p.templates["email.base"])
	if err != nil {
		return body, err
	}

	if err := t.ExecuteTemplate(&body, "email.base", data); err != nil {
		return body, err
	}
	return body, nil
}
