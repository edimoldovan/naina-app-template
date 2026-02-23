package handler

import (
	"embed"
	"example/internal/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

//go:embed html
var embeddedHTML embed.FS

var pages = map[string]*template.Template{}

func InitTemplates() {
	layout := "html/layouts/base.html"
	entries, err := embeddedHTML.ReadDir("html/pages")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		name := strings.TrimSuffix(e.Name(), ".html")
		t, err := template.ParseFS(embeddedHTML, layout, "html/pages/"+e.Name())
		if err != nil {
			log.Fatal(err)
		}
		pages[name] = t
	}
}

func render(w http.ResponseWriter, name string, data map[string]any) {
	if data == nil {
		data = map[string]any{}
	}
	data["IsDev"] = config.IsDev()

	if config.IsDev() {
		t, err := template.ParseFiles(
			filepath.Join("internal", "handler", "html", "layouts", "base.html"),
			filepath.Join("internal", "handler", "html", "pages", name+".html"),
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := t.ExecuteTemplate(w, name, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	t, ok := pages[name]
	if !ok {
		http.Error(w, "template not found: "+name, http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
