package main
import "fmt"

// The greet function takes a string parameter called name and prints a personalized greeting message.
func greet(name string) {
	fmt.Println("Hello,", name)
}
func main() {
	//Here we are calling the greet function with different arguments to print personalized greetings.
	greet("Alice")
	greet("Bob")
}
