package govector

import (
	"testing"
)

func TestVec3(t *testing.T) {
	var a, b, c Vec3

	a = Vec3{1, 0, 0};
	b = Vec3{0, 1, 0};
	c = Vec3{0, 0, 1.5};

	cross := Cross(a, b);

	if !cross.Equals(c) {
		t.Errorf("Expected cross product be: %v, got %v", c, cross)
	}
}