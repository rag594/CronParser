package parser

import (
	"cronParser/models"
	"fmt"
	"strconv"
	"strings"
)

type CronMinuteParser struct {
	MinuteExpression string
}

func (c *CronMinuteParser) WildCardParser() []int {
	var minutes []int

	for i := 0; i <= 59; i++ {
		minutes = append(minutes, i)
	}

	return minutes
}

func (c *CronMinuteParser) NumberParser(values []int) []int {
	var minutes []int

	if strings.Contains(c.MinuteExpression, models.Range.Value()) ||
		strings.Contains(c.MinuteExpression, models.WildCard.Value()) {
		return values
	}

	parts := strings.FieldsFunc(c.MinuteExpression, func(r rune) bool {
		return r == '/'
	})

	number, _ := strconv.Atoi(parts[0])

	if strings.Contains(c.MinuteExpression, models.Step.Value()) {
		for i := 0; i < len(values); i++ {
			if values[i] >= number {
				minutes = append(minutes, values[i])
			}
		}
		return minutes
	}

	return []int{number}
}

func (c *CronMinuteParser) StepParser(values []int) []int {
	var minutes []int

	if !strings.Contains(c.MinuteExpression, models.Step.Value()) {
		return values
	}
	parts := strings.FieldsFunc(c.MinuteExpression, func(r rune) bool {
		return r == '/'
	})
	step, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(values); i++ {
		if values[i]%step == values[0]%step {
			minutes = append(minutes, values[i])
		}
	}

	return minutes
}

func (c *CronMinuteParser) RangeParser(values []int) []int {
	var minutes []int
	if !strings.Contains(c.MinuteExpression, models.Range.Value()) {
		return values
	}
	parts := strings.FieldsFunc(c.MinuteExpression, func(r rune) bool {
		return r == '-' || r == '/'
	})
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	for i := 0; i < len(values); i++ {
		if values[i] >= start && values[i] <= end {
			minutes = append(minutes, values[i])
		}
	}
	return minutes
}

func (c *CronMinuteParser) Parse() []int {

	x := c.WildCardParser()
	f := c.NumberParser(x)
	y := c.RangeParser(f)
	z := c.StepParser(y)

	return z
}
