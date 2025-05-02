package goldenweek

import (
	"time"

	"github.com/rickar/cal/v2"
)

func SundayBefore(day time.Time) (time.Time, int) {
	weekday := day.Weekday()
	if weekday == time.Sunday {
		return day, 0
	}
	
	diff := int(weekday)
	return day.AddDate(0, 0, -diff), diff
}

func IsHolidayOrWeekend(day time.Time, holidayInstances []time.Time) bool {
	if cal.IsWeekend(day) {
		return true
	}

	for _, holiday := range holidayInstances {
		if day.Equal(holiday) {
			return true
		}
	}

	return false
}

func FormatDay(day time.Time) string {
	return day.Format("2006/1/2")
}
