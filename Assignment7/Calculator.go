/*
	package : go.core/assignments

   Go program that prompts the user to enter a two numbers and operator,
   calculates the result for valid numbers, and produces the output":
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var operator string
	var number1, number2 float64

	// Taking inputs and assigning in number1, number2  and operator using Scanln
	fmt.Println("Enter the first number:")
	number1 = getNumber()

	fmt.Println("Enter the second number:")
	number2 = getNumber()

	fmt.Println("Enter The Operation :")
	fmt.Scanln(&operator)

	//calculating result
	res := calculateResult(number1, number2, operator)

	//prints result in console
	fmt.Printf("result is: %f\n", res)
}

// takes input converts string to float64 or prints errors
func getNumber() float64 {
	var number string
	fmt.Scanln(&number)
	num, err := strconv.ParseFloat(number, 64)
	if err != nil {
		fmt.Printf("Please Enter Integer '%s' is not supported\n", number)
		os.Exit(1)
	}
	return num
}

// performs operation based on user choice
func calculateResult(num1 float64, num2 float64, operator string) float64 {
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0.0 { // cheking if user is trying to divide by 0
			fmt.Println("error: you tried to divide by zero.")
			os.Exit(1)
		}
		result = num1 / num2

	default: //printing default msg is user gives invalid operator
		fmt.Println("Invalid Operator: Only '+' '-' '/' '*' is allowed")
		os.Exit(1)
	}
	return result
}
