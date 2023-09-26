/*
	package : go.core/assignments

   Go program to allow user to input numbers or commands like 'count', 'mean', 'min','even','odd', and 'max'.
   It calculates and displays the count, mean, minimum, or maximum or and gets even and odd values of entered numbers using switch statements and separate functions.
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var numbers []int

	fmt.Print("Enter a number, or word ('count', 'mean', 'min', 'max','even','odd') or operator ('+')  \n")
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
		case "even":
			printEven(numbers)
		case "odd":
			printOdd(numbers)
		default:
			numbers = append(numbers, parseNumber(num)) //appending numbers .
		}
	}
}

// checks for valid number or throws error
func parseNumber(num string) int {
	number, err := strconv.Atoi(num) //converting string to float64
	if err != nil {                  //throwing error for invalid inputs and exits
		fmt.Println("Invalid input.")
		os.Exit(1)
	}
	return number
}

// sum of all number
func printSum(numbers []int) {
	var result int
	for _, num := range numbers {
		result += num
	}
	fmt.Println("sum: ", result)
	os.Exit(0)
}

// calculating and printing Min number
func printMin(numbers []int) {
	if len(numbers) == 0 { //throws error and exits if user enters min before appending numbers.
		fmt.Print("enter numbers first\n")
		os.Exit(1)
	}
	sort.Ints(numbers)
	fmt.Printf("Minimum: %d\n", numbers[0])
}

// calculating and printing Max number
func printMax(numbers []int) {
	if len(numbers) == 0 {
		fmt.Print("enter numbers first\n")
		os.Exit(1)
	}
	sort.Ints(numbers)
	fmt.Printf("Maximum: %d\n", numbers[len(numbers)-1])
}

// calculating and printing Mean
func printMean(numbers []int) {
	if len(numbers) == 0 {
		fmt.Printf("Mean: %.f\n", 0.0)
		os.Exit(0)
	}
	var sum int
	for _, num := range numbers {
		sum += num
	}
	mean := sum / len(numbers)
	fmt.Printf("Mean: %d\n", mean)
}

// calculating and printing even numbers
func printEven(numbers []int) {
	var count int
	var nums []int
	for _, num := range numbers {
		if num%2 == 0.0 {
			count++
			nums = append(nums, num)
		}
	}
	fmt.Println("Even Numbers: ", nums, "\nCount: ", count)
	os.Exit(0)
}

// calculating and prints odd numbers
func printOdd(numbers []int) {
	var count int
	var nums []int
	for _, num := range numbers {
		if num%2 != 0 {
			count++
			nums = append(nums, num)
		}
	}
	fmt.Println("Odd Numbers: ", nums, "\nCount: ", count)
	os.Exit(0)
}
