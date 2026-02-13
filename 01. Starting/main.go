/*
Package main is the entry point of the program.
It contains the main function which is the starting point of execution.
*/
package main

import (
	"fmt"
	"github.com/fatih/color"
)

/*
The main function is the entry point of the program.
the function name must be main and it must not take any arguments and must not return any value.
*/
func main() {
	fmt.Println("Hello, World!")
	color.Red("This is a red text")
}
// go mod init hello-go -- This command is used to initialize a new Go module in the current directory. It creates a go.mod file which is used to manage dependencies for the project.
// go get github.com/fatih/color -- This command is used to install the color package which is used to print colored text in the terminal.