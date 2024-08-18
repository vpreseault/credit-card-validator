package main

import (
	"log"
	"net/http"

	"github.com/vpreseault/credit-card-validator/server"
)


func main() {
	const filepathRoot = "."
	const port = "8080"

	cfg := server.Config{
		History: []string{},
	}

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	mux.HandleFunc("/", server.HandlerRoot)
	mux.HandleFunc("GET /validate", cfg.HandlerValidate)


	srv := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}