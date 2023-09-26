package main

import (
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {

	//slice of holiday list
	holidays := []time.Time{
		time.Date(2023, time.August, 01, 0, 0, 0, 0, time.Local),
		time.Date(2023, time.August, 30, 0, 0, 0, 0, time.Local),
		time.Date(2023, time.August, 15, 0, 0, 0, 0, time.Local),
	}

	//structure of input and expected output
	tests := []struct {
		inputDate   time.Time
		expectedMsg string
	}{
		{time.Date(2023, time.August, 01, 0, 0, 0, 0, time.Local), "It's a holiday!"},
		{time.Date(2023, time.August, 15, 0, 0, 0, 0, time.Local), "It's a holiday!"},
		{time.Date(2023, time.August, 30, 0, 0, 0, 0, time.Local), "It's a holiday!"},
		{time.Date(2023, time.August, 17, 0, 0, 0, 0, time.Local), "It's a working day!"},
		{time.Date(2023, time.August, 28, 0, 0, 0, 0, time.Local), "It's a working day!"},
	}

	//Loop over Testcases mention in structure
	for _, test := range tests {
		t.Run(test.inputDate.Format("2006-01-02"), func(t *testing.T) {
			date := test.inputDate
			isHoliday := isHoliday(date, holidays)
			var expectedOutput string
			if isHoliday {
				expectedOutput = "It's a holiday!"
			} else {
				expectedOutput = "It's a working day!"
			}
			if expectedOutput != test.expectedMsg {
				t.Errorf("For date %s, expected '%s', but got '%s'", test.inputDate, test.expectedMsg, expectedOutput)
			}
		})
	}
}
