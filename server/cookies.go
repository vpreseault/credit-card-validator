package server

import (
	"strings"
)

func encodeHistoryCookie(history map[string]bool) string {
	cookie := ""
	for ccn, valid := range history {
		historyItem := ""
		if valid {
			historyItem += "s"
		} else {
			historyItem += "f"
		}
		historyItem += ccn

		cookie += historyItem + "|"
	}
	return cookie
}

func decodeHistoryCookie(cookie string) map[string]bool {
	history := map[string]bool{}
	historyItems := strings.Split(cookie, "|")
	for _, historyItem := range historyItems {
		if len(historyItem) < 2 {
			continue
		}
		valid := false
		if historyItem[0] == 's' {
			valid = true
		}
		ccn := historyItem[1:]
		history[ccn] = valid
	}
	return history
}