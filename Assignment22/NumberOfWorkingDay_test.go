package main

import (
	"testing"
	"time"
)

func TestCalculateWorkingDays(t *testing.T) {

	//slice of holiday list
	holidays := []time.Time{
		parseDate("2023-08-06"),
		parseDate("2023-08-10"),
		parseDate("2023-08-15"),
		parseDate("2023-08-18"),
		parseDate("2023-08-20"),
	}

	//structure of input and expected output
	tests := []struct {
		startDate      string
		endDate        string
		expectedResult int
	}{
		{"2023-08-01", "2023-08-31", 20},
		{"2023-08-17", "2023-08-17", 1},
		{"2023-08-13", "2023-08-22", 5},
		{"2023-08-01", "2023-08-25", 16},
		{"2023-08-10", "2023-08-25", 9},
	}

	//Loop over Testcases mention in structure
	for _, test := range tests {
		t.Run(test.startDate+" "+test.endDate, func(t *testing.T) {
			startDate := parseDate(test.startDate)
			endDate := parseDate(test.endDate)

			result := calculateWorkingDays(startDate, endDate, holidays)

			if result != test.expectedResult {
				t.Errorf("For date range %s to %s, expected %d working days but got %d", test.startDate, test.endDate, test.expectedResult, result)
			}
		})
	}
}

func parseDate(dateStr string) time.Time {
	date, _ := time.Parse("2006-01-02", dateStr)
	return date
}
