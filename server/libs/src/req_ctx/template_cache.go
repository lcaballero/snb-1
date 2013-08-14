package req_ctx

import (
	"html/template"
	"path"
	"fmt"
	"io/ioutil"
)

type HtmlTemplateCache struct {
	TemplateDirs map[string]bool
	mapping map[string]*template.Template
}

var template_cache *HtmlTemplateCache = nil

func TemplateCache() *HtmlTemplateCache {
	if template_cache == nil {
		template_cache = &HtmlTemplateCache{
			TemplateDirs: make(map[string]bool),
			mapping: make(map[string]*template.Template),
		}
	}
	return template_cache
}

func (h *HtmlTemplateCache) addRoot(root string) {
	_, ok := h.TemplateDirs[root]

	if !ok {
		h.TemplateDirs[root] = true
	}
}

func (h *HtmlTemplateCache) Add(root, file string, t *template.Template) {
	cache := *TemplateCache()
	cache.addRoot(root)

	key := path.Join(root, file)
	cache.mapping[key] = t
}

func (h *HtmlTemplateCache) AddAll(root string, fn func(r, f string) (*template.Template), paths ...string) {
	cache := *h
	cache.addRoot(root)

	for _, p := range paths {
		key := path.Join(root, p)
		cache.mapping[key] = fn(root, p)
	}
}

func (h *HtmlTemplateCache) FindTemplate(file string) (*template.Template, bool) {
	cache := *TemplateCache()

	for k, v := range cache.mapping {
		fmt.Println(k, v == nil)
	}

	for root, _ := range cache.TemplateDirs {
		file := path.Join(root, file)
		fmt.Println("Looking up: ", file)
		v, ok := cache.mapping[file]

		if ok {
			return v, ok;
		}
	}

	return nil, false;
}

func CompileHtmlTemplate(r, f string) *template.Template {

	p := path.Join(r, f)
	bytes, err1 := ioutil.ReadFile(p)

	if err1 != nil {
		fmt.Println("Couldn't ReadFile: ", p)
		return nil
	}

	tmpl, err2 := template.New(p).Parse(string(bytes))

	if err2 != nil {
		fmt.Println("Couldn't parse template: ", p)
		return nil
	}

	return tmpl
}