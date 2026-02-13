package main

import (
	"fmt"
)
//This is a package-level variable declaration
// It can be accessed from any function within the same package
var packageLevelVariable string = "John"
func main(){
	var name = "John"
	//Short variable declaration, it's possible only inside a function, 
	// it infers the type of the variable from the value assigned to it
	age := 30
	var height float32 = 1.70
	fmt.Println(name, "is", age, "years old, his height is", height, "meters.");
	//John is 30 years old, his height is 1.75 meters
	// Output is spaced by default because of the commas in the Println function
	fmt.Printf("%s is %d years old, his height is %.2f meters.\n", name, age, height)
}