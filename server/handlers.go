package server

import (
	"fmt"
	"net/http"

	"github.com/vpreseault/credit-card-validator/luhn"
)

type Config struct {
	History []string
}

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func (cfg *Config) HandlerValidate(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	ccn := params["cc-number"][0]
	valid, err := luhn.IsValidLuhn(ccn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	template := cfg.createTemplate(valid, ccn)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(template))
	if err != nil {
		http.Error(w, "Unable to write response", http.StatusInternalServerError)
	}
}

func (cfg *Config) createTemplate(valid bool, ccn string) string {
	iconFailure := `<img src="/static/icons/failure2.png" alt="failure icon" class="icons">`
	iconSuccess := `<img src="/static/icons/success1.png" alt="success icon" class="icons">`

	resultTemplate := ``
	historyItem := ``
	if valid {
		resultTemplate = `<div id="result"><p class="result-success">%s is a valid credit card number!</p></div>`
		historyItem = fmt.Sprintf(`<li class="history-item">%s %s</li>`, iconSuccess, ccn)
	} else {
		resultTemplate = `<div id="result"><p class="result-failure">%s is an invalid credit card number.</p></div>`
		historyItem = fmt.Sprintf(`<li class="history-item">%s %s</li>`, iconFailure, ccn)
	}
	resultTemplate = fmt.Sprintf(resultTemplate, ccn)

	cfg.History = append(cfg.History, historyItem)

	historyListTemplate := cfg.createHistoryListTemplate()

	template := fmt.Sprintf(`
		%s
		%s
	`, resultTemplate, historyListTemplate)

	return template
}

func (cfg *Config) createHistoryListTemplate() string {
	historyItemsTemplate := ``
	for i := len(cfg.History) - 1; i >= 0; i-- {
		historyItemsTemplate += cfg.History[i]
	}

	template := fmt.Sprintf(`
		<div id="history-container">
			<h2>History</h2>
			<ul id="history">
				%s
			</ul>
    	</div>
	`, historyItemsTemplate)

	return template
}