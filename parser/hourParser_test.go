package parser

import (
	"reflect"
	"testing"
)

func TestParseHours(t *testing.T) {
	t.Run("WildCard of format *", func(t *testing.T) {
		cronHourParser := &CronHourParser{HourExpression: "*"}
		actualHours := cronHourParser.Parse()
		expectedHours := func() []int {
			var hours []int
			for i := 0; i <= 23; i++ {
				hours = append(hours, i)
			}
			return hours
		}()

		if !reflect.DeepEqual(actualHours, expectedHours) {
			t.Errorf("actual %v, wanted %v", actualHours, expectedHours)
		}
	})

	t.Run("Range of format 3-14", func(t *testing.T) {
		cronHourParser := &CronHourParser{HourExpression: "3-14"}
		actualHours := cronHourParser.Parse()
		expectedHours := func() []int {
			var hours []int
			for i := 3; i <= 14; i++ {
				hours = append(hours, i)
			}
			return hours
		}()

		if !reflect.DeepEqual(actualHours, expectedHours) {
			t.Errorf("actual %v, wanted %v", actualHours, expectedHours)
		}
	})

	t.Run("Step of format with */10", func(t *testing.T) {
		cronHourParser := &CronHourParser{HourExpression: "*/10"}
		actualHours := cronHourParser.Parse()
		expectedHours := func() []int {
			hours := []int{0, 10, 20}

			return hours
		}()

		if !reflect.DeepEqual(actualHours, expectedHours) {
			t.Errorf("actual %v, wanted %v", actualHours, expectedHours)
		}
	})

	t.Run("step/range combination of format 1-12/3", func(t *testing.T) {
		cronHourParser := &CronHourParser{HourExpression: "1-12/3"}
		actualHours := cronHourParser.Parse()
		expectedHours := func() []int {
			hours := []int{1, 4, 7, 10}

			return hours
		}()

		if !reflect.DeepEqual(actualHours, expectedHours) {
			t.Errorf("actual %v, wanted %v", actualHours, expectedHours)
		}
	})

}
