/*
package: go.core/assignments
Go program  to prompt the user to enter their name,
reads the input, and then prints a greeting message with the provided name.
*/
package main

import (
	"fmt"
)

func main() {
	var name string // Declare a variable named 'name' of type string to store the user's input.

	fmt.Println("Enter your name") // Print a message to the console asking the user to enter their name.
	fmt.Scanf("%s", &name)         // Use Scanf to read a string input from the user and store it in the 'name' variable.

	fmt.Printf("Hello %s!\n", name) // Print a greeting message to the console using the value stored in 'name'.
}
