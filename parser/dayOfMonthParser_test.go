package parser

import (
	"reflect"
	"testing"
)

func TestParseDayOfMonth(t *testing.T) {
	t.Run("WildCard of format *", func(t *testing.T) {
		cronDayOfMonthParser := &CronDayOfMonthParser{DayOfMonthExpression: "*"}
		actualDayOfMonths := cronDayOfMonthParser.Parse()
		expectedDayOfMonths := func() []int {
			var dayOfMonths []int
			for i := 1; i <= 31; i++ {
				dayOfMonths = append(dayOfMonths, i)
			}
			return dayOfMonths
		}()

		if !reflect.DeepEqual(actualDayOfMonths, expectedDayOfMonths) {
			t.Errorf("actual %v, wanted %v", actualDayOfMonths, expectedDayOfMonths)
		}
	})

	t.Run("Range of format 2-15", func(t *testing.T) {
		cronDayOfMonthParser := &CronDayOfMonthParser{DayOfMonthExpression: "2-15"}
		actualDayOfMonths := cronDayOfMonthParser.Parse()
		expectedDayOfMonths := func() []int {
			var dayOfMonths []int
			for i := 2; i <= 15; i++ {
				dayOfMonths = append(dayOfMonths, i)
			}
			return dayOfMonths
		}()

		if !reflect.DeepEqual(actualDayOfMonths, expectedDayOfMonths) {
			t.Errorf("actual %v, wanted %v", actualDayOfMonths, expectedDayOfMonths)
		}
	})

	t.Run("Step of format with */10", func(t *testing.T) {
		cronDayOfMonthParser := &CronDayOfMonthParser{DayOfMonthExpression: "*/10"}
		actualDayOfMonths := cronDayOfMonthParser.Parse()
		expectedDayOfMonths := func() []int {
			dayOfMonths := []int{1, 11, 21, 31}

			return dayOfMonths
		}()

		if !reflect.DeepEqual(actualDayOfMonths, expectedDayOfMonths) {
			t.Errorf("actual %v, wanted %v", actualDayOfMonths, expectedDayOfMonths)
		}
	})

	t.Run("step/range combination of format 4-12/3", func(t *testing.T) {
		cronDayOfMonthParser := &CronDayOfMonthParser{DayOfMonthExpression: "4-12/3"}
		actualDayOfMonths := cronDayOfMonthParser.Parse()
		expectedDayOfMonths := func() []int {
			dayOfMonths := []int{4, 7, 10}

			return dayOfMonths
		}()

		if !reflect.DeepEqual(actualDayOfMonths, expectedDayOfMonths) {
			t.Errorf("actual %v, wanted %v", actualDayOfMonths, expectedDayOfMonths)
		}
	})

	t.Run("list of format 4,23,25", func(t *testing.T) {
		cronDayOfMonthParser := &CronDayOfMonthParser{DayOfMonthExpression: "4,23,25"}
		actualDayOfMonths := cronDayOfMonthParser.Parse()
		expectedDayOfMonths := func() []int {
			dayOfMonths := []int{4, 23, 25}

			return dayOfMonths
		}()

		if !reflect.DeepEqual(actualDayOfMonths, expectedDayOfMonths) {
			t.Errorf("actual %v, wanted %v", actualDayOfMonths, expectedDayOfMonths)
		}
	})

}
