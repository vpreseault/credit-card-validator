package server

import "fmt"

func createTemplate(valid bool, ccn string, history map[string]bool) string {
	resultTemplate := createResultTemplate(valid, ccn)
	historyListTemplate := createHistoryListTemplate(history)

	template := fmt.Sprintf(`
		%s
		%s
	`, resultTemplate, historyListTemplate)

	return template
}

func createResultTemplate(valid bool, ccn string) string {
	resultTemplate := ``
	if valid {
		resultTemplate = `<div id="result"><p class="result-success">%s is a valid credit card number!</p></div>`
	} else {
		resultTemplate = `<div id="result"><p class="result-failure">%s is an invalid credit card number.</p></div>`
	}
	resultTemplate = fmt.Sprintf(resultTemplate, ccn)

	return resultTemplate
}

func createHistoryListTemplate(history map[string]bool) string {
	iconFailure := `<img src="/static/icons/failure2.png" alt="failure icon" class="icons">`
	iconSuccess := `<img src="/static/icons/success1.png" alt="success icon" class="icons">`

	historyItemsTemplate := ``
	for ccn, valid := range history {
		historyItem := ``
		if valid {
			historyItem = fmt.Sprintf(`<li class="history-item">%s %s</li>`, iconSuccess, ccn)
		} else {
			historyItem = fmt.Sprintf(`<li class="history-item">%s %s</li>`, iconFailure, ccn)
		}
		historyItemsTemplate = historyItem + historyItemsTemplate
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