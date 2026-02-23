package middleware

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, mw ...Middleware) http.HandlerFunc {
	for _, m := range mw {
		f = m(f)
	}
	return f
}
