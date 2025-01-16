package main

import "fmt"

func main() {
	var f float32
	var i int

	fmt.Println("Enter a floating point number:")

	_, err := fmt.Scan(&f)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	i = int(f)
	fmt.Println("As integer:", i)
}
