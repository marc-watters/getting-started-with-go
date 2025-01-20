package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	Write a program which reads information from a file and represents it in a
	slice of structs. Assume that there is a text file which contains a series of names.
	Each line of the text file has a first name and a last name, in that order,
	separated by a single space on the line.

names.txt:

john doe
jane smith
billy bob
alice cooper
thisisaverylongfirstname thisisaverylonglastname

*/

var names = make([]name, 0)

// Your program will define a name struct which has two fields, fname for the
// first name, and lname for the last name.
type name struct {
	fname string
	lname string
}

func main() {
	// Your program should prompt the user for the name of the text file.
	fmt.Print("\nEnter the name of the file containing the list of names: ")
	inputScanner := bufio.NewScanner(os.Stdin)
	if ok := inputScanner.Scan(); !ok {
		fmt.Fprintf(os.Stderr, "error: %s", inputScanner.Err())
		os.Exit(1)

	}

	fileName := inputScanner.Text()
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %s", err)
	}
	defer file.Close()

	// Your program will successively read each line of the text file and create a struct
	// which contains the first and last names found in the file.
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fullName := strings.Fields(line)

		// Each field will be a string of size 20 (characters).
		for i := range 2 {
			if len(fullName[i]) > 20 {
				fullName[i] = fullName[i][:21]
			}
		}

		// Each struct created will be added to a slice, and after all lines have been
		// read from the file, your program will have a slice containing one struct for
		// each line in the file.
		newName := name{fullName[0], fullName[1]}
		names = append(names, newName)
	}

	// After reading all lines from the file, your program should iterate through
	// your slice of structs and print the first and last names found in each struct.
	for i, name := range names {
		fmt.Printf("[%d] - first name: %s\tlast name: %s\n", i, name.fname, name.lname)
	}
}
