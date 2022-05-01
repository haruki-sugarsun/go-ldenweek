package main

import (
	"flag"
	"fmt"
	"time"
	"os"
	"log"

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

func formatDay(day time.Time) string {
	return day.Format("2006/1/2")
}

func main() {
	// Flag interpretation and Setup.
    target_year_flag := flag.Int("year", -1, "Target year of the Christian Era.")
    allowed_gap_flag := flag.Int("gap", 1, "Maximum gap between holidays to make them continuos for Go-ldenweek.")
    verbose_flag := flag.Bool("verbose", false, "Verbose debugging output if given.")
	flag.Parse()

	// Constants
	// TODO: do we have better way to declare a constant?
	go_ldenweek_ingredients := []*cal.Holiday{
		jp.ShowaDay,
		jp.ConstitutionMemorialDay,
		jp.GreeneryDay,
		jp.ChildrensDay,
	}
	var target_year = *target_year_flag
	if target_year <= 0 { target_year = time.Now().Year() }
	allowed_gap := *allowed_gap_flag
	if allowed_gap < 0 {
		fmt.Println("allowed_gap cannot be negative.")
        os.Exit(1)
	}

	// Core calculation.
	fmt.Printf("Calculating the Go-ldenweek of %d.\n", target_year)
	holiday_instances := funk.Map(go_ldenweek_ingredients[:],
		func(x *cal.Holiday) time.Time {
			actual, observed := x.Calc(target_year)
			if *verbose_flag {
				log.Printf("     ----- %#v \n", x)
				log.Printf("actual: %s\n", actual)
				log.Printf("observed: %s\n", observed)
				log.Printf("observed.Weekday(): %s\n", observed.Weekday())
			}
			return observed
		}).([]time.Time)

	// Determine the "start" of Go-ldenweek.
	// Have the obserbed ShowaDay as a temporary start, and include the nearest weekend just before.
	start := (func() time.Time {
		_, temp_start := jp.ShowaDay.Calc(target_year) // TODO: Reconsider var/const naming.
		sunday, diff := sundayBefore(temp_start)
		if diff <= allowed_gap+1 {
			// Diff 1 means the days are continuos, so days allowed_gap+1 away can be connected.
			// this weekend can be a part of Go-ldenweek.
			return sunday.AddDate(0, 0, -1)
		} else {
			return temp_start
		}
	})()
	if *verbose_flag {
		log.Printf("start is %s.", start)
	}

	// Sweep forward to check "connected" holidays and weekends.
	// TODO: Implement more intuitive approach? e.g. merging the vacation periods.
	end := (func() time.Time {
		var cursor = start // day to test if it is a vacation-ish.

		// TODO: this only output the possible Go-ldenweek including Showa day.
		// Fix to detect the longest possible. e.g. for 2020.
	sweepings:
		for cursor.Year() == target_year {
			for i := 1; i <= allowed_gap+1; i++ {
				day_to_check := cursor.AddDate(0, 0, i)
				if *verbose_flag { log.Printf("   --- checking %s\n", day_to_check) }
				if isHolidayOrWeekend(day_to_check, holiday_instances) {
					if *verbose_flag { log.Printf("   --- vacation-ish! %s\n", day_to_check) }

					cursor = day_to_check
					continue sweepings
				}
			}
			// did not reach to the next vacation-ish day.
			break
		} // sweepings
		return cursor
	})()

	fmt.Printf("Go-ldenweek is %s ~ %s\n", formatDay(start), formatDay(end))
}
