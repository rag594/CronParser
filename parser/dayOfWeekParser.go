package parser

import (
	"cronParser/models"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

type CronDayOfWeekParser struct {
	DayOfWeekExpression string
	DayOfWeekInFull     bool
}

var (
	NumToDayOfWeekEncoding = map[int]string{
		1: "MON",
		2: "TUE",
		3: "WED",
		4: "THU",
		5: "FRI",
		6: "SAT",
		7: "SUN",
	}
	DayOfWeekToNumEncoding = map[string]int{
		"MON": 1,
		"TUE": 2,
		"WED": 3,
		"THU": 4,
		"FRI": 5,
		"SAT": 6,
		"SUN": 7,
	}
	DaysOfWeekInStrSlice = slices.Collect(maps.Keys(DayOfWeekToNumEncoding))
)

func (c *CronDayOfWeekParser) WildCardParser() []int {
	var dayOfWeeks []int

	for i := 1; i <= 7; i++ {
		dayOfWeeks = append(dayOfWeeks, i)
	}

	return dayOfWeeks
}

func (c *CronDayOfWeekParser) NumberParser(values []int) []int {
	var dayOfWeeks []int
	var number int

	if strings.Contains(c.DayOfWeekExpression, models.Range.Value()) ||
		strings.Contains(c.DayOfWeekExpression, models.WildCard.Value()) {
		return values
	}

	parts := strings.FieldsFunc(c.DayOfWeekExpression, func(r rune) bool {
		return r == '/'
	})

	if slices.Contains(DaysOfWeekInStrSlice, parts[0]) {
		c.DayOfWeekInFull = true
		number = DayOfWeekToNumEncoding[parts[0]]
	} else {
		number, _ = strconv.Atoi(parts[0])
	}

	if strings.Contains(c.DayOfWeekExpression, models.Step.Value()) {
		for i := 0; i < len(values); i++ {
			if values[i] >= number {
				dayOfWeeks = append(dayOfWeeks, values[i])
			}
		}
		return dayOfWeeks
	}

	return []int{number}
}

func (c *CronDayOfWeekParser) StepParser(values []int) []int {
	var dayOfWeeks []int

	if !strings.Contains(c.DayOfWeekExpression, models.Step.Value()) {
		return values
	}
	parts := strings.FieldsFunc(c.DayOfWeekExpression, func(r rune) bool {
		return r == '/'
	})
	step, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(values); i++ {
		if values[i]%step == values[0]%step {
			dayOfWeeks = append(dayOfWeeks, values[i])
		}
	}

	return dayOfWeeks
}

func (c *CronDayOfWeekParser) RangeParser(values []int) []int {
	var dayOfWeeks []int
	var start, end int
	if !strings.Contains(c.DayOfWeekExpression, models.Range.Value()) {
		return values
	}
	parts := strings.FieldsFunc(c.DayOfWeekExpression, func(r rune) bool {
		return r == '-' || r == '/'
	})

	if slices.Contains(DaysOfWeekInStrSlice, parts[0]) {
		start = DayOfWeekToNumEncoding[parts[0]]
		end = DayOfWeekToNumEncoding[parts[1]]
		c.DayOfWeekInFull = true
	} else {
		start, _ = strconv.Atoi(parts[0])
		end, _ = strconv.Atoi(parts[1])
	}

	for i := 0; i < len(values); i++ {
		if values[i] >= start && values[i] <= end {
			dayOfWeeks = append(dayOfWeeks, values[i])
		}
	}
	return dayOfWeeks
}

func (c *CronDayOfWeekParser) Parse() []string {

	var strSlice []string

	x := c.WildCardParser()
	f := c.NumberParser(x)
	y := c.RangeParser(f)
	z := c.StepParser(y)

	if c.DayOfWeekInFull {
		for _, num := range z {
			strSlice = append(strSlice, NumToDayOfWeekEncoding[num])
		}
		return strSlice
	}

	for _, num := range z {
		strSlice = append(strSlice, strconv.Itoa(num))
	}

	return strSlice
}
