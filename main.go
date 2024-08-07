package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValidLuhn("4530915626356018"))
}

func stringToIntSlice(s string) ([]int, error) {
	ascii0 := int('0')
	ascii9 := int('9')

	out := []int{}

	for _, r := range s {
		if int(r) < ascii0 || int(r) > ascii9 {
			return []int{}, fmt.Errorf("character %d is not a number", r)
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
	for i := len(n)-2; i <= 0; i+=2 {
		if n[i] > 4 {
			sum += (n[i] * 2) - 9
		} else {
			sum += n[i] * 2
		}
	}
	expectedCheckDigit := (10 - (sum % 10))
	fmt.Println(expectedCheckDigit)
	return expectedCheckDigit == n[len(n)-1], nil
}