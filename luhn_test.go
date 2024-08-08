package main

import "testing"

func TestIsValidLuhn(t *testing.T) {
	testCases := []struct {
		name, input string;
		expected bool;
    }{
    	{"valid credit card 1", "5555555555554444", true},
		{"valid credit card 2", "4012888888881881", true},
		{"valid credit card 3", "4111111111111111", true},
		{"valid credit card 4", "6011000990139424", true},
    	{"valid luhn number 1", "17893729974", true},
    	{"invalid luhn number 1", "17893729977", false},
		{"invalid credit card 1", "4012888888888888", false},
    }

    for _, tc := range testCases {
		// Variable Capturing
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := isValidLuhn(tc.input)

			if err != nil {
				t.Fatalf("isValidLuhn(%v) throws: %v", tc.input, err)
			}

			if result != tc.expected {
				t.Errorf("isValidLuhn(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})	
    }
}

func TestStringToIntSlice(t *testing.T) {
	testCases := []struct {
        input string;
		expected []int;
    }{
    	{"5555555555554444", []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4}},
		{"4012888888881881", []int{4, 0, 1, 2, 8, 8, 8, 8, 8, 8, 8, 8, 1, 8, 8, 1}},
		{"4111111111111111", []int{4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{"6011000990139424", []int{6, 0, 1, 1, 0, 0, 0, 9, 9, 0, 1, 3, 9, 4, 2, 4}},
    	{"17893729974", []int{1, 7, 8, 9, 3, 7, 2, 9, 9, 7, 4}},
    	{"17893729977", []int{1, 7, 8, 9, 3, 7, 2, 9, 9, 7, 7}},
		{"4012888888888888", []int{4, 0, 1, 2, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8}},
    }

    for _, tc := range testCases {
        result, err := stringToIntSlice(tc.input)

		if err != nil {
			t.Fatalf("stringToIntSlice(%v) throws: %v", tc.input, err)
		}

		if len(result) != len(tc.expected) {
			t.Fatalf("stringToIntSlice(%v) len(%v) != len(%v)", tc.input, result, tc.expected)
		}
		
		for i := 0; i < len(result); i++ {
			if result[i] != tc.expected[i] {
				t.Errorf("stringToIntSlice(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		}
    }
}

func TestTrimWhitespace(t *testing.T) {
	testCases := []struct {
        name, input, expected string;
    }{
    	{"separating spaces", "5555 5555 5555 4444", "5555555555554444" },
		{"trailing tab", "4012 8888 8888 1881	", "4012888888881881" },
		{"separating tabs", "4111	1111	1111	1111", "4111111111111111"},
		{"leading tab", "	6011000990139424", "6011000990139424"},
    }

    for _, tc := range testCases {
		// Variable Capturing
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := trimWhitespace(tc.input)
			
			if result != tc.expected {
				t.Errorf("trimWhitespace(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
    }
}