package goldenweek

import (
	"testing"
	"time"
)

func TestSundayBefore(t *testing.T) {
	tests := []struct {
		name           string
		date           time.Time
		expectedSunday time.Time
		expectedDiff   int
	}{
		{
			name:           "Sunday returns itself",
			date:           time.Date(2022, 5, 1, 0, 0, 0, 0, time.UTC), // Sunday
			expectedSunday: time.Date(2022, 5, 1, 0, 0, 0, 0, time.UTC),
			expectedDiff:   0,
		},
		{
			name:           "Monday returns previous Sunday",
			date:           time.Date(2022, 5, 2, 0, 0, 0, 0, time.UTC), // Monday
			expectedSunday: time.Date(2022, 5, 1, 0, 0, 0, 0, time.UTC),
			expectedDiff:   1,
		},
		{
			name:           "Saturday returns previous Sunday",
			date:           time.Date(2022, 5, 7, 0, 0, 0, 0, time.UTC), // Saturday
			expectedSunday: time.Date(2022, 5, 1, 0, 0, 0, 0, time.UTC),
			expectedDiff:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSunday, gotDiff := SundayBefore(tt.date)
			if !gotSunday.Equal(tt.expectedSunday) {
				t.Errorf("SundayBefore() gotSunday = %v, want %v", gotSunday, tt.expectedSunday)
			}
			if gotDiff != tt.expectedDiff {
				t.Errorf("SundayBefore() gotDiff = %v, want %v", gotDiff, tt.expectedDiff)
			}
		})
	}
}

func TestIsHolidayOrWeekend(t *testing.T) {
	tests := []struct {
		name       string
		date       time.Time
		holidays   []time.Time
		expected   bool
	}{
		{
			name:       "Weekend (Sunday) is true",
			date:       time.Date(2022, 5, 1, 0, 0, 0, 0, time.UTC), // Sunday
			holidays:   []time.Time{},
			expected:   true,
		},
		{
			name:       "Weekend (Saturday) is true",
			date:       time.Date(2022, 5, 7, 0, 0, 0, 0, time.UTC), // Saturday
			holidays:   []time.Time{},
			expected:   true,
		},
		{
			name:       "Weekday without holiday is false",
			date:       time.Date(2022, 5, 2, 0, 0, 0, 0, time.UTC), // Monday
			holidays:   []time.Time{},
			expected:   false,
		},
		{
			name:       "Holiday is true",
			date:       time.Date(2022, 5, 3, 0, 0, 0, 0, time.UTC), // Tuesday
			holidays:   []time.Time{time.Date(2022, 5, 3, 0, 0, 0, 0, time.UTC)},
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsHolidayOrWeekend(tt.date, tt.holidays)
			if got != tt.expected {
				t.Errorf("IsHolidayOrWeekend() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFormatDay(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected string
	}{
		{
			name:     "Format day correctly",
			date:     time.Date(2022, 5, 1, 0, 0, 0, 0, time.UTC),
			expected: "2022/5/1",
		},
		{
			name:     "Format day with single-digit month and day",
			date:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: "2022/1/1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatDay(tt.date)
			if got != tt.expected {
				t.Errorf("FormatDay() = %v, want %v", got, tt.expected)
			}
		})
	}
}
