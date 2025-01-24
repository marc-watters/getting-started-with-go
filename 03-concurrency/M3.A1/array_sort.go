package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/* Write a program to sort an array of integers. */
func main() {
	const (
		prompt      = "Enter a list of 4 or more space separated integers"
		minIntCount = 4
		listCount   = 4
	)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// The program should prompt the user to input a series of integers.
		fmt.Printf("%s: ", prompt)
		if ok := scanner.Scan(); !ok {
			fmt.Fprintf(os.Stderr, "error reading input: %v", scanner.Err())
			os.Exit(1)
		}

		input := strings.Fields(scanner.Text())
		if len(input) < minIntCount {
			fmt.Fprint(os.Stderr, "minimum of 4 numbers required\n\n")
			continue
		}

		// The program should partition the array into 4 parts, each of which is
		// sorted by a different goroutine. Each partition should be of approximately
		// equal size.
		lists := make([][]int64, listCount)
		for i, val := range input {
			if int, err := strconv.ParseInt(val, 10, 32); err == nil {
				lists[i%listCount] = append(lists[i%listCount], int)
			} else {
				fmt.Fprintf(os.Stderr, "invalid integer: %v", err)
				continue
			}
		}

		channel := make(chan []int64)
		for i := range listCount {
			go func(list []int64, c chan []int64) {
				// Each goroutine which sorts ¼ of the array should print the subarray
				// that it will sort.
				fmt.Printf("list %d: %v ", i, list)
				slices.Sort(list)
				fmt.Printf("sorted: %v\n", list)
				c <- list
			}(lists[i], channel)
		}

		for range listCount {
			<-channel
		}

		// Then the main goroutine should merge the 4 sorted subarrays into one
		// large sorted array.
		sorted := make([]int64, 0)
		for i := 0; i < listCount; i++ {
			sorted = append(sorted, lists[i]...)
		}
		slices.Sort(sorted)

		//  When sorting is complete, the main goroutine should print the entire sorted list.
		fmt.Printf("\nsorted: %v\n", sorted)
	}
}
