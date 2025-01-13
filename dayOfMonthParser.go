package main

type ICronDayOfMonthParser interface {
	ListParser(values []int) []int
	WildCardParser() []int
	NumberParser(values []int) []int
	StepParser(values []int) []int
	RangeParser(values []int) []int
	Parse() []int
}
