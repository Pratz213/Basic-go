package main

import "fmt"

/*
Problem 1: Palindrome Check
Write a function that takes a string as input and returns true if the string is a palindrome
(reads the same forwards and backwards), and false otherwise.
*/

/*
	Using map and returning all string palidrome or not.
*/

func main() {
	namesList := getList()
	printPalindrome(namesList)
}

func printPalindrome(namesList map[string]bool) {
	fmt.Println("--------------------\nPlidrome List: ")
	for item := range namesList {
		if isPalidrome(item) {
			namesList[item] = true
		}
	}

	for item := range namesList {
		if namesList[item] {
			fmt.Println(item)
		}
	}
}

func isPalidrome(inputString string) bool {
	var palidrome = false
	var result string
	for _, v := range inputString {
		result = string(v) + result
	}
	if result == inputString {
		palidrome = true
	}
	return palidrome
}

func getList() map[string]bool {
	var namesList = make(map[string]bool)
	fmt.Println("Enter String to check Palidrome or enter 'Done'")
	var input string
	for {
		fmt.Scanln(&input)
		if input == "done" {
			return namesList
		}
		namesList[input] = false
	}
}
