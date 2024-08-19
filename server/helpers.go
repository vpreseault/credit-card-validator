package server

func addHistoryItem(history []HistoryItem, ccn string, valid bool) []HistoryItem {

	for i, item := range history {
		if item.CCN == ccn {
			history = append(history[:i], history[i+1:]...)
		}
	}

	history = append(history, HistoryItem{CCN: ccn, Valid: valid})
	
	if len(history) > 5 {
		history = history[1:]
	}

	return history
}