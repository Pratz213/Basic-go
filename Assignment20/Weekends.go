/*
This Go program accepts start and end dates,
calculates and prints the number of weekend days (Saturday and Sunday) between them using the countWeekendDays function.
*/
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// Accept input for start date and end date
	fmt.Println("Enter the start date (in the format YYYY-MM-DD):")
	startDate := getDate()

	fmt.Println("Enter the end date (in the format YYYY-MM-DD):")
	endDate := getDate()

	weekendDays := countWeekendDays(startDate, endDate)
	fmt.Println("Number of weekend days between the given dates:", weekendDays)
}

// counts the number of weekend days  between two dates.
func countWeekendDays(startDate, endDate time.Time) int {

	// checks startDate is earlier than endDate
	if startDate.After(endDate) {
		startDate, endDate = endDate, startDate
	}

	weekendCount := 0
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		weekday := d.Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			weekendCount++
		}
	}
	return weekendCount
}

// gets inpput and parsed date
func getDate() time.Time {
	var inputDate string
	fmt.Scanln(&inputDate)
	dateFormat := "2006-01-02" //parsing format 

	parsedDate, err := time.Parse(dateFormat, inputDate)
	if err != nil {
		fmt.Println("invalid input")
		os.Exit(1)
	}
	return parsedDate
}
