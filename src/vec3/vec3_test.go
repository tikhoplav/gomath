package vec3

import (
	"testing"
)

func TestEquals(t *testing.T) {
	var a = New(1, 0, 0)
	var b = New(1, 0, 0)

	if !Equals(a, b) {
		t.Errorf("Vectors: %v, %v expected to be equal", a, b)
	}

	b = New(0, 1, 0)

	if Equals(a, b) {
		t.Errorf("Vectors: %v, %v expected to be unequal", a, b)
	}
}

func TestEqualsApprox(t *testing.T) {
	var a = New(1, 0, 0)
	var b = New(1, 0.0001, 0.0001)
	var e = 0.001

	if !EqualsApprox(a, b, e) {
		t.Errorf("Vectors: %v, %v expected equal", a, b)
	}

	b = New(1, 0.002, 0)

	if EqualsApprox(a, b, e) {
		t.Errorf("Vectors: %v, %v expected unequal", a, b)
	}
}

func TestSquaredLength(t *testing.T) {
	var a = New(1, 0, 0)

	if Len2(a) != 1 {
		t.Errorf("Vector %v expected squared length is: 1, got %v", a, Len2(a))
	}

	a = New(2, 1, 1)

	if Len2(a) != 6 {
		t.Errorf("Vector %v expected squared length is: 6, got %v", a, Len2(a))
	}
}

func TestLength(t *testing.T) {
	var a = New(1, 0, 0)

	if Len(a) != 1 {
		t.Errorf("Vector %v expected squared length is: 1, got %v", a, Len(a))
	}

	a = New(2, 1, 2)

	if Len(a) != 3 {
		t.Errorf("Vector %v expected squared length is: 3, got %v", a, Len(a))
	}
}

func TestNormalized(t *testing.T) {
	var a = New(2, 1, 2)
	var n = New(0.666, 0.333, 0.666)

	if !EqualsApprox(Normalize(a), n, 0.001) {
		t.Errorf("Vector %v expected normalized version is %v, got %v", a, n, Normalize(a))
	}
}

func TestScale(t *testing.T) {	
	var a = New(2, 4, 0)
	var b = New(1, 2, 0)

	if !Equals(Scale(a, 0.5), b) {
		t.Errorf("Vector %v scaled by 0.5 expected to be %v, got %v", a, b, Scale(a, 0.5))
	}
}

func TestAdd(t *testing.T) {
	var a = New(1, 0, 0)
	var b = New(0, 2, 0)
	var c = New(1, 2, 0)

	if !Equals(Add(a, b), c) {
		t.Errorf("Vector %v added vector %v expected to be %v, got %v", a, b, c, Add(a, b))
	}
}

func TestSub(t *testing.T) {
	var a = New(1, 0, 0)
	var b = New(0, 2, 0)
	var c = New(1, -2, 0)

	if !Equals(Sub(a, b), c) {
		t.Errorf("Vector %v substracted vector %v expected to be %v, got %v", a, b, c, Sub(a, b))
	}
}

func TestDot(t *testing.T) {
	var a = New(2, 1, 0)
	var b = New(-1, 2, 0)

	if Dot(a, b) != 0 {
		t.Errorf("Expected dot product of vectors %v, %v is equal to 0, got %v", a, b, Dot(a, b))
	}

	// Since dot product with normalized vector is a projection
	// We could predict the result of dot by using basis vector
	a = New(1, 2, 0)
	b = New(1, 0, 0)

	if Dot(a, b) != 1 {
		t.Errorf("Expected dot product of vectors %v, %v is equal to 1, got %v", a, b, Dot(a, b))
	}
}

func TestCross(t *testing.T) {
	var a = New(1, 0, 0)
	var b = New(0, 1, 0)
	var c = New(0, 0, 1)

	if !Equals(Cross(a, b), c) {
		t.Errorf("Expected cross product of %v and %v to be %v, got %v", a, b, c, Cross(a, b))
	}
}

func TestProject(t *testing.T) {
	var a = New(1, 2, 0)
	var b = New(-0.5, 1, 0)
	var c = New(-0.6, 1.2, 0)

	if !Equals(Project(a, b), c) {
		t.Errorf("Expected projection of %v to %v be equal to %v, got %v", a, b, c, Project(a, b))
	}
}