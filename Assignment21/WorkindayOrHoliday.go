/*
package : go.core/assignment

*/

package main

import (
	"fmt"
	"os"
	"time"
)

const (
	DATE_FORMAT = "2006-01-02"
)

func main() {

	// Accept holiday list from user and stores in slice
	holidayList := getHolidayList()

	//accept date to check holiday or not
	fmt.Println("Enter date(YYYY-MM-DD) to check holiday or working day:")
	var dateToCheck string
	fmt.Scanln(&dateToCheck)

	//Parsing string to date
	parsedDate, err := time.Parse(DATE_FORMAT, dateToCheck)
	checkErrIsNil("Invalid Date! ", err)

	if isHoliday(parsedDate, holidayList) {
		fmt.Println("It's a holiday!")
	} else {
		fmt.Println("It's a working day!")
	}
}

// gets list of holidays from user
func getHolidayList() []time.Time {
	fmt.Println("Enter a holiday date (YYYY-MM-DD) or word 'proceed'")
	var holidays []time.Time
	for {
		var inputDate string
		fmt.Scanln(&inputDate)

		if inputDate == "proceed" {
			return holidays
		}

		parsedDate, err := time.Parse("2006-01-02", inputDate)
		checkErrIsNil("Invalid Date! ", err)
		holidays = append(holidays, parsedDate)
	}
}

// isHoliday checks whether a given date is a holiday, Sunday, or Saturday.
func isHoliday(date time.Time, holidayList []time.Time) bool {
	isHoliday := false

	// Check if the date is a Saturday or Sunday
	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		isHoliday = true
	} else {
		// Check if the date is in the list of holidays
		for _, holidayDay := range holidayList {
			if holidayDay.Equal(date) {
				isHoliday = true
				break //breaks loop as date is present in  holiday list
			}
		}
	}
	return isHoliday
}

// Prints Error!
func checkErrIsNil(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}
