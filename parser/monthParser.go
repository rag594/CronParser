package parser

import (
	"cronParser/models"
	"fmt"
	"strconv"
	"strings"
)

type CronMonthParser struct {
	MonthExpression string
}

func (c *CronMonthParser) WildCardParser() []int {
	var months []int

	for i := 1; i <= 12; i++ {
		months = append(months, i)
	}

	return months
}

func (c *CronMonthParser) NumberParser(values []int) []int {
	var months []int

	if strings.Contains(c.MonthExpression, models.Range.Value()) ||
		strings.Contains(c.MonthExpression, models.WildCard.Value()) {
		return values
	}

	parts := strings.FieldsFunc(c.MonthExpression, func(r rune) bool {
		return r == '/'
	})

	number, _ := strconv.Atoi(parts[0])

	if strings.Contains(c.MonthExpression, models.Step.Value()) {
		for i := 0; i < len(values); i++ {
			if values[i] >= number {
				months = append(months, values[i])
			}
		}
		return months
	}

	return []int{number}
}

func (c *CronMonthParser) StepParser(values []int) []int {
	var months []int

	if !strings.Contains(c.MonthExpression, models.Step.Value()) {
		return values
	}
	parts := strings.FieldsFunc(c.MonthExpression, func(r rune) bool {
		return r == '/'
	})
	step, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(values); i++ {
		if values[i]%step == values[0]%step {
			months = append(months, values[i])
		}
	}

	return months
}

func (c *CronMonthParser) RangeParser(values []int) []int {
	var months []int
	if !strings.Contains(c.MonthExpression, models.Range.Value()) {
		return values
	}
	parts := strings.FieldsFunc(c.MonthExpression, func(r rune) bool {
		return r == '-' || r == '/'
	})
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	for i := 0; i < len(values); i++ {
		if values[i] >= start && values[i] <= end {
			months = append(months, values[i])
		}
	}
	return months
}

func (c *CronMonthParser) Parse() []int {

	x := c.WildCardParser()
	f := c.NumberParser(x)
	y := c.RangeParser(f)
	z := c.StepParser(y)

	return z
}
