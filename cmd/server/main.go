package main

import (
	"log"
	"net/http"

	"github.com/vpreseault/credit-card-validator/server"
)


func main() {
	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/", server.HandlerRoot)
	mux.HandleFunc("GET /api/validate", server.HandlerValidate)


	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}