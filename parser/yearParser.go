package parser

import (
	"cronParser/models"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CronYearParser struct {
	YearExpression string
}

const (
	maxYearLimit = 2100
)

func (c *CronYearParser) WildCardParser() []int {
	var years []int
	if len(c.YearExpression) == 0 {
		return years
	}
	currentYear := time.Now().Year()
	for i := currentYear; i <= maxYearLimit; i++ {
		years = append(years, i)
	}

	return years
}

// -
func (c *CronYearParser) RangeParser(values []int) []int {
	if !strings.Contains(c.YearExpression, models.Range.Value()) || len(c.YearExpression) == 0 {
		return values
	}

	parts := strings.FieldsFunc(c.YearExpression, func(r rune) bool {
		return r == '-'
	})

	startRange, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println(err)
	}
	endRange, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println(err)
	}

	var years []int
	for i := 0; i < len(values); i++ {
		if values[i] >= startRange && values[i] <= endRange {
			years = append(years, values[i])
		}
	}

	return years

}

func (c *CronYearParser) Parse() []int {
	var years []int
	if len(c.YearExpression) == 0 {
		return years
	}

	allYears := c.WildCardParser()
	rangeValues := c.RangeParser(allYears)

	return rangeValues
}
