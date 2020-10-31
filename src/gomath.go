package gomath

import (
	"math"
)

// Represents default precision of floating point number equality calculation.
// Two numbers are considered equal if absolute value of their substraction is
// less or equal specified epsilon: |x - y| <= 0. Approximate equality is used
// in order to prevent errors caused by floating point number representation.
const EPSILON = 0.000001

// Returns true if two numbers are are equal within the specified precision.
// Third parameter may be used to specify epsilon. If none provided default
// value EPSILON is used.
func Equals(a float64, b float64, args ...float64) bool {
	var e = EPSILON
	if len(args) > 0 {
		e = args[0]
	}
	return math.Abs(a - b) <= e
}

// Return square root of the given number.
func Sqrt(a float64) float64 {
	// Implemetation of sqrt function may be changed in future, in order to
	// optimize a computational time. Despite this, function structre will not
	// be changed.
	return math.Sqrt(a)
}