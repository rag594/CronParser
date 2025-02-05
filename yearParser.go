package main

type ICronYearParser interface {
	WildCardParser() []int
	RangeParser(values []int) []int
	Parse() []int
}
