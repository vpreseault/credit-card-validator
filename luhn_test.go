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
        result, _ := isValidLuhn(tc.s)
        if result != tc.expected {
            t.Errorf("isValidLuhn(%v) = %v; want %v", tc.s, result, tc.expected)
        }
    }
}