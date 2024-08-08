package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    if len(os.Args) < 2 {
        scanner := bufio.NewScanner(os.Stdin)

        for scanner.Scan() {
            line := scanner.Text()
            checkInput(line)
        }

        if err := scanner.Err(); err != nil {
            fmt.Fprintln(os.Stderr, "Error reading input:", err)
            os.Exit(1)
        }
    } else {
		input := strings.Join(os.Args[1:], " ")
		checkInput(input)
	}
}

func checkInput(input string) {
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