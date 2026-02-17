package constants
import "fmt"

func main(){
	//Constants are immutable values that cannot be changed after they are declared
	//their value must be assigned at the time of declaration and cannot be left uninitialized
	const pi = 3.14
	// constants or variables declared with 
	// PascalCase are exported and can be accessed from other packages
	const Greeting = "Hello, World!"
	const isGoFun = true

	fmt.Println("Pi:", pi)
	fmt.Println("Greeting:", Greeting)
	fmt.Println("Is Go Fun?", isGoFun)
}