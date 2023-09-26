/*

package : go.core/assignments

Go code compares date

*/

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("enter date 1 in 'YYYY-MM-DD'  format")
	inputDate1 := getDate()

	fmt.Println("enter date 2 in 'YYYY-MM-DD' format ")
	inputDate2 := getDate()

	fmt.Println(compareDates(inputDate1, inputDate2))
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

// Function to compare dates and return the comparison result
func compareDates(date1, date2 time.Time) string {
	if date1.After(date2) {
		return fmt.Sprintf("\n%s is after %s", date1, date2)
	} else if date1.Before(date2) {
		return fmt.Sprintf("\n%s is before %s", date1, date2)
	} else {
		return fmt.Sprintf("\n%s is equal to %s", date1, date2)
	}
}
