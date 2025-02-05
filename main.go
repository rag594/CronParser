package main

import (
	"cronParser/models"
	"cronParser/parser"
	"os"
	"strings"
)

func main() {
	cronArgs := os.Args[1:]
	cSplitStr := strings.Split(cronArgs[0], " ")

	command := strings.Join(cronArgs[1:], " ")

	cronCommand := models.CronCommand{
		Expression: models.CronExpression{
			Minute:     cSplitStr[0],
			Hour:       cSplitStr[1],
			DayOfMonth: cSplitStr[2],
			Month:      cSplitStr[3],
			DayOfWeek:  cSplitStr[4],
			Year: func() string {
				if len(cSplitStr) <= 5 {
					return ""
				}
				return cSplitStr[5]
			}(),
		},
		Command: command,
	}

	// 34/4 */4 1,15 * 1-4 2025,2026 or 2025-2027 (list and range)

	cronExpressionParser := NewCronExpressionParser(
		WithCronMinuteParser(&parser.CronMinuteParser{MinuteExpression: cronCommand.Expression.Minute}),
		WithCronHourParser(&parser.CronHourParser{HourExpression: cronCommand.Expression.Hour}),
		WithCronDayOfMonthParser(&parser.CronDayOfMonthParser{DayOfMonthExpression: cronCommand.Expression.DayOfMonth}),
		WithCronMonthParser(&parser.CronMonthParser{MonthExpression: cronCommand.Expression.Month}),
		WithCronDayOfWeekParser(&parser.CronDayOfWeekParser{DayOfWeekExpression: cronCommand.Expression.DayOfWeek}),
		WithCronYearParser(&parser.CronYearParser{YearExpression: cronCommand.Expression.Year}),
	)

	cronExpressionParser.Display(cronCommand.Command)

}
