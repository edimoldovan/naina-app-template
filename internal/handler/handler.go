package handler

import (
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render(w, "home", map[string]any{
		"Title": "nexample",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	render(w, "login", map[string]any{
		"Title": "Login",
	})
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	render(w, "dashboard", map[string]any{
		"Title": "Dashboard",
	})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func PathID(r *http.Request, key string) uint {
	id, err := strconv.ParseUint(r.PathValue(key), 10, 32)
	if err != nil {
		return 0
	}
	return uint(id)
}
