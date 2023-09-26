/*
	package : go.core/assignments

   Go program to allow user to input numbers or commands like 'count', 'mean', 'min', and 'max'.
   It calculates and displays the count, mean, minimum, or maximum values of entered numbers using switch statements and separate functions,
   while ensuring valid inputs and providing error handling.
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var numbers []float64

	fmt.Print("Enter a number, or word ('count', 'mean', 'min', 'max') or operator ('+','-','*','/') \n")
	//to take inputs multiple times
	for {
		var num string
		fmt.Scan(&num)

		switch num {
		case "count":
			fmt.Printf("Count: %d\n", len(numbers))
			return
		case "mean":
			printMean(numbers)
			return
		case "min":
			printMin(numbers)
			return
		case "max":
			printMax(numbers)
			return
		case "+":
			printSum(numbers)
		default:
			numbers = append(numbers, parseNumber(num)) //appending numbers .
		}
	}
}

// checks for valid number or throws error
func parseNumber(num string) float64 {
	number, err := strconv.ParseFloat(num, 64) //converting string to float64
	if err != nil {                            //throwing error for invalid inputs and exits
		fmt.Println("Invalid input.")
		os.Exit(1)
	}
	return number
}

// sum of all number
func printSum(numbers []float64) {
	var result float64
	for _, num := range numbers {
		result += num
	}
	fmt.Println("sum: ", result)
	os.Exit(0)
}

// calculating and printing Min number
func printMin(numbers []float64) {
	if len(numbers) == 0 { //throws error and exits if user enters min before appending numbers.
		fmt.Print("enter numbers first\n")
		os.Exit(1)
	}
	sort.Float64s(numbers)
	fmt.Printf("Minimum: %.2f\n", numbers[0])
}

// calculating and printing Max number
func printMax(numbers []float64) {
	if len(numbers) == 0 {
		fmt.Print("enter numbers first\n")
		os.Exit(1)
	}
	sort.Float64s(numbers)
	fmt.Printf("Maximum: %.2f\n", numbers[len(numbers)-1])
}

// calculating and printing Mean
func printMean(numbers []float64) {
	if len(numbers) == 0 {
		fmt.Printf("Mean: %.f\n", 0.0)
		os.Exit(0)
	}
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	mean := sum / float64(len(numbers))
	fmt.Printf("Mean: %.2f\n", mean)
}
