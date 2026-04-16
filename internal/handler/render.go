package handler

import (
	"embed"
	"io/fs"
	"html/template"
	"log"
	"net/http"
	"nexample/internal/config"
	"path/filepath"
	"strings"
)

//go:embed html
var embeddedHTML embed.FS

var pages = map[string]*template.Template{}

func InitTemplates() {
	partials, err := fs.Glob(embeddedHTML, "html/partials/*.html")
	if err != nil {
		log.Fatal(err)
	}

	entries, err := embeddedHTML.ReadDir("html/pages")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		name := strings.TrimSuffix(e.Name(), ".html")
		files := append([]string{}, partials...)
		files = append(files, "html/pages/"+e.Name())
		t, err := template.ParseFS(embeddedHTML, files...)
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

	pageFile := name + ".html"

	if config.IsDev() {
		partials, err := filepath.Glob(filepath.Join("internal", "handler", "html", "partials", "*.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		files := append([]string{}, partials...)
		files = append(files, filepath.Join("internal", "handler", "html", "pages", pageFile))
		t, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := t.ExecuteTemplate(w, pageFile, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	t, ok := pages[name]
	if !ok {
		http.Error(w, "template not found: "+name, http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(w, pageFile, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
