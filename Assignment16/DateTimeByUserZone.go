/*

package : go.core/assignments

Go code to This Go program prompts the user to input a time zone and a date format choice.
It then displays the current date and time in the user specified time zone using the chosen date format.

*/

package main

import (
	"fmt"
	"time"
)

var zone string

var supportedTimezones = []string{
	"America/New_York",
	"America/Los_Angeles",
	"Europe/London",
	"Europe/Berlin",
	"Asia/Kolkata",
	"Australia/Sydney",
}

func main() {

	// Prompt user for a valid time zone
	fmt.Print("\nEnter a valid time zone from the below list: ")
	var userTimeZone string = getTimeZone()

	loc, err := time.LoadLocation(userTimeZone)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//stores the current time from the user-selected time zone
	currentTime := time.Now().In(loc)

	// gets User selected time format
	dateFormat := getDateFormat()

	fmt.Println("\nCurrent Date/Time in", userTimeZone, ":")
	fmt.Println(currentTime.Format(dateFormat))

}

func getTimeZone() string {
	// Print the list of supported time zones
	for _, zone = range supportedTimezones {
		fmt.Println(zone)
	}
	fmt.Println("\n--------------")
	var timeZone string
	fmt.Scanln(&timeZone)

	switch timeZone {
	case "America/New_York":
		return timeZone
	case "America/Los_Angeles":
		return timeZone
	case "Europe/London":
		return timeZone
	case "Europe/Berlin":
		return timeZone
	case "Asia/Kolkata":
		return timeZone
	case "Australia/Sydney":
		return timeZone
	default:
		return "Invalid time zone selected."
	}
}

func getDateFormat() string {
	var formatChoice int
	fmt.Println("\nSelect a format no:", "\n1. 16 Mar 2022", "\n2. Mar 16, 2022", "\n3. 2022-03-16", "\n4. 2022-03-16T15:52:00Z", "\n5. Tuesday, 16 March 2022")
	fmt.Scanln(&formatChoice)

	switch formatChoice {
	case 1:
		return "02 Jan 2006"
	case 2:
		return "Jan 02, 2006"
	case 3:
		return "2006-01-02"
	case 4:
		return "2006-01-02T15:04:05Z"
	case 5:
		return "Monday, 02 January 2006"
	default:
		return "Invalid format choice."
	}
}
