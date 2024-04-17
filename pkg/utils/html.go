package utils

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	rAllowedExtensions = regexp.MustCompile(`\.(html|tpl)$`)
	rExtends           = regexp.MustCompile(`{{\s*extends\s+"([^"]+)"\s*}}`)
	rYield             = regexp.MustCompile(`{{\s*yield\s*}}`)
	rBlock             = regexp.MustCompile(`{{\s*block\s+"([^"]+)"\s*}}\s*(.*?)\s*{{\s*end\s*}}`)
)

type HtmlParser struct {
	root      string
	templates map[string]*Template
}

func NewHtmlParser(root string) *HtmlParser {
	h := &HtmlParser{root: root, templates: make(map[string]*Template)}
	if err := h.LoadTemplates(); err != nil {
		log.Fatalf("failed to load templates: %+v\n", err)
	}
	return h
}

func (p *HtmlParser) LoadTemplates() error {
	if err := filepath.Walk(p.root, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && rAllowedExtensions.MatchString(path) {
			name, _ := filepath.Rel(p.root, path)
			p.templates[name] = &Template{name, path, p, make(map[string]string)}
		}
		return err
	}); err != nil {
		return err
	}
	return nil
}
func (p *HtmlParser) Parse(name string) string {
	if _, ok := p.templates[name]; !ok {
		log.Fatalf("template %s not found\n", name)
	}
	return p.templates[name].Render()
}

type Template struct {
	name   string
	path   string
	parser *HtmlParser
	blocks map[string]string
}

func (t *Template) Render() string {
	path := t.path

	for {
		bytes, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("failed to read template: %+v\n", err)
		}
		res := rExtends.FindStringSubmatch(string(bytes))
		if len(res) == 0 {
			return string(bytes)
		}
		bytes = rExtends.ReplaceAllFunc(bytes, func(b []byte) []byte {
			tpl := t.parser.templates[res[1]]
			if tpl == nil {
				log.Printf("template %s not found\n", res[1])
				return nil
			}
			return []byte(tpl.Render())
		})
		path = filepath.Join(t.parser.root, res[1])
	}
	return ""
}
