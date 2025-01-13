package parser

import (
	"cronParser/models"
	"fmt"
	"strconv"
	"strings"
)

type CronHourParser struct {
	HourExpression string
}

func (c *CronHourParser) WildCardParser() []int {
	var hours []int

	for i := 0; i <= 23; i++ {
		hours = append(hours, i)
	}

	return hours
}

func (c *CronHourParser) NumberParser(values []int) []int {
	var hours []int

	if strings.Contains(c.HourExpression, models.Range.Value()) ||
		strings.Contains(c.HourExpression, models.WildCard.Value()) {
		return values
	}

	parts := strings.FieldsFunc(c.HourExpression, func(r rune) bool {
		return r == '/'
	})

	number, _ := strconv.Atoi(parts[0])

	if strings.Contains(c.HourExpression, models.Step.Value()) {
		for i := 0; i < len(values); i++ {
			if values[i] >= number {
				hours = append(hours, values[i])
			}
		}
		return hours
	}

	return []int{number}
}

func (c *CronHourParser) StepParser(values []int) []int {
	var hours []int

	if !strings.Contains(c.HourExpression, models.Step.Value()) {
		return values
	}
	parts := strings.FieldsFunc(c.HourExpression, func(r rune) bool {
		return r == '/'
	})
	step, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(values); i++ {
		if values[i]%step == values[0]%step {
			hours = append(hours, values[i])
		}
	}

	return hours
}

func (c *CronHourParser) RangeParser(values []int) []int {
	var hours []int
	if !strings.Contains(c.HourExpression, models.Range.Value()) {
		return values
	}
	parts := strings.FieldsFunc(c.HourExpression, func(r rune) bool {
		return r == '-' || r == '/'
	})
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	for i := 0; i < len(values); i++ {
		if values[i] >= start && values[i] <= end {
			hours = append(hours, values[i])
		}
	}
	return hours
}

func (c *CronHourParser) Parse() []int {

	x := c.WildCardParser()
	f := c.NumberParser(x)
	y := c.RangeParser(f)
	z := c.StepParser(y)

	return z
}
