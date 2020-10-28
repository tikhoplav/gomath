package vec3

import (
	"math"
)

type Vec3 [3]float64

// Creates new vector from x,y and z components.
func New(x float64, y float64, z float64) Vec3 {
	return Vec3{x, y, z}
}

// Checks if specified vectors are equal. In order to prevent errors casued
// by floating point number representation, it's recomended to use EqualsApprox
// function instead.
func Equals(a Vec3, b Vec3) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2]
}

// Checks if differences between vectors components less then specified epsilon
// This function may prevent falty denials caused by floating point number
// representation.
func EqualsApprox(a Vec3, b Vec3, e float64) bool {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	dz := a[2] - b[2]
	return math.Abs(dx) <= e && math.Abs(dy) <= e && math.Abs(dz) <= e
}

// Returns vectors squared magnitude. This function recomended to be used where
// possible, since lower computational cost compared to Len function.
func Len2(a Vec3) float64 {
	return a[0]*a[0] + a[1]*a[1] + a[2]*a[2]
}

// Returns magnitude of specified vector.
func Len(a Vec3) float64 {
	return math.Sqrt(Len2(a))
}

// Returns scaled version of given vector by the given scalar.
func Scale(a Vec3, f float64) Vec3 {
	return Vec3{a[0] * f, a[1] * f, a[2] * f}
}

// Returns normalized version of given vector.
func Normalize(a Vec3) Vec3 {
	return Scale(a, 1/Len(a))
}

// Returns product of vectors addition.
func Add(a Vec3, b Vec3) Vec3 {
	return Vec3{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
	}
}

// Returns product of vectors substraction.
func Sub(a Vec3, b Vec3) Vec3 {
	return Vec3{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
	}
}

// Returns dot product of specified vectors.
func Dot(a Vec3, b Vec3) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

// Returns cross product of specified vectors.
func Cross(a Vec3, b Vec3) Vec3 {
	return Vec3{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[2],
	}
}

// Returns projection of given vector A to an axis specified with vector B.
func Project(a Vec3, b Vec3) Vec3 {
	var n = Normalize(b)
	var l = Dot(a, n)
	return Scale(n, l)
}
