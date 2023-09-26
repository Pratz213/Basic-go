/*
	package : go.core/assignments

   Go program that prompts the user to enter a sequence of numbers,
   calculates the sum of valid numbers, and produces the output when the user enters "proceed":
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PROCEED = "proceed"
)

func main() {

	//Declaring variables
	var numbers string
	var total float64

	fmt.Println("enter number or enter word 'proceed'")
	for {
		fmt.Scanln(&numbers) //Taking  input and assigning in numbers variable

		//if input is proceed it exit from loop.
		if strings.ToLower(numbers) == PROCEED {
			break
		}

		//converts string to float64 and store in nums or store error if there are any.
		nums, err := strconv.ParseFloat(numbers, 64)
		if err != nil { //prints error and terimates program
			panic("only number or word 'proceed' is allowed")
		}
		total += nums //calculates total each time user enter a valid number
	}

	//prints total on console
	fmt.Println("\nTotal : ", total)
}
