package util

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	now := time.Date(2023, 10, 25, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "Birthday has passed this year",
			dob:      time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
			expected: 33,
		},
		{
			name:     "Birthday is today",
			dob:      time.Date(1990, 10, 25, 0, 0, 0, 0, time.UTC),
			expected: 33,
		},
		{
			name:     "Birthday is tomorrow",
			dob:      time.Date(1990, 10, 26, 0, 0, 0, 0, time.UTC),
			expected: 32,
		},
		{
			name:     "Birthday next month",
			dob:      time.Date(1990, 11, 1, 0, 0, 0, 0, time.UTC),
			expected: 32,
		},
		{
			name:     "Born today",
			dob:      time.Date(2023, 10, 25, 0, 0, 0, 0, time.UTC),
			expected: 0,
		},
		{
			name:     "Leap year birthday (Feb 29) on non-leap year",
			dob:      time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
			expected: 23,
			// Explanation: Feb 29 treated as Mar 1 in non-leap years.
			// In Oct 2023, the birthday has passed.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateAge(tt.dob, now)
			if got != tt.expected {
				t.Errorf("CalculateAge() = %v, want %v", got, tt.expected)
			}
		})
	}
}
