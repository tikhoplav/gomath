package gomath

// Represents 3x3 matrix. New instance can be created as: var m = Mat3{0, 1, 2}
// All omitted components will be set to 0 by default. For optimisation purpose
// indecies are represented by a single number from 0 to 8, in contrast with
// common math notation, where components are specified by two indecies. Each
// component is represented by the floating point number and can be accessed by
// it's index, for example: var a21 = m[3]. As well as can be set explicitly:
// m[3] = 3.14.
type Mat3 [9]float64

// Returns true if matrices are equal and false if matrices are unequal. Second
// parameter epsilon can be used to define precision of the comparison. Default
// value of epsion is set by constant EPSILON. In order to prevent error caused
// by floating point number representation, calculations done approximately.
func (a Mat3) Equals(b Mat3, e ...float64) bool {
	res := true
	for i := 0; i < 9; i++ {
		res = res && Equals(a[i], b[i], e...)
	}
	return res
}

// Returns a determinant of a matrix.
func (a Mat3) Det() float64 {
	return a[0]*(a[4]*a[8]-a[5]*a[7]) -
		a[1]*(a[3]*a[8]-a[5]*a[6]) +
		a[2]*(a[3]*a[7]-a[4]*a[6])
}

// Returns scaled version of matrix by given scalar.
func (a Mat3) Scale(s float64) Mat3 {
	return Mat3{
		s * a[0], s * a[1], s * a[2],
		s * a[3], s * a[4], s * a[5],
		s * a[6], s * a[7], s * a[8],
	}
}

// Returns transposed matrix to the given.
func (a Mat3) Trans() Mat3 {
	return Mat3{
		a[0], a[3], a[6],
		a[1], a[4], a[7],
		a[2], a[5], a[8],
	}
}

// Returns inverse matrix to the given.
func (a Mat3) Inv() Mat3 {
	var i = Mat3{
		a[4]*a[8] - a[5]*a[7], a[5]*a[6] - a[3]*a[8], a[3]*a[7] - a[4]*a[6],
		a[2]*a[7] - a[1]*a[8], a[0]*a[8] - a[2]*a[6], a[1]*a[6] - a[0]*a[7],
		a[1]*a[5] - a[2]*a[4], a[2]*a[3] - a[0]*a[5], a[0]*a[4] - a[1]*a[3],
	}
	return i.Trans().Scale(1 / a.Det())
}
