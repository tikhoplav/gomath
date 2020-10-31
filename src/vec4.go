package gomath

// Represents four dimensional vector. For clarity purposes component's names
// are set as x, y, z and w, and are represented by floating point numbers.
// New instance can be created as: var a = Vec4{x, y, z, w}. Each component can
// be set explicitly: a.X = 3.14. Vec4 also stands for Quatrenion, with real
// part represented by w.
type Vec4 struct {
	X, Y, Z, W float64
}

// Returns true if vectors are equal and false if vectors are unequal. Second
// parameter epsilon can be used to define precision of the comparison. Default
// value of epsion is set by constant EPSILON. In order to prevent error caused
// by floating point number representation, calculations done approximately.
func (a Vec4) Equals(b Vec4, e ...float64) bool {
	dx := Equals(a.X, b.X, e...)
	dy := Equals(a.Y, b.Y, e...)
	dz := Equals(a.Z, b.Z, e...)
	dw := Equals(a.W, b.W, e...)
	return dx && dy && dz && dw
}