package main

import "fmt"

// CronExpressionParser parsers for different parts in the cron expression
// each part has a different parser and is assigned an interface which helps
// in plug-n-play model of using different implementations
type CronExpressionParser struct {
	Minute     ICronMinuteParser
	Hour       ICronHourParser
	DayOfMonth ICronDayOfMonthParser
	Month      ICronMonthParser
	DayOfWeek  ICronDayOfWeekParser
	Year       ICronYearParser
}

func (c *CronExpressionParser) Display(command string) {
	fmt.Printf("minute         %v\n", c.Minute.Parse())
	fmt.Printf("hour           %v\n", c.Hour.Parse())
	fmt.Printf("day of month   %v\n", c.DayOfMonth.Parse())
	fmt.Printf("month          %v\n", c.Month.Parse())
	fmt.Printf("day of week    %v\n", c.DayOfWeek.Parse())
	if len(c.Year.Parse()) != 0 {
		fmt.Printf("Year    %v\n", c.Year.Parse())
	}
	fmt.Printf("command        %v\n", command)
}

// CronExpressionParserOptions uses the optional pattern to help in plug-n-play model of which part to use and not
type CronExpressionParserOptions func(parser *CronExpressionParser)

func NewCronExpressionParser(opts ...CronExpressionParserOptions) *CronExpressionParser {
	cp := &CronExpressionParser{}

	for _, opt := range opts {
		opt(cp)
	}

	return cp
}

func WithCronMinuteParser(minuteParser ICronMinuteParser) CronExpressionParserOptions {
	return func(parser *CronExpressionParser) {
		parser.Minute = minuteParser
	}
}

func WithCronHourParser(hourParser ICronHourParser) CronExpressionParserOptions {
	return func(parser *CronExpressionParser) {
		parser.Hour = hourParser
	}
}

func WithCronDayOfMonthParser(dayOfMonthParser ICronDayOfMonthParser) CronExpressionParserOptions {
	return func(parser *CronExpressionParser) {
		parser.DayOfMonth = dayOfMonthParser
	}
}

func WithCronMonthParser(monthParser ICronMonthParser) CronExpressionParserOptions {
	return func(parser *CronExpressionParser) {
		parser.Month = monthParser
	}
}

func WithCronDayOfWeekParser(dayOfWeekParser ICronDayOfWeekParser) CronExpressionParserOptions {
	return func(parser *CronExpressionParser) {
		parser.DayOfWeek = dayOfWeekParser
	}
}

func WithCronYearParser(yearParser ICronYearParser) CronExpressionParserOptions {
	return func(parser *CronExpressionParser) {
		parser.Year = yearParser
	}
}
