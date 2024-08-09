package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vpreseault/credit-card-validator/luhn"
)

func main() {
    if len(os.Args) > 1 {
		input := strings.Join(os.Args[1:], " ")
		checkInput(input)
	} else {
		fileInfo, err := os.Stdin.Stat()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error stating stdin:", err)
			os.Exit(1)
		}
		
		scanner := bufio.NewScanner(os.Stdin)
		if (fileInfo.Mode() & os.ModeCharDevice) != 0 {
			fmt.Println("Please enter a credit card number: ")
			if scanner.Scan() {
				line := scanner.Text()
				checkInput(line)
			}
		} else {
			for scanner.Scan() {
				line := scanner.Text()
				checkInput(line)
			}
		}

        if err := scanner.Err(); err != nil {
            fmt.Fprintln(os.Stderr, "Error reading input:", err)
            os.Exit(1)
        }
	}
}

func checkInput(input string) {
	result, err := luhn.IsValidLuhn(input)
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