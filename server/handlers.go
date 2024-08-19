package server

import (
	"log"
	"net/http"

	"github.com/vpreseault/credit-card-validator/luhn"
)

type Config struct {
	History []string
}

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

	history := make(map[string]bool)
	cookie, err := r.Cookie("validationHistory")
	if err == nil {
		history = decodeHistoryCookie(cookie.Value)
	} else {
		log.Printf("Error getting cookie: %v", err)
	}
	
	history[ccn] = valid

	historyString := encodeHistoryCookie(history)
	http.SetCookie(w, &http.Cookie{
		Name:  "validationHistory",
		Value: historyString,
		Path:  "/",
	})

	template := createTemplate(valid, ccn, history)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(template))
	if err != nil {
		http.Error(w, "Unable to write response", http.StatusInternalServerError)
	}
}