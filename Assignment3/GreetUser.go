/*
package : go.core/assignments

Go program reads a command-line argument provided by the user and prints a greeting message with the provided argument.
If the user provides more than one argument or no arguments at all,
it displays an error message indicating that only one input is allowed.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var name string = os.Args[1] //storing command-line argument in the 'name' variable.

	if len(os.Args) <= 2 { // Checks if the number of command-line arguments is less than two (program name + 1 input arguments).
		fmt.Printf("Only 1 input is allowed. You entered %d argument(s).\n", len(os.Args)-1) // Print an error message.
		fmt.Printf("Hello %s!\n", name)                                                      // If only one input argument is provided, print a greeting message using the value stored in 'name'.
	}
}
