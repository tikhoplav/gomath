package gomath

import (
	"testing"
)

func TestMat3Equals(t *testing.T) {
	var a = Mat3{1, 2, 3}
	var b = Mat3{1, 2, 3, 0, 0, 0}
	if !a.Equals(b) {
		f := `Next matrices expected to be equal
		%v and
		%v`
		t.Errorf(f, a, b)
	}
}

func TestMat3EqualsApprox(t *testing.T) {
	var a = Mat3{1, 0, 1}
	var b = Mat3{1.001, 0.001, 1.001}
	var e = 0.001
	if !a.EqualsApprox(b, e) {
		f := `Expected next matrices to be approximately equal
		with e = 0.001:
		%v and
		%v`
		t.Errorf(f, a, b)
	}

	e = 0.0001
	if a.EqualsApprox(b, e) {
		f := `Expected next matrices to be approximately unequal
		with e = 0.0001:
		%v and
		%v`
		t.Errorf(f, a, b)
	}
}

func TestMat3Determinant(t *testing.T) {
	var a = Mat3{1, 3, 2, -3, -1, -3, 2, 3, 1}
	if a.Det() != -15 {
		f := `Determinant of the next matrix expected to be equal -15:
		%v,
		got %v`
		t.Errorf(f, a, a.Det())
	}
}

func TestMat3Transpose(t *testing.T) {
	var a = Mat3{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var b = Mat3{1, 4, 7, 2, 5, 8, 3, 6, 9}
	var c = a.Trans()
	if !c.Equals(b) {
		f := `Transposed matrix of
		%v expected to be
		%v,
		got %v`
		t.Errorf(f, a, b, c)
	}
}

func TestMat3Inverse(t *testing.T) {
	var a = Mat3{0, 1, 2, 1, 4, 0, 0, 8, 6} // Det is 10
	var b = Mat3{2.4, 1, -0.8, -0.6, 0, 0.2, 0.8, 0, -0.1}
	var c = a.Inv()
	if !c.EqualsApprox(b, 0.001) {
		f := `Inverse matrix to
		%v expected to be
		%v,
		got %v`
		t.Errorf(f, a, b, c)
	}
}
