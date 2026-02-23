package static

import (
	"embed"
	"net/http"
	"nexample/internal/config"

	"github.com/go-chi/chi/v5"
)

//go:embed assets
var embeddedAssets embed.FS

func Setup(r chi.Router) {
	if config.IsDev() {
		fs := http.FileServer(http.Dir("internal/static"))
		r.Get("/assets/*", http.StripPrefix("/", fs).ServeHTTP)
	} else {
		fs := http.FileServer(http.FS(embeddedAssets))
		r.Get("/assets/*", http.StripPrefix("/", fs).ServeHTTP)
	}
}
