package main

import "testing"

func TestIsValidLuhn(t *testing.T) {
	testCases := []struct {
        s string;
		expected bool;
    }{
    	{"5555555555554444", true},
		{"4012888888881881", true},
		{"4111111111111111", true},
		{"6011000990139424", true},
    	{"17893729974", true},
    	{"17893729977", false},
		{"4012888888888888", false},
    }

    for _, tc := range testCases {
        result, err := isValidLuhn(tc.s)

		if err != nil {
			t.Fatalf("stringToIntSlice(%v) throws: %v", tc.s, err)
		}

        if result != tc.expected {
            t.Errorf("isValidLuhn(%v) = %v; want %v", tc.s, result, tc.expected)
        }
    }
}

func TestStringToIntSlice(t *testing.T) {
	testCases := []struct {
        s string;
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
        result, err := stringToIntSlice(tc.s)

		if err != nil {
			t.Fatalf("stringToIntSlice(%v) throws: %v", tc.s, err)
		}

		if len(result) != len(tc.expected) {
			t.Fatalf("stringToIntSlice(%v) len(%v) != len(%v)", tc.s, result, tc.expected)
		}
		
		for i := 0; i < len(result); i++ {
			if result[i] != tc.expected[i] {
				t.Errorf("stringToIntSlice(%v) = %v; want %v", tc.s, result, tc.expected)
			}
		}
    }
}

func TestTrimWhitespace(t *testing.T) {
	testCases := []struct {
        s, expected string;
    }{
    	{"5555 5555 5555 4444", "5555555555554444" },
		{"4012 8888 8888 1881	", "4012888888881881" },
		{"4111	1111	1111	1111", "4111111111111111"},
		{"	6011000990139424", "6011000990139424"},
    }

    for _, tc := range testCases {
        result := trimWhitespace(tc.s)
		
		if result != tc.expected {
            t.Errorf("isValidLuhn(%v) = %v; want %v", tc.s, result, tc.expected)
        }
    }
}