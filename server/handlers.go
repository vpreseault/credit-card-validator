package server

import (
	"net/http"

	"github.com/vpreseault/credit-card-validator/luhn"
)

type HistoryItem struct {
	CCN   string
	Valid bool
}

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func HandlerBow(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/bow.html")
}

func HandlerValidate(w http.ResponseWriter, r *http.Request) {
	ccn := r.URL.Query().Get("cc-number")
	valid, err := luhn.IsValidLuhn(ccn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	history := []HistoryItem{}
	cookie, err := r.Cookie("validationHistory")
	if err == nil {
		history = decodeHistoryCookie(cookie.Value)
	}
	
	history = addHistoryItem(history, ccn, valid)

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

func HandlerGetHistory(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("validationHistory")
	if err != nil {
		// No history cookie
		return
	}

	history := decodeHistoryCookie(cookie.Value)
	template := createHistoryListTemplate(history)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(template))
	if err != nil {
		http.Error(w, "Unable to write response", http.StatusInternalServerError)
	}
}