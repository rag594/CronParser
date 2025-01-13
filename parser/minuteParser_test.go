package parser

import (
	"reflect"
	"testing"
)

func TestParseMinutes(t *testing.T) {
	t.Run("WildCard of format *", func(t *testing.T) {
		cronMinuteParser := &CronMinuteParser{MinuteExpression: "*"}
		actualMinutes := cronMinuteParser.Parse()
		expectedMinutes := func() []int {
			var minutes []int
			for i := 0; i <= 59; i++ {
				minutes = append(minutes, i)
			}
			return minutes
		}()

		if !reflect.DeepEqual(actualMinutes, expectedMinutes) {
			t.Errorf("actual %v, wanted %v", actualMinutes, expectedMinutes)
		}
	})

	t.Run("Range of format 3-45", func(t *testing.T) {
		cronMinuteParser := &CronMinuteParser{MinuteExpression: "3-45"}
		actualMinutes := cronMinuteParser.Parse()
		expectedMinutes := func() []int {
			var minutes []int
			for i := 3; i <= 45; i++ {
				minutes = append(minutes, i)
			}
			return minutes
		}()

		if !reflect.DeepEqual(actualMinutes, expectedMinutes) {
			t.Errorf("actual %v, wanted %v", actualMinutes, expectedMinutes)
		}
	})

	t.Run("Step of format with */10", func(t *testing.T) {
		cronMinuteParser := &CronMinuteParser{MinuteExpression: "*/10"}
		actualMinutes := cronMinuteParser.Parse()
		expectedMinutes := func() []int {
			minutes := []int{0, 10, 20, 30, 40, 50}

			return minutes
		}()

		if !reflect.DeepEqual(actualMinutes, expectedMinutes) {
			t.Errorf("actual %v, wanted %v", actualMinutes, expectedMinutes)
		}
	})

	t.Run("step/range combination of format 5-45/10", func(t *testing.T) {
		cronMinuteParser := &CronMinuteParser{MinuteExpression: "5-45/10"}
		actualMinutes := cronMinuteParser.Parse()
		expectedMinutes := func() []int {
			minutes := []int{5, 15, 25, 35, 45}

			return minutes
		}()

		if !reflect.DeepEqual(actualMinutes, expectedMinutes) {
			t.Errorf("actual %v, wanted %v", actualMinutes, expectedMinutes)
		}
	})

}
