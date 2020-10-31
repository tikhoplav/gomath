package gomath

import (
	"testing"
)

func TestEquals(t *testing.T) {
	var a = 0.0001
	var b = 0.0001
	if !Equals(a, b) {
		f := `Expected numbers to be equal
		%v and
		%v`
		t.Errorf(f, a, b)
	}
	b = 0.00011
	if Equals(a, b) {
		f := `Expected numbers to be unequal
		%v and
		%v,
		with epsilon %v`
		t.Errorf(f, a, b, EPSILON)
	}
	if !Equals(a, b, 0.0001) {
		f := `Expected numbers to be equal
		%v and
		%v,
		with epsilon %v`
		t.Errorf(f, a, b, 0.0001)
	}
}

func TestSqrt(t *testing.T) {
	var a = 4.0
	var b = 2.0
	if !Equals(Sqrt(a), b) {
		f := `Expected square root of %v
		to be equal %v,
		got %v`
		t.Errorf(f, a, b, Sqrt(a))
	}
	a = 9.0
	b = 3.0
	if !Equals(Sqrt(a), b) {
		f := `Expected square root of %v
		to be equal %v,
		got %v`
		t.Errorf(f, a, b, Sqrt(a))
	}
	a = 10.0
	b = 3.162277
	if !Equals(Sqrt(a), b) {
		f := `Expected square root of %v
		to be equal %v,
		got %v`
		t.Errorf(f, a, b, Sqrt(a))
	}
}