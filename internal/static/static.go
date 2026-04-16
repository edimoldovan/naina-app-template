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
		r.Get("/assets/*", http.StripPrefix("/", noCache(fs)).ServeHTTP)
	} else {
		fs := http.FileServer(http.FS(embeddedAssets))
		r.Get("/assets/*", http.StripPrefix("/", fs).ServeHTTP)
	}
}

func noCache(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		h.ServeHTTP(w, r)
	})
}
