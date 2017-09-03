package main

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/pressly/chi"
)

func main() {
	router := chi.NewRouter()
	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Got a request")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"data": "hello world"}`)
	})
	if err := http.ListenAndServe(":80", router); err != nil {
		log.WithError(err).Fatal("Failed to start HTTP server")
	}
}
