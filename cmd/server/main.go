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

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	mux.HandleFunc("/", server.HandlerRoot)
	mux.HandleFunc("/bow", server.HandlerBow)
	mux.HandleFunc("GET /validate", server.HandlerValidate)
	mux.HandleFunc("GET /history", server.HandlerGetHistory)

	srv := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}