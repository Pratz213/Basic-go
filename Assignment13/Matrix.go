/*

package : go.core/assignments

Go code prompts the user for the number of rows , columns, and  character to print.
The getInput function handles input validation, converting the user's input to an integer.
The printMatrix function prints the specified character in rows and columns to create the matrix.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Prompt for inputs
	fmt.Print("Enter the number of rows: ")
	numRows := getInput()
	fmt.Print("Enter the number of cols: ")
	numCols := getInput()

	fmt.Print("Enter a character to be used for display: ")
	var character string
	fmt.Scanln(&character)

	// print the rectangular matrix
	printMatrix(character, numRows, numCols)

}

// Takes input from console
func getInput() (num int) {
	var number string
	fmt.Scanln(&number)
	num, err := strconv.Atoi(number)
	if err != nil {
		fmt.Printf("invalid number: %s\n", number)
		os.Exit(1)
	}
	if num <= 0 {
		fmt.Printf("number should be positive integer\n")
	}
	return num
}

// Prints character
func printMatrix(character string, numRows, numCols int) {
	char := strings.Split(character, "")
	chars := string(character[0])

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			fmt.Printf("%c ", character[0])
		}
		fmt.Println()

		fmt.Println("-------------------")

		for i := 0; i < numRows; i++ {
			for j := 0; j < numCols; j++ {
				fmt.Print(chars, " ")
			}
			fmt.Println()
		}
		fmt.Println("-------------------")
		for i := 0; i < numRows; i++ {
			for j := 0; j < numCols; j++ {
				fmt.Print(char[0], " ")
			}
			fmt.Println()
		}
		fmt.Println("-------------------")

		for _, value := range character {
			fmt.Println(string(value))
		}

		fmt.Println("-------------------")

		data := []string{character, "Demo", "niwnjc"}

		for _, char := range data {
			fmt.Println(char)
			for _, ch := range char {
				fmt.Printf("%q\n ", ch)
			}
		}

	}
}
