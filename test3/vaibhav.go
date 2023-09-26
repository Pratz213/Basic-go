package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/slices"
)

// Define the list of holidays (hardcoded)
var holidays = []time.Time{
	time.Date(2023, time.August, 7, 0, 0, 0, 0, time.UTC),
	time.Date(2023, time.July, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2023, time.December, 25, 0, 0, 0, 0, time.UTC),
}

func addBusinessDays(inputDate time.Time, daysToAdd int) time.Time {
	startDay := inputDate

	// Calculate the number of weekends within the given duration
	weekends := daysToAdd / 5 * 2

	// Adjust daysToAdd to account for weekends
	adjustedDaysToAdd := daysToAdd + weekends

	// Start with the initial input date and add the adjusted days
	currentDate1 := startDay.AddDate(0, 0, adjustedDaysToAdd)

	for _, holiday := range holidays {
		if (holiday.Equal(startDay) || holiday.After(startDay)) && (holiday.Before(currentDate1) || holiday.Equal(currentDate1)) {
			currentDate1 = currentDate1.AddDate(0, 0, 1)
			fmt.Println(currentDate1)
		}
	}

	// Check if the target date falls on a weekend (Saturday or Sunday), and if so, adjust it
	if currentDate1.Weekday() == time.Saturday {
		currentDate1 = currentDate1.AddDate(0, 0, 2) // Move to the next Monday
	} else if currentDate1.Weekday() == time.Sunday {
		currentDate1 = currentDate1.AddDate(0, 0, 1) // Move to the next Monday
	} else if slices.Contains(holidays, currentDate1) {
		currentDate1 = currentDate1.AddDate(0, 0, 1)
	}

	// Subtract holidays from the working days

	return currentDate1
}

func main() {
	inputDate := time.Date(2023, time.August, 4, 0, 0, 0, 0, time.UTC) // Your input date
	daysToAdd := 1                                                     // Number of business days to add

	resultDate := addBusinessDays(inputDate, daysToAdd)
	fmt.Println("Input Date:", inputDate.Format("2006-01-02"))
	fmt.Println("Result Date:", resultDate.Format("2006-01-02"))
}
