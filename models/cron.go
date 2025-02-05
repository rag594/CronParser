package models

type CronSymbol string

func (c CronSymbol) Value() string {
	return string(c)
}

const (
	WildCard CronSymbol = "*"
	Range    CronSymbol = "-"
	Step     CronSymbol = "/"
	List     CronSymbol = ","
)

type CronCommand struct {
	Expression CronExpression
	Command    string
}

type CronExpression struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
	Year       string
}

func NewCronExpression(minute string) *CronExpression {
	return &CronExpression{
		Minute:     minute,
		Hour:       "",
		DayOfMonth: "",
		Month:      "",
		DayOfWeek:  "",
	}
}
