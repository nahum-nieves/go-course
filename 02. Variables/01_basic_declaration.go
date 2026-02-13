package main

import (
	"fmt"
)

func main() {
	// Variables are used to store values in memory. They have a name, a type, and a value.
	// the type of a variable can be inferred from the value assigned to it,
	// or it can be explicitly declared.
	var info = "Hello, World!"
	fmt.Println(info)
	info = "Hello, Go!"
	fmt.Printf(info)
	// This will cause a compile-time error because the variable "info"
	//  is of type string and cannot be assigned an integer value.
	//info = 20;
}
