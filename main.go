package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: ccv <your creditcard number>")
        return
    }

    input := strings.Join(os.Args[1:], " ")

	result, err := isValidLuhn(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	if result {
		fmt.Printf("%v is a valid credit card number.\n", input)
	} else {
		fmt.Printf("%v is not a valid credit card number.\n", input)
	}
}