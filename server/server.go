package server

import (
	"fmt"
	"net/http"
)
func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func HandlerValidate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandlerValidate")
}