/*
package : go.core/assignments

Go program takes two positive whole numbers as input, calculates and prints their multiplication table.
The getNumber() function retrieves a positive integer from the console, and the printMultiplicationTable()
function generates and displays the multiplication table for the first input number up to the second input number.
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter positive whole num 1")
	num1 := getNumber()
	fmt.Println("Enter positive whole num 2")
	num2 := getNumber()
	printMultiplicationTable(num1, num2)
}

// prints table
func printMultiplicationTable(num1, num2 int) {
	fmt.Println("Multiplication Table of ", num1, " upto", num2)
	for i := 1; i <= num2; i++ {
		fmt.Println(num1, " x ", i, " = ", num1*i)
	}
}

// gets whole number from console and parse it in int.
func getNumber() int {
	var input string
	fmt.Scanln(&input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if num < 1 {
		fmt.Println("enter whole positive number only ")
		os.Exit(1)
	}
	return num
}
