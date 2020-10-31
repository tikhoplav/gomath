package gomath

import (
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

// Checks if specified vectors are equal. In order to prevent errors casued by
// floating point number representation, it's recomended to use EqualsApprox
// function instead.
func (a Vec3) Equals(b Vec3) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}

// Checks if differences between vectors components less then specified epsilon
// This function may prevent falty denials caused by floating point number
// representation.
func (a Vec3) EqualsApprox(b Vec3, e float64) bool {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return math.Abs(dx) <= e && math.Abs(dy) <= e && math.Abs(dz) <= e
}

// Returns vectors squared magnitude. This function recomended to be used where
// possible, since lower computational cost compared to Len function.
func (a Vec3) Len2() float64 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

// Returns magnitude of specified vector.
func (a Vec3) Len() float64 {
	return math.Sqrt(a.Len2())
}

// Returns scaled version of given vector by the given scalar.
func (a Vec3) Scale(f float64) Vec3 {
	return Vec3{a.X * f, a.Y * f, a.Z * f}
}

// Returns normalized version of given vector.
func (a Vec3) Normalize() Vec3 {
	return a.Scale(1 / a.Len())
}

// Returns product of vectors addition.
func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Returns product of vectors subtraction.
func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Returns dot product of specified vectors.
func (a Vec3) Dot(b Vec3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Returns projection of given vector onto an axis specified with vector B.
func (a Vec3) Project(b Vec3) Vec3 {
	var axis = b.Normalize()
	var length = a.Dot(axis)
	return axis.Scale(length)
}

// Returns cross product of specified vectors.
func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

// Returns squared cosine of angle between two vectors. It is recommended to
// use Cos2 instead of Cos where possible due to less number of calculations.
func (a Vec3) Cos2(b Vec3) float64 {
	var d = a.Dot(b)
	return d * d / (a.Len2() * b.Len2())
}

// Returns squared sine of angle between two vectors. It is recommended to use
// Sin2 instead of Sin where possible due to less number of calculations.
// Sin2 is calculated based on Pyphagorean theorem (sin2 + cos2 = 1).
func (a Vec3) Sin2(b Vec3) float64 {
	return 1 - a.Cos2(b)
}

// Returns squared tangent of angle between two vectors. It is recommended
// to use Tan2 instead of Tan where possible due to less calculations required.
// Sin2 is calculated based on Pyphagorean theorem (sin2 + cos2 = 1).
func (a Vec3) Tan2(b Vec3) float64 {
	var cos2 = a.Cos2(b)
	return (1 - cos2) / cos2
}

// Returns squared cotangent of angle between two vectors. It is recommended
// to use Cot2 instead of Cot where possible due to less calculations required.
// Sin2 is calculated based on Pyphagorean theorem (sin2 + cos2 = 1).
func (a Vec3) Cot2(b Vec3) float64 {
	var cos2 = a.Cos2(b)
	return cos2 / (1 - cos2)
}

// Returns cosine of angle between two vectors.
func (a Vec3) Cos(b Vec3) float64 {
	return math.Sqrt(a.Cos2(b))
}

// Returns sine of angle between two vectors.
func (a Vec3) Sin(b Vec3) float64 {
	return math.Sqrt(a.Sin2(b))
}

// Returns tangent of angle between two vectors
func (a Vec3) Tan(b Vec3) float64 {
	return math.Sqrt(a.Tan2(b))
}

// Return cotangent of angle between two vectors
func (a Vec3) Cot(b Vec3) float64 {
	return math.Sqrt(b.Cot2(b))
}

// Return product of matrix application to the vector. Corresponds to math
// notation as y = Ax, where A - is a matrix, x - is given vector. If multiple
// transformation should be applied, for example y = ABCx, next sequence should
// be used: var y = x.Transform(C).Transform(B).Transform(A)
func (a Vec3) Transform(m Mat3) Vec3 {
	return Vec3{
		a.X*m[0] + a.Y*m[1] + a.Z*m[2],
		a.X*m[3] + a.Y*m[4] + a.Z*m[5],
		a.X*m[6] + a.Y*m[7] + a.Z*m[8],
	}
}
