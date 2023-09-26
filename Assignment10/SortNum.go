// This program reads numbers from the user until they enter "proceed".
// It calculates the total of the entered numbers and displays the sorted numbers along with the total.
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	PROCEED = "proceed"
)

func main() {
	var input string
	var nums []float64
	var total float64

	fmt.Println("Enter a number or type 'proceed':")
	fmt.Scanln(&input)

	for strings.ToLower(input) != PROCEED {
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input")
			break
		}
		total += num // Calculate total each time a valid number is entered
		nums = append(nums, num)
		fmt.Scanln(&input)
	}

	//after user takes proceed prints sorted array and total of all elements of array
	if strings.ToLower(input) == PROCEED && len(nums) > 0 {
		sort.Float64s(nums)
		fmt.Println("Sorted numbers:", nums, "\nTotal:", total)
	}
}
