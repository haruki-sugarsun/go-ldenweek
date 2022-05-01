package main

import (
	"fmt"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/jp"
	"github.com/thoas/go-funk"
)

// Go-ldenweek
// https://ja.wikipedia.org/wiki/%E3%82%B4%E3%83%BC%E3%83%AB%E3%83%87%E3%83%B3%E3%82%A6%E3%82%A3%E3%83%BC%E3%82%AF
// TODO: Support silverweek too?

// Utilities
// Returns the nearest Sunday from the day given, and the difference in day.
func sundayBefore(day time.Time) (time.Time, int) {
	switch day.Weekday() { // TODO: Can be improved.
	case time.Sunday:
		return day, 0
	case time.Monday:
		return day.AddDate(0, 0, -1), 1
	case time.Tuesday:
		return day.AddDate(0, 0, -2), 2
	case time.Wednesday:
		return day.AddDate(0, 0, -3), 3
	case time.Thursday:
		return day.AddDate(0, 0, -4), 4
	case time.Friday:
		return day.AddDate(0, 0, -5), 5
	case time.Saturday:
		return day.AddDate(0, 0, -6), 6
	}

	// Should not fall here.
	panic(fmt.Sprintf("%s has no Weekday.", day))
}

func previousDay(day time.Time) time.Time {
	return day.AddDate(0, 0, -1)
}

func nextDay(day time.Time) time.Time {
	return day.AddDate(0, 0, -1)
}

func isHolidayOrWeekend(day time.Time, holiday_instances []time.Time) bool {
	if cal.IsWeekend(day) {
		return true
	}

	for _, ing := range holiday_instances {
		if day.Equal(ing) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("Go-ldenweek!")

	// Constants
	// TODO: do we have better way to declare a constant?
	go_ldenweek_ingredients := []*cal.Holiday{
		jp.ShowaDay,
		jp.ConstitutionMemorialDay,
		jp.GreeneryDay,
		jp.ChildrensDay,
	}
	current_year := 2022
	allowed_gap := 1

	integers := funk.Map([]int{1, 2, 3}, func(i int) int { return i + 1 })
	fmt.Printf("integers = %s\n", integers)

	tmp_holiday_instances := funk.Map(go_ldenweek_ingredients[:], func(x *cal.Holiday) time.Time { _, o := x.Calc(current_year); return o }).([]time.Time)
	holiday_instances := tmp_holiday_instances

	for _, ex := range go_ldenweek_ingredients {
		var actual, observed = ex.Calc(current_year)

		fmt.Printf("     ----- %s \n", ex)
		fmt.Printf("actual: %s\n", actual)
		fmt.Printf("observed: %s\n", observed)

		fmt.Printf("observed.Weekday(): %s\n", observed.Weekday())

	}

	// Determine the "start" of Go-ldenweek.
	// Have the obserbed ShowaDay as a temporary start, and include the nearest weekend just before.
	start := (func() time.Time {
		_, temp_start := jp.ShowaDay.Calc(current_year) // TODO: Reconsider var/const naming.
		sunday, diff := sundayBefore(temp_start)
		if diff <= allowed_gap+1 {
			// Diff 1 means the days are continuos, so days allowed_gap+1 away can be connected.
			// this weekend can be a part of Go-ldenweek.
			return previousDay(sunday)
		} else {
			return temp_start
		}
	})()
	fmt.Printf("start is %s.", start)

	// Sweep forward to check "connected" holidays and weekends.
	// TODO: Implement more intuitive approach? e.g. merging the vacation periods.
	end := (func() time.Time {
		var cursor = start // day to test if it is a vacation-ish.

	sweepings:
		for cursor.Year() == current_year {
			for i := 1; i <= allowed_gap+1; i++ {
				day_to_check := cursor.AddDate(0, 0, i)
				fmt.Printf("   --- checking %s\n", day_to_check)
				if isHolidayOrWeekend(day_to_check, holiday_instances) {
					fmt.Printf("   --- vacation-ish! %s\n", day_to_check)

					cursor = day_to_check
					continue sweepings
				}
			}
			// did not reach to the next vacation-ish day.
			break
		} // sweepings
		return cursor
	})()

	fmt.Printf("Go-ldenweek is %s ~ %s", start, end)
}
