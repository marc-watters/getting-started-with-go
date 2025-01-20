package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Write a program which prompts the user to first enter a name, and then enter an
	// address.
	var (
		name    string
		address string
	)

	scanner := bufio.NewScanner(os.Stdin)

	for name == "" || address == "" {
		var input *string

		if name == "" {
			input = &name
			fmt.Print("\nEnter a name: ")
		} else if address == "" {
			input = &address
			fmt.Print("Enter an address: ")
		}

		if ok := scanner.Scan(); !ok {
			fmt.Fprintf(os.Stderr, "error scanning input: %s", scanner.Err())
			os.Exit(1)
		}

		*input = scanner.Text()
	}

	// Your program should create a map and add the name and address to the map using
	// the keys “name” and “address”, respectively.
	location := map[string]string{
		"name":    name,
		"address": address,
	}

	// Your program should use Marshal() to create a JSON object from the map, and
	// then your program should print the JSON object.
	object, err := json.Marshal(location)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshalling location: %s", err)
	}

	fmt.Printf("\n%s\n\n", object)
}
