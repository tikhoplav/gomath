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
