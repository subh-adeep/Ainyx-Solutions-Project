package util

import (
	"time"
)

// CalculateAge calculates age given a date of birth and current time.
// It accounts for month and day to ensure accuracy.
func CalculateAge(dob time.Time, now time.Time) int {
	age := now.Year() - dob.Year()
	
	// Subtract 1 if the birthday hasn't occurred yet this year
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}
	
	return age
}
