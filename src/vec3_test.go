package gomath

import (
	"github.com/tikhoplav/gomath/vec3"
	"github.com/tikhoplav/gomath/mat3"
	"testing"
)

func TestMatrix3AppliedToVector3(t *testing.T) {
	var a = vec3.New(21, 5, -1)
	var b = vec3.New(3, 1, -2)
	var m = mat3.New(-2, 8, -5, 3, -11, 7, 9, -34, 21)

	var c = Vec3ApplyMat3(a, m)

	if !vec3.Equals(c, b) {
		t.Errorf("Matrix %v applyed to vector %v expected to be %v, got %v", m, a, b, c)
	}
}