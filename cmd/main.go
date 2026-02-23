package main

import (
	"example/internal/config"
	"example/internal/database"
	"example/internal/handler"
	"example/internal/router"
	"example/internal/session"
	"example/internal/static"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	database.Init()
	database.Migrate()
	session.Init()
	handler.InitTemplates()

	r := chi.NewRouter()
	router.Setup(r)
	static.Setup(r)

	port := ":8080"
	if !config.IsDev() {
		port = ":80"
	}

	log.Printf("listening on %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
