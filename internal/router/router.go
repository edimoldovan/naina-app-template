package router

import (
	"net/http"
	"nexample/internal/handler"
	"nexample/internal/middleware"
	"strings"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

var (
	public  = []middleware.Middleware{}
	private = []middleware.Middleware{middleware.Auth}
)

func Setup(r chi.Router) chi.Router {
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.URL.Path, "/assets/") {
				chimw.Logger(next).ServeHTTP(w, r)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	})

	r.Get("/up", handler.HealthCheck)
	r.Get("/health", handler.HealthCheck)

	// public
	r.Get("/", middleware.Chain(handler.Home, public...))
	r.Get("/login", middleware.Chain(handler.Login, public...))

	// private
	r.Get("/dashboard", middleware.Chain(handler.Dashboard, private...))

	return r
}
