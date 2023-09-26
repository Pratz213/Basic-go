/*
package : go.core/assignments

checks if the year is a leap year, and prints whether it is or isn't.
It uses the IsLeapYear function to determine leap years based on the established rules.
*/
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var inputDate string
	//promt user for date
	fmt.Print("Enter a date (YYYY-MM-DD): ")
	fmt.Scanln(&inputDate)

	//parsing string to date
	dateFormat := "2006-01-02" // Specify the layout of the expected date string
	parsedDate, err := time.Parse(dateFormat, inputDate)
	if err != nil {
		fmt.Print("Invalid Date input\n")
		os.Exit(1)
	}

	year := parsedDate.Year() //extract year from date
	isLeapYear := IsLeapYear(year)

	if isLeapYear {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}

// checks leap year or not
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}
