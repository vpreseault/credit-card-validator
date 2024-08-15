package server

import (
	"fmt"
	"net/http"

	"github.com/vpreseault/credit-card-validator/luhn"
)
func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func HandlerValidate(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	ccn := params["cc-number"][0]
	valid, err := luhn.IsValidLuhn(ccn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	template := ``
	if valid {
		template = `<p class="result-success">%s is a valid credit card number!</p>`
	} else {
		template = `<p class="result-failure">%s is an invalid credit card number.</p>`
	}
	
	response := fmt.Sprintf(template, ccn)	

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(response))
	if err != nil {
		http.Error(w, "Unable to write response", http.StatusInternalServerError)
	}
}