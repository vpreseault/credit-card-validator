package server

import (
	"strings"
)

func encodeHistoryCookie(history []HistoryItem) string {
	cookie := ""
	for _, item := range history {
		historyItemString := ""
		if item.Valid {
			historyItemString += "s"
		} else {
			historyItemString += "f"
		}
		historyItemString += item.CCN

		cookie += historyItemString + "|"
	}
	return cookie
}

func decodeHistoryCookie(cookie string) []HistoryItem {
	history := []HistoryItem{}
	historyItemStrings := strings.Split(cookie, "|")
	for _, historyItemString := range historyItemStrings {
		if len(historyItemString) < 2 {
			continue
		}
		valid := false
		if historyItemString[0] == 's' {
			valid = true
		}
		ccn := historyItemString[1:]
		history = append(history, HistoryItem{CCN: ccn, Valid: valid})
	}
	return history
}