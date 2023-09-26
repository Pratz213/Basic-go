/*

package : go.core/assignments

Go code to Print Date and time in various formats

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now().UTC()

	// Format 1: "16 Mar 2022"
	fmt.Println(currentTime.Format("02 Jan 2006"))

	// Format 2: "Mar 16, 2022"
	fmt.Println(currentTime.Format("Jan 02, 2006"))

	// Format 3: "2022-03-16"
	fmt.Println(currentTime.Format("2006-01-02"))

	// Format 4: "2022-03-16T15:52:00Z"
	fmt.Println(currentTime.Format("2006-01-02T15:52:00Z"))

	// Format 5: "Tuesday, 16 March 2022"
	fmt.Println(currentTime.Format("Monday, 02 January 2006"))

	// Format 6: "2023-07-28"
	fmt.Println(time.DateOnly)

	// Format 7: "9:45AM"
	fmt.Println(time.Kitchen)

	// Format 8: 9 August 2023 10:29:46
	fmt.Println(currentTime.Day(), currentTime.Month(), currentTime.Year(), currentTime.Format("15:04:05"))
}
