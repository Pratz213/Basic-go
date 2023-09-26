package main

import (
	"testing"
	"time"
)

func TestCalculateBusinessDays(t *testing.T) {
	holidays := []time.Time{
		parseDate("2023-08-22"),
		parseDate("2023-08-26"),
		parseDate("2023-08-18"),
	}

	testCases := []struct {
		startDate    string
		businessDays int
		expectedDate string
	}{
		{"2023-08-23", 1, "2023-08-24"},
		{"2023-08-25", 0, "2023-08-25"},
		{"2023-08-22", -1, "2023-08-21"},
		{"2023-08-24", -2, "2023-08-21"},
		{"2023-08-17", 5, "2023-08-28"}, // Considering holidays and weekends
	}

	for _, tc := range testCases {
		t.Run(tc.startDate, func(t *testing.T) {
			startDate := parseDate(tc.startDate)
			expectedDate := parseDate(tc.expectedDate)

			actualDate := calculateBusinessDays(startDate, tc.businessDays, holidays)

			if !actualDate.Equal(expectedDate) {
				t.Errorf("For start date %s and business days %d, expected date %s, but got %s",
					tc.startDate, tc.businessDays, expectedDate.Format("2006-01-02"), actualDate.Format("2006-01-02"))
			}
		})
	}
}

func parseDate(dateStr string) time.Time {
	date, _ := time.Parse("2006-01-02", dateStr)
	return date
}
