package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	// Accept holiday list from user and stores in slice
	holidayList := getHolidayList()

	// Accept input for start date
	fmt.Println("Enter the start date (YYYY-MM-DD):")
	startDate := getDate()

	// Accept input no. of input date
	fmt.Println("Enter the business days:")
	businessNum := getBusinessDaysNum()

	workingDate := calculateBusinessDays(startDate, businessNum, holidayList)

	fmt.Println("working date: ", workingDate.Format("2006-01-02"))
}

// calculateBusinessDays calculates the output date by considering business days while excluding holidays.
func calculateBusinessDays(inputDate time.Time, numBusinessDays int, holidays []time.Time) time.Time {
	itetrateDay := 1

	// If the number of business days is negative, itetrating in backward
	if numBusinessDays < 0 {
		itetrateDay = -1
		numBusinessDays = -numBusinessDays
	}

	// Loop for the business days
	for i := 0; i < numBusinessDays; i++ {
		inputDate = inputDate.AddDate(0, 0, itetrateDay)
		if isHoliday(inputDate, holidays) {
			i--
		}
	}
	return inputDate
}

// gets date from user
func getDate() time.Time {
	var inputDate string
	fmt.Scanln(&inputDate)
	dateFormat := "2006-01-02" //parsing format
	parsedDate, err := time.Parse(dateFormat, inputDate)
	checkErrIsNil("invalid input! ", err)
	return parsedDate
}

// checks for valid number or throws error
func getBusinessDaysNum() int {
	var num string
	fmt.Scanln(&num)
	number, err := strconv.Atoi(num) //converting string to int
	checkErrIsNil("invalid input! ", err)
	return number
}

// gets list of holidays from user
func getHolidayList() []time.Time {
	fmt.Println("Enter a holiday date (YYYY-MM-DD) or word 'proceed'")
	var holidays []time.Time
	for {
		var inputDate string
		fmt.Scanln(&inputDate)
		dateFormat := "2006-01-02" //parsing format

		if inputDate == "proceed" {
			return holidays
		}

		parsedDate, err := time.Parse(dateFormat, inputDate)
		checkErrIsNil("invalid input! ", err)
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
