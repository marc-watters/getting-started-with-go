package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Write a program which prompts the user to enter integers and stores the
// integers in a sorted slice.
func main() {
	/*
		The program should be written as a loop. Before entering the loop, the
		program should create an empty integer slice of size (startLen) 3.
		The slice must grow in size to accommodate any number of integers which the
		user decides to enter.
	*/
	var (
		ints    = make([]int, 3)
		scanner = bufio.NewScanner(os.Stdin)
	)

	// The program should only quit (exiting the loop) when the user enters the
	// character ‘X’ instead of an integer.
	for input := ""; input != "X"; {

		// During each pass through the loop, the program prompts the user to enter
		// an integer to be added to the slice.
		fmt.Print("\nEnter an integer: ")
		if ok := scanner.Scan(); !ok {
			fmt.Fprintf(os.Stderr, "error: %s", scanner.Err())
			os.Exit(1)
		}

		// convert to integer
		input = scanner.Text()
		integer, err := strconv.Atoi(input)
		if err != nil {
			if input != "X" {
				fmt.Fprintf(os.Stderr, "invalid input: %s\n", err)
			}
			continue
		}

		// The program adds the integer to the slice, sorts the slice, and prints the
		// contents of the slice in sorted order.
		ints = append(ints, integer)
		sort.Ints(ints)
		fmt.Printf("Sorted: %v\n", ints[3:])
	}
}
