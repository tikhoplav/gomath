package gomath

import (
	"testing"
)

func TestVec3Equals(t *testing.T) {
	var a = Vec3{1, 0, 0}
	var b = Vec3{1, 0, 0}
	if !a.Equals(b) {
		f := `Vectors
		%v and
		%v
		expected to be equal`
		t.Errorf(f, a, b)
	}

	b = Vec3{0, 1, 0}
	if a.Equals(b) {
		f := `Vectors
		%v and
		%v
		expected to be unequal`
		t.Errorf(f, a, b)
	}
}

func TestVec3EqualsApprox(t *testing.T) {
	var a = Vec3{1, 0, 0}
	var b = Vec3{1, 0.001, 0.001}
	var e = 0.001
	if !a.EqualsApprox(b, e) {
		f := `Expected vectors
		%v and
		%v
		to be approximately equal with e = 0.001`
		t.Errorf(f, a, b)
	}

	e = 0.0001
	if a.EqualsApprox(b, e) {
		f := `Expected vectors
		%v and
		%v
		to be approximately unequal with e = 0.0001`
		t.Errorf(f, a, b)
	}
}

func TestVec3SquaredLength(t *testing.T) {
	var a = Vec3{1, 0, 0}
	if a.Len2() != 1 {
		f := `Squared lenght of vector %v
		expected to be 1,
		got %v`
		t.Errorf(f, a, a.Len2())
	}

	a = Vec3{2, 1, 1}
	if a.Len2() != 6 {
		f := `Squared lenght of vector %v
		expected to be 6,
		got %v`
		t.Errorf(f, a, a.Len2())
	}
}

func TestVec3Length(t *testing.T) {
	var a = Vec3{1, 0, 0}
	if a.Len() != 1 {
		f := `Lenght of vector %v
		expected to be 1,
		got %v`
		t.Errorf(f, a, a.Len())
	}

	a = Vec3{2, 1, 2}
	if a.Len() != 3 {
		f := `Lenght of vector %v
		expected to be 3,
		got %v`
		t.Errorf(f, a, a.Len())
	}
}

func TestVec3Scale(t *testing.T) {
	var a = Vec3{2, 4, 0}
	var b = Vec3{1, 2, 0}
	var c = a.Scale(0.5)
	if !c.Equals(b) {
		f := `Vector %v scaled by 0.5
		expected to be %v,
		got %v`
		t.Errorf(f, a, b, c)
	}
}

func TestVec3Normalized(t *testing.T) {
	var a = Vec3{2, 1, 2}
	var b = Vec3{0.666, 0.333, 0.666}
	var c = a.Normalize()
	if !c.EqualsApprox(b, 0.001) {
		f := `Vector %v normalized
		expect to be equal %v,
		got %v`
		t.Errorf(f, a, b, c)
	}
}

func TestVec3Add(t *testing.T) {
	var a = Vec3{1, 2, 3}
	var b = Vec3{4, 3, 2}
	var c = Vec3{5, 5, 5}
	if !a.Add(b).Equals(c) {
		f := `Additions of vectors
		%v and
		%v
		expected to be %v,
		got %v`
		t.Errorf(f, a, b, c, a.Add(b))
	}
}

func TestVec3Sub(t *testing.T) {
	var a = Vec3{1, 2, 3}
	var b = Vec3{4, 3, 2}
	var c = Vec3{-3, -1, 1}
	if !a.Sub(b).Equals(c) {
		f := `Substraction of vectors
		%v and
		%v
		expected to be %v,
		got %v`
		t.Errorf(f, a, b, c, a.Sub(b))
	}
}

func TestVec3Dot(t *testing.T) {
	var a = Vec3{1, 2, 3}
	var b = Vec3{4, 3, 2}
	if a.Dot(b) != 16 {
		f := `Dot product of vectors
		%v and
		%v
		expected to ve equal to 16,
		got %v`
		t.Errorf(f, a, b, a.Dot(b))
	}
}

func TestVec3Project(t *testing.T) {
	var a = Vec3{1, 2, 3}
	var b = Vec3{4, 3, 2}
	var c = Vec3{2.2, 1.65, 1.1}
	var p = a.Project(b)
	if !p.EqualsApprox(c, 0.01) {
		f := `Projection of vector
		%v onto vector
		%v
		expected to be %v,
		got %v`
		t.Errorf(f, a, b, c, p)
	}
}

func TestVec3Cross(t *testing.T) {
	var a = Vec3{1, 2, 3}
	var b = Vec3{4, 3, 2}
	var c = Vec3{-5, 10, -5}
	var r = a.Cross(b)
	if !r.Equals(c) {
		f := `Cross product of vectors
		%v and
		%v
		expected to be %v,
		got %v`
		t.Errorf(f, a, b, c, r)
	}
}

func TestVec3Transform(t *testing.T) {
	var a = Vec3{2, 5, -9}
	var b = Vec3{1, -2, 2}
	var m = Mat3{4, -5, -2, 5, -6, -2, -8, 9, 3}
	var c = a.Transform(m)
	if !c.Equals(b) {
		f := `Expected result of vector %v
		transformed by %v
		be equal to %v,
		got %v`
		t.Errorf(f, a, m, b, c)
	}
}
