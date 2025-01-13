package main

type ICronDayOfWeekParser interface {
	WildCardParser() []int
	NumberParser(values []int) []int
	StepParser(values []int) []int
	RangeParser(values []int) []int
	Parse() []string
}
