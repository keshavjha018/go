package main

import (
	"fmt"
)

func main() {

	// memory allocation
	// new() => memory init with zeros
	// make() => init by garbage vals

	// Pointers
	var ptr *int
	fmt.Println("Pointer referencing nowhere: ", ptr)

	num := 2
	var ptr2 = &num

	fmt.Print("Pointer: ", ptr2)             // prints address
	fmt.Println(", referencing to: ", *ptr2) // prints val stored
}
