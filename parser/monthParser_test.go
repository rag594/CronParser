package parser

import (
	"reflect"
	"testing"
)

func TestParseMonth(t *testing.T) {
	t.Run("WildCard of format *", func(t *testing.T) {
		cronMonthParser := &CronMonthParser{MonthExpression: "*"}
		actualMonths := cronMonthParser.Parse()
		expectedMonths := func() []int {
			var Months []int
			for i := 1; i <= 12; i++ {
				Months = append(Months, i)
			}
			return Months
		}()

		if !reflect.DeepEqual(actualMonths, expectedMonths) {
			t.Errorf("actual %v, wanted %v", actualMonths, expectedMonths)
		}
	})

	t.Run("Range of format 4-10", func(t *testing.T) {
		cronMonthParser := &CronMonthParser{MonthExpression: "4-10"}
		actualMonths := cronMonthParser.Parse()
		expectedMonths := func() []int {
			var Months []int
			for i := 4; i <= 10; i++ {
				Months = append(Months, i)
			}
			return Months
		}()

		if !reflect.DeepEqual(actualMonths, expectedMonths) {
			t.Errorf("actual %v, wanted %v", actualMonths, expectedMonths)
		}
	})

	t.Run("Step of format with */3", func(t *testing.T) {
		cronMonthParser := &CronMonthParser{MonthExpression: "*/3"}
		actualMonths := cronMonthParser.Parse()
		expectedMonths := func() []int {
			Months := []int{1, 4, 7, 10}

			return Months
		}()

		if !reflect.DeepEqual(actualMonths, expectedMonths) {
			t.Errorf("actual %v, wanted %v", actualMonths, expectedMonths)
		}
	})

	t.Run("step/range combination of format 4-12/3", func(t *testing.T) {
		cronMonthParser := &CronMonthParser{MonthExpression: "4-12/3"}
		actualMonths := cronMonthParser.Parse()
		expectedMonths := func() []int {
			Months := []int{4, 7, 10}

			return Months
		}()

		if !reflect.DeepEqual(actualMonths, expectedMonths) {
			t.Errorf("actual %v, wanted %v", actualMonths, expectedMonths)
		}
	})

}
