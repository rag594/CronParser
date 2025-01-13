package parser

import (
	"reflect"
	"strconv"
	"testing"
)

func TestParseDayOfWeek(t *testing.T) {
	t.Run("WildCard of format *", func(t *testing.T) {
		cronDayOfWeekParser := &CronDayOfWeekParser{DayOfWeekExpression: "*"}
		actualDayOfWeeks := cronDayOfWeekParser.Parse()
		expectedDayOfWeeks := func() []string {
			var dayOfWeeks []string
			for i := 1; i <= 7; i++ {
				dayOfWeeks = append(dayOfWeeks, strconv.Itoa(i))
			}
			return dayOfWeeks
		}()

		if !reflect.DeepEqual(actualDayOfWeeks, expectedDayOfWeeks) {
			t.Errorf("actual %v, wanted %v", actualDayOfWeeks, expectedDayOfWeeks)
		}
	})

	t.Run("Range of format 2-5", func(t *testing.T) {
		cronDayOfWeekParser := &CronDayOfWeekParser{DayOfWeekExpression: "2-5"}
		actualDayOfWeeks := cronDayOfWeekParser.Parse()
		expectedDayOfWeeks := func() []string {
			var dayOfWeeks []string
			for i := 2; i <= 5; i++ {
				dayOfWeeks = append(dayOfWeeks, strconv.Itoa(i))
			}
			return dayOfWeeks
		}()

		if !reflect.DeepEqual(actualDayOfWeeks, expectedDayOfWeeks) {
			t.Errorf("actual %v, wanted %v", actualDayOfWeeks, expectedDayOfWeeks)
		}
	})

	t.Run("Step of format with */2", func(t *testing.T) {
		cronDayOfWeekParser := &CronDayOfWeekParser{DayOfWeekExpression: "*/2"}
		actualDayOfWeeks := cronDayOfWeekParser.Parse()
		expectedDayOfWeeks := func() []string {
			dayOfWeeks := []string{"1", "3", "5", "7"}

			return dayOfWeeks
		}()

		if !reflect.DeepEqual(actualDayOfWeeks, expectedDayOfWeeks) {
			t.Errorf("actual %v, wanted %v", actualDayOfWeeks, expectedDayOfWeeks)
		}
	})

	t.Run("step/range combination of format 1-5/2", func(t *testing.T) {
		cronDayOfWeekParser := &CronDayOfWeekParser{DayOfWeekExpression: "1-5/2"}
		actualDayOfWeeks := cronDayOfWeekParser.Parse()
		expectedDayOfWeeks := func() []string {
			dayOfWeeks := []string{"1", "3", "5"}

			return dayOfWeeks
		}()

		if !reflect.DeepEqual(actualDayOfWeeks, expectedDayOfWeeks) {
			t.Errorf("actual %v, wanted %v", actualDayOfWeeks, expectedDayOfWeeks)
		}
	})

	t.Run("step/range combination of format MON-WED", func(t *testing.T) {
		cronDayOfWeekParser := &CronDayOfWeekParser{DayOfWeekExpression: "MON-WED"}
		actualDayOfWeeks := cronDayOfWeekParser.Parse()
		expectedDayOfWeeks := func() []string {
			dayOfWeeks := []string{"MON", "TUE", "WED"}

			return dayOfWeeks
		}()

		if !reflect.DeepEqual(actualDayOfWeeks, expectedDayOfWeeks) {
			t.Errorf("actual %v, wanted %v", actualDayOfWeeks, expectedDayOfWeeks)
		}
	})

}
