package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var result int

	for result == 0 {
		fmt.Print("Enter a floating point number: ")

		var input string
		if _, err := fmt.Scan(&input); err != nil {
			fmt.Fprintf(os.Stderr, "\nerror: %v\n\n", err)
			continue
		}

		float, err := strconv.ParseFloat(input, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nerror: %v\n\n", err)
			continue
		}

		result = int(float)
	}

	fmt.Println(result)
}
