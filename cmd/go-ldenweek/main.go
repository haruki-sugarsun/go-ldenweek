package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/haruki-sugarsun/go-ldenweek/internal/goldenweek"
	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/jp"
	"github.com/thoas/go-funk"
)

func main() {
	targetYearFlag := flag.Int("year", -1, "Target year of the Christian Era.")
	allowedGapFlag := flag.Int("gap", 1, "Maximum gap between holidays to make them continuos for Go-ldenweek.")
	verboseFlag := flag.Bool("verbose", false, "Verbose debugging output if given.")
	flag.Parse()

	goldenWeekIngredients := []*cal.Holiday{
		jp.ShowaDay,
		jp.ConstitutionMemorialDay,
		jp.GreeneryDay,
		jp.ChildrensDay,
	}
	targetYear := *targetYearFlag
	if targetYear <= 0 {
		targetYear = time.Now().Year()
	}
	allowedGap := *allowedGapFlag
	if allowedGap < 0 {
		fmt.Println("allowed_gap cannot be negative.")
		os.Exit(1)
	}

	fmt.Printf("Calculating the Go-ldenweek of %d.\n", targetYear)
	holidayInstances := funk.Map(goldenWeekIngredients[:],
		func(x *cal.Holiday) time.Time {
			actual, observed := x.Calc(targetYear)
			if *verboseFlag {
				log.Printf("     ----- %#v \n", x)
				log.Printf("actual: %s\n", actual)
				log.Printf("observed: %s\n", observed)
				log.Printf("observed.Weekday(): %s\n", observed.Weekday())
			}
			return observed
		}).([]time.Time)

	start := (func() time.Time {
		_, tempStart := jp.ShowaDay.Calc(targetYear)
		sunday, diff := goldenweek.SundayBefore(tempStart)
		if diff <= allowedGap+1 {
			return sunday.AddDate(0, 0, -1)
		} else {
			return tempStart
		}
	})()
	if *verboseFlag {
		log.Printf("start is %s.", start)
	}

	end := (func() time.Time {
		var cursor = start // day to test if it is a vacation-ish.

	sweepings:
		for cursor.Year() == targetYear {
			for i := 1; i <= allowedGap+1; i++ {
				dayToCheck := cursor.AddDate(0, 0, i)
				if *verboseFlag {
					log.Printf("   --- checking %s\n", dayToCheck)
				}
				if goldenweek.IsHolidayOrWeekend(dayToCheck, holidayInstances) {
					if *verboseFlag {
						log.Printf("   --- vacation-ish! %s\n", dayToCheck)
					}

					cursor = dayToCheck
					continue sweepings
				}
			}
			break
		} // sweepings
		return cursor
	})()

	fmt.Printf("Go-ldenweek is %s ~ %s\n", goldenweek.FormatDay(start), goldenweek.FormatDay(end))
}
