package main

import "fmt"

func main() {
	age := 3
	//age = "three" // This will cause a compile-time error because "three" is a string, not an integer.
	fmt.Println("Age ", age)
	var size int
	fmt.Println("Size ", size) // This will print "Size 0" because the default value of an int is 0.
	var name string
	fmt.Println("Name ", name) // This will print "Name " because the default value of a string is an empty string.
	var isActive bool
	fmt.Println("Is Active ", isActive) // This will print "Is Active false" because the default value of a bool is false.
	
	
	//This is a compile-time error because the result of the expression is a float64, 
	 // and we cannot assign it to an int variable without explicit conversion.
	//approximatedDays := 365.25 * age; ⬅️



	// This will work because we are explicitly converting the result to an int.
	//Go doesn't convert types implicitly,
	//  so we need to use a type conversion to assign the result to an int variable.
	approximatedDays := int(365.25 * float64(age)) 
	fmt.Println("Approximated Days ", approximatedDays)
}
