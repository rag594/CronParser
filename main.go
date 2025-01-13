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
	cronCommand := models.CronCommand{
		Expression: models.CronExpression{
			Minute:     cSplitStr[0],
			Hour:       cSplitStr[1],
			DayOfMonth: cSplitStr[2],
			Month:      cSplitStr[3],
			DayOfWeek:  cSplitStr[4],
		},
		Command: cronArgs[1],
	}

	cronExpressionParser := NewCronExpressionParser(
		WithCronMinuteParser(&parser.CronMinuteParser{MinuteExpression: cronCommand.Expression.Minute}),
		WithCronHourParser(&parser.CronHourParser{HourExpression: cronCommand.Expression.Hour}),
		WithCronDayOfMonthParser(&parser.CronDayOfMonthParser{DayOfMonthExpression: cronCommand.Expression.DayOfMonth}),
		WithCronMonthParser(&parser.CronMonthParser{MonthExpression: cronCommand.Expression.Month}),
		WithCronDayOfWeekParser(&parser.CronDayOfWeekParser{DayOfWeekExpression: cronCommand.Expression.DayOfWeek}),
	)

	cronExpressionParser.Display(cronCommand.Command)

}
