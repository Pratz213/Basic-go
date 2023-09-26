/*
	package : go.core/assignments

   Go program reads two command-line arguments provided by the user: a template string and a name string.
   It checks if both arguments are provided, and if not, it panics and displays an error message.
   If both arguments are provided, it checks if the template string contains the constant PLACEHOLDER.
   If the placeholder is present, it replaces all occurrences of PLACEHOLDER with the name string and prints the result.
   If the placeholder is not present in the template string, it prints a message indicating that.
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

// Declaring constants
const (
	PLACEHOLDER string = "$name$"
)

func main() {
	// Declaring Variable to store strings.
	var template string
	var name string

	// taking Command-line arguments and assigning it to variables.
	template = os.Args[1]
	name = os.Args[2]

	// Checks If the number of command-line arguments is less than 3 (program name + 2 input arguments)
	if len(os.Args) != 3 {
		// Panic :: terminates Program and print an error message if arguments not equals to three.
		panic("Only 2 inputs are allowed")
	} else if strings.Contains(template, PLACEHOLDER) { //checks if the template string contains the constant PLACEHOLDER,
		fmt.Println(strings.ReplaceAll(template, PLACEHOLDER, name))
	} else {
		fmt.Println("PLACEHOLDER not present")
	}
}
