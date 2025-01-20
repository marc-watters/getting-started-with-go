package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	Write a Bubble Sort program in Go. The program should print the integers out
	on one line, in sorted order, from least to greatest. Use your favorite search
	tool to find a description of how the bubble sort algorithm works.
*/

func main() {
	// The program should prompt the user to type in a sequence of up to 10 integers.
	var (
		integers = []int{}
		scanner  = bufio.NewScanner(os.Stdin)
	)

	fmt.Print("Enter up to 10 integers: ")
	if ok := scanner.Scan(); !ok {
		fmt.Fprintf(os.Stderr, "error: %s", scanner.Err())
		os.Exit(1)
	}

	input := scanner.Text()
	tokens := strings.Fields(input)

	for i, v := range tokens {
		if i > 10 {
			break
		}

		int, err := strconv.Atoi(v)
		if err != nil {
			continue
		}

		integers = append(integers, int)
	}

	BubbleSort(integers)

	fmt.Println(integers)
}

// As part of this program, you should write a function called BubbleSort() which
// takes a slice of integers as an argument and returns nothing. The BubbleSort()
// function should modify the slice so that the elements are in sorted order.
func BubbleSort(ints []int) {
	for i := 0; i < len(ints); i++ {
		Swap(ints, i)
	}
}

// A recurring operation in the bubble sort algorithm is the Swap operation which
// swaps the position of two adjacent elements in the slice. You should write a
// Swap() function which performs this operation. Your Swap() function should take
// two arguments, a slice of integers and an index value i which indicates a
// position in the slice. The Swap() function should return nothing, but it should
// swap the contents of the slice in position i with the contents in position i+1.
func Swap(ints []int, i int) {
	for j := 0; j < len(ints)-1-i; j++ {
		if ints[j] > ints[j+1] {
			ints[j], ints[j+1] = ints[j+1], ints[j]
		}
	}
}
