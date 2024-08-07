package main

import (
	"fmt"
	"strings"
	"unicode"
)

func trimWhitespace(s string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsSpace(r) {
            return -1
        }
        return r
    }, s)
}

func stringToIntSlice(s string) ([]int, error) {
	ascii0 := int('0')

	out := []int{}

	for i, r := range s {
		if !unicode.IsDigit(r) {
			return []int{}, fmt.Errorf("character [%v] at index %d is not a number", string(r), i)
		}
		out = append(out, int(r) - ascii0)
	}

	return out, nil
}

func isValidLuhn(s string) (bool, error) {

	n, err := stringToIntSlice(s)
	if err != nil {
		return false, err
	}

	sum := 0
	for i := len(n)-2; i >= 0; i-=2 {
		if n[i] > 4 {
			sum += (n[i] * 2) - 9
		} else {
			sum += n[i] * 2
		}

		if i - 1 >= 0 {
			sum += n[i-1]
		}
	}
	expectedCheckDigit := (10 - (sum % 10))
	return expectedCheckDigit == n[len(n)-1], nil
}