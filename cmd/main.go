package main

import (
	"log"
	"net/http"
	"nexample/internal/database"
	"nexample/internal/handler"
	"nexample/internal/router"
	"nexample/internal/session"
	"nexample/internal/static"
	"os"

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

	port := ":" + os.Getenv("NEXAMPLE_PORT")

	log.Printf("listening on %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
