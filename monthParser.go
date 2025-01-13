package main

type ICronMonthParser interface {
	WildCardParser() []int
	NumberParser(values []int) []int
	StepParser(values []int) []int
	RangeParser(values []int) []int
	Parse() []int
}
