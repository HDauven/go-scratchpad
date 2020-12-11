// Package hof - Higher order functions, showcases a number of such functions.
package hof

// Times is a HOF that takes a multiplier and returns a new function that uses this multiplier to multiply a newly given integer (the multiplicand) with the multiplier.
func Times(multiplier int) func(multiplicand int) int {
	return func(multiplicand int) int {
		return multiplier * multiplicand
	}
}
