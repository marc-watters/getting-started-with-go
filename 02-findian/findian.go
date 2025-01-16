package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		input string
		found bool
		err   error

		reader = bufio.NewReader(os.Stdin)
	)

	for input == "" {
		fmt.Print("Enter a string: ")

		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			continue
		}

		input = strings.ToLower(input)
		input = strings.TrimSpace(input)

		if strings.HasPrefix(input, "i") {
			if strings.Contains(input, "a") {
				if strings.HasSuffix(input, "n") {
					found = true
				}
			}
		}
	}

	if found {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
