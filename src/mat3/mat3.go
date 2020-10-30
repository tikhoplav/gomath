package mat3

import (
	"math"
)

type Mat3 [9]float64

// Returns new instance of 3x3 matrix.
// If number of arguments is less then 9, all other matrix components will be
// set to 0 by default. Arguments with number more then 9 will be ignored.
// It is possible to create empty matrix by providing none arguments.
func New(values ...float64) Mat3 {
	for i := len(values); i < 9; i++ {
		values = append(values, 0)
	}

	return Mat3{
		values[0], values[1], values[2],
		values[3], values[4], values[5],
		values[6], values[7], values[8],
	}
}

// Checks if matrices are identical. In order to prevent falty denials
// function EqualsApprox can be used instead.
func Equals(a Mat3, b Mat3) bool {
	res := true
	for i := 0; i < 9; i++ {
		res = res && a[i] == b[i]
	}
	return res
}

// Checks if components of specified matrices differs not more then provided
// epsilon. This function is recomended to use instead of regular Equals in
// order to prevent falty denials due to floating point number representation.
func EqualsApprox(a Mat3, b Mat3, e float64) bool {
	res := true
	for i := 0; i < 9; i++ {
		res = res && math.Abs(a[i]-b[i]) <= e
	}
	return res
}

// Returns a determinant of a matrix.
func Determinant(a Mat3) float64 {
	return a[0]*(a[4]*a[8] - a[5]*a[7]) -
		a[1]*(a[3]*a[8] - a[5]*a[6]) +
		a[2]*(a[3]*a[7] - a[4]*a[6])
}

// Returns scaled version of matrix by given scalar.
func Scale(a Mat3, s float64) Mat3 {
	return Mat3 {
		s*a[0], s*a[1], s*a[2],
		s*a[3], s*a[4], s*a[5],
		s*a[6], s*a[7], s*a[8],
	}
}

// Returns transposed matrix to the given.
func Transpose(a Mat3) Mat3 {
	return Mat3 {
		a[0], a[3], a[6],
		a[1], a[4], a[7],
		a[2], a[5], a[8],
	}
}

// Returns inverse matrix to the given.
func Inverse(a Mat3) Mat3 {
	var i = Mat3 {
		a[4]*a[8] - a[5]*a[7], a[5]*a[6] - a[3]*a[8], a[3]*a[7] - a[4]*a[6],
		a[2]*a[7] - a[1]*a[8], a[0]*a[8] - a[2]*a[6], a[1]*a[6] - a[0]*a[7],
		a[1]*a[5] - a[2]*a[4], a[2]*a[3] - a[0]*a[5], a[0]*a[4] - a[1]*a[3],
	}
	return Scale(Transpose(i), Determinant(a))
}
