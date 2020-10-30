package mat3

import (
	"testing"
)

func TestEquals(t *testing.T) {
	var a = New(1, 2, 3)
	var b = New(1, 2, 3, 0, 0, 0)

	if !Equals(a, b) {
		t.Errorf("Matrices %v and %v expected to be equal", a, b)
	}
}

func TestDeterminant(t *testing.T) {
	var a = New(1, 3, 2, -3, -1, -3, 2, 3, 1)

	if Determinant(a) != -15 {
		t.Errorf("Determinant of %v expected to be equal -15, got %v", a, Determinant(a))
	}
}

func TestTranspose(t *testing.T) {
	var a = New(1, 2, 3, 4, 5, 6, 7, 8, 9)
	var b = New(1, 4, 7, 2, 5, 8, 3, 6, 9)

	if !Equals(Transpose(a), b) {
		t.Errorf("Transposed matrix for %v expected to be equal %v, got %v", a, b, Transpose(a))
	}
}

func TestInverse(t *testing.T) {
	var a = New(7, 2, 1, 0, 3, -1, -3, 4, -2)
	var b = New(-2, 8, -5, 3, -11, 7, 9, -34, 21)
	
	var i = Inverse(a)
	if !Equals(i, b) {
		t.Errorf("Inverse matrix for %v expected to be %v, got %v", a, b, i)
	}
}