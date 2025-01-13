package parser

import (
	"cronParser/models"
	"fmt"
	"strconv"
	"strings"
)

type CronDayOfMonthParser struct {
	DayOfMonthExpression string
}

func (c *CronDayOfMonthParser) ListParser(values []int) []int {
	var days []int

	if !strings.Contains(c.DayOfMonthExpression, models.List.Value()) {
		return values
	}

	splitSlice := strings.Split(c.DayOfMonthExpression, models.List.Value())
	for _, numStr := range splitSlice {
		numInt, _ := strconv.Atoi(numStr)
		days = append(days, numInt)
	}

	return days
}

func (c *CronDayOfMonthParser) WildCardParser() []int {
	var days []int

	for i := 1; i <= 31; i++ {
		days = append(days, i)
	}

	return days
}

func (c *CronDayOfMonthParser) NumberParser(values []int) []int {
	var days []int

	if strings.Contains(c.DayOfMonthExpression, models.Range.Value()) ||
		strings.Contains(c.DayOfMonthExpression, models.WildCard.Value()) ||
		strings.Contains(c.DayOfMonthExpression, models.List.Value()) {
		return values
	}

	parts := strings.FieldsFunc(c.DayOfMonthExpression, func(r rune) bool {
		return r == '/'
	})

	number, _ := strconv.Atoi(parts[0])

	if strings.Contains(c.DayOfMonthExpression, models.Step.Value()) {
		for i := 0; i < len(values); i++ {
			if values[i] >= number {
				days = append(days, values[i])
			}
		}
		return days
	}

	return []int{number}
}

func (c *CronDayOfMonthParser) StepParser(values []int) []int {
	var days []int

	if !strings.Contains(c.DayOfMonthExpression, models.Step.Value()) {
		return values
	}
	parts := strings.FieldsFunc(c.DayOfMonthExpression, func(r rune) bool {
		return r == '/'
	})
	step, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(values); i++ {
		if values[i]%step == values[0]%step {
			days = append(days, values[i])
		}
	}

	return days
}

func (c *CronDayOfMonthParser) RangeParser(values []int) []int {
	var days []int
	if !strings.Contains(c.DayOfMonthExpression, models.Range.Value()) {
		return values
	}
	parts := strings.FieldsFunc(c.DayOfMonthExpression, func(r rune) bool {
		return r == '-' || r == '/'
	})
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	for i := 0; i < len(values); i++ {
		if values[i] >= start && values[i] <= end {
			days = append(days, values[i])
		}
	}
	return days
}

func (c *CronDayOfMonthParser) Parse() []int {

	x := c.WildCardParser()
	j := c.ListParser(x)
	f := c.NumberParser(j)
	y := c.RangeParser(f)
	z := c.StepParser(y)

	return z
}
