package main

import (
	"fmt"
	"time"
)

/*
	This program is a demonstration of a race condition. A race condition occurs
	when memory is shared between threads or processes without synchronization.
	What this means in practice is the value of read/written memory is non-deterministic
	as there is no way to know in advance what values will be read or written.
*/

func main() {
	var variable string

	go func() {
		variable = "written to by routine 1"
	}()

	go func() {
		variable = "written to by routine 2"
	}()

	for variable == "" {
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println(variable)
}
