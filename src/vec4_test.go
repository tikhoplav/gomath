package gomath

import (
	"testing"
)

func TestVec4Equals(t *testing.T) {
	var a = Vec4{4, 2, 3, 1}
	var b = Vec4{4, 2, 3, 1}
	if !a.Equals(b) {
		f := `Next vectors expected to be equal:
		%v and
		%v`
		t.Errorf(f, a, b)
	}
	b = Vec4{4.01, 2, 3.01, 1}
	if a.Equals(b) {
		f := `Next vectors expected to be unequal:
		%v and
		%v`
		t.Errorf(f, a, b)
	}
	if !a.Equals(b, 0.01) {
		f := `Next vectors expected to be equal:
		%v and
		%v,
		with epsilon = 0.1`
		t.Errorf(f, a, b)
	}
}