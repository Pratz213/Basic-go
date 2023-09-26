package main

/*
	package : go.core/assignments

   Go program reads two inputs provided by the user: a number1 string and a number2 string.
   it checks if the both arguments can be converted string to int or it throws error if not.
   calculates sum of both values and prints result
*/

import (
	"fmt"
	"strconv"
)

func main() {
	//Declaring variables
	var number1, number2 string

	// Taking input and assigning to num1 using Scanln
	fmt.Println("Enter the first number:")
	fmt.Scanln(&number1)

	// Converting String to int and calculating the result
	num1, err := strconv.Atoi(number1)
	if err != nil {
		panic(err)
	}

	// Taking input and assigning to num2 using Scanln
	fmt.Println("Enter the second number:")
	fmt.Scanln(&number2)
	num2, err := strconv.Atoi(number2)
	if err != nil {
		panic(err)
	}

	//printing result
	fmt.Printf("num1 = %d\nnum2 = %d\n--------\nTotal = %d\n", num1, num2, num1+num2)
}
