package main
import "fmt"

const message = "Hello,"
// The greet function takes a string parameter called name and prints a personalized greeting message.
func greet(name string) {
	fmt.Println(message, name)
}
func main() {
	
	//Here we are calling the greet function with different arguments to print personalized greetings.
	greet("Alice")
	greet("Bob")
}
