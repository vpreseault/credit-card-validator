package server

import (
	"testing"
)
	

func TestAddHistoryItem(t *testing.T) {
	testCases := []struct {
		name string;
		inputHistory, expectedHistory []HistoryItem;
		ccn string;
		valid bool;
	}{
		{
			"add item", 
			[]HistoryItem{}, 
			[]HistoryItem{{"1111", true}}, 
			"1111", 
			true,
		},
		{
			"add 6th item", 
			[]HistoryItem{{"1111", true}, {"2222", false}, {"3333", true}, {"4444", false}, {"5555", true}}, 
			[]HistoryItem{{"2222", false}, {"3333", true}, {"4444", false}, {"5555", true}, {"6666", true}}, 
			"6666", 
			true,
		},
		{
			"add already present item", 
			[]HistoryItem{{"1111", true}, {"2222", false}, {"3333", true}},
			[]HistoryItem{{"2222", false}, {"3333", true}, {"1111", true}},
			"1111",
			true,
		},
		{
			"add already present 6th item",
			[]HistoryItem{{"1111", true}, {"2222", false}, {"3333", true}, {"4444", false}, {"5555", true}},
			[]HistoryItem{{"1111", true}, {"2222", false}, {"4444", false}, {"5555", true}, {"3333", true}},
			"3333",
			true,
		},
	}

	for _, tc := range testCases {
		// Variable Capturing
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := addHistoryItem(tc.inputHistory, tc.ccn, tc.valid)
			
			for i, item := range result {
				if item != tc.expectedHistory[i] {
					t.Errorf("item at index %v = %v; want %v", i, item, tc.expectedHistory[i])
				}
			}	
		})
	}
}