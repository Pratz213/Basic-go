package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//accepting holiday list from user
	holidaysList := getHolidayList()

	//accepting start date and end date  from user
	fmt.Print("Enter the start date (YYYY-MM-DD): ")
	startDateInput := getDate()
	fmt.Print("Enter the end date (YYYY-MM-DD): ")
	endDateInput := getDate()

	//calculates No. of working days from startdate till endate.
	workingDays := calculateWorkingDays(startDateInput, endDateInput, holidaysList)
	fmt.Printf("Number of working days between %s and %s: %d\n", startDateInput.Format("02 Jan 2006"), endDateInput.Format("02 Jan 2006"), workingDays)
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
		if err != nil {
			fmt.Println("invalid input")
			os.Exit(1)
		}
		holidays = append(holidays, parsedDate)
	}
}

// calculates Working Days
func calculateWorkingDays(startDate, endDate time.Time, holidays []time.Time) int {
	days := 0

	//loop from start date till end date (includes starting date and ending date)
	for currentDate := startDate; !currentDate.After(endDate); currentDate = currentDate.AddDate(0, 0, 1) {
		if !isHoliday(currentDate, holidays) {
			days++
		}
	}
	return days
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
