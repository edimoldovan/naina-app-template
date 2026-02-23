package middleware

import (
	"example/internal/session"
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acc := session.Get(r)
		if acc.ID == 0 {
			if r.Header.Get("X-Requested-With") == "xmlhttprequest" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			} else {
				http.Redirect(w, r, "/login", http.StatusFound)
			}
			return
		}
		next(w, r)
	}
}
