/*

package : go.core/assignments

Go code to This Go program prompts the user to input start date and end date .
calculates difference and prints.

*/

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Prompt the user to enter the start date and get the parsed date
	fmt.Println("enter start date in format: YYYY-MM-DDTHH:MM:SS")
	startDate := getDate()

	// Prompt the user to enter the end date and get the parsed date
	fmt.Println("enter end date in format: YYYY-MM-DDTHH:MM:SS")
	endDate := getDate()

	// Calculate the duration between the two dates
	years, months, days, hours, minutes := CalculateDuration(startDate, endDate)

	// Print the calculated difference
	fmt.Printf("Difference between the two date-times: %d years %d months %d days %d hours %d minutes\n", years, months, days, hours, minutes)
}

// gets inpput and parsed date
func getDate() time.Time {
	var inputDate string
	fmt.Scanln(&inputDate)
	dateFormat := "2006-01-02T15:04:05" //parsing format

	parsedDate, err := time.Parse(dateFormat, inputDate)
	if err != nil {
		fmt.Println("invalid input")
		os.Exit(1)
	}
	return parsedDate
}

// calculates duration between two dates
func CalculateDuration(startDate, endDate time.Time) (years, months, days, hours, minutes int) {

	// Checks startDate is earlier than endDate
	if startDate.After(endDate) {
		startDate, endDate = endDate, startDate
	}

	// Calculate the difference in years, months, days,hours and minutes
	years = endDate.Year() - startDate.Year()
	months = int(endDate.Month()) - int(startDate.Month())
	days = endDate.Day() - startDate.Day()
	hours = endDate.Hour() - startDate.Hour()
	minutes = endDate.Minute() - startDate.Minute()

	fmt.Printf("TESTING: %d years %d months %d days %d hours %d minutes\n", years, months, days, hours, minutes)

	// Adjust negative minutes by subtracting an hour and adding 60 minutes
	if minutes < 0 {
		hours--
		minutes += 60
	}

	// Adjust negative hours by subtracting a day and adding 24 hours
	if hours < 0 {
		days--
		hours += 24
	}

	// Adjust negative days by subtracting a month and adding last month's days
	if days < 0 {
		months--
		lastMonthEndDate := endDate.AddDate(0, -1, 0)
		days += lastMonthEndDate.Day()
	}

	// Adjust negative months by subtracting a year and adding 12 months
	if months < 0 {
		years--
		months += 12
	}
	return years, months, days, hours, minutes
}
