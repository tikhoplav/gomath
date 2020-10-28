package vec3

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestCos2(t *testing.T) {
	// Set angle between vectors to pi/4 apriori
	var a = New(1, 0, 0)
	var b = New(1, 1, 0)

	if Cos2(a, b) != 0.5 {
		t.Errorf("Expected squared cos between %v and %v to be 0.5, got %v", a, b, Cos2(a, b))
	}
}

func TestSin2(t *testing.T) {
	// Set angle between vectors to pi/4 apriori
	var a = New(1, 0, 0)
	var b = New(1, 1, 0)

	if Sin2(a, b) != 0.5 {
		t.Errorf("Expected squared cos between %v and %v to be 0.5, got %v", a, b, Sin2(a, b))
	}
}

func TestCosTheorem(t *testing.T) {
	err := 0.0
	n := 10
	e := 0.000001

	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())

		var a = New(rand.Float64(), rand.Float64(), rand.Float64())
		var b = New(rand.Float64(), rand.Float64(), rand.Float64())

		var c = Sub(a, b)

		err = err + Len2(c) - Len2(a) - Len2(b) + 2*Len(a)*Len(b)*Cos(a, b)
	}

	err = err / float64(n)

	if math.Abs(err) > e {
		t.Errorf("Cosine law is violated with epsi: %v", e)
	}
}

func TestSinTheorem(t *testing.T) {
	err := 0.0
	n := 10
	e := 0.000001

	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())

		var a = New(rand.Float64(), rand.Float64(), rand.Float64())
		var b = New(rand.Float64(), rand.Float64(), rand.Float64())

		var c = Sub(a, b)

		err = err + Sin2(a, b)/Len2(c) - Sin2(a, c)/Len2(b)
	}

	err = err / float64(n)

	if math.Abs(err) > e {
		t.Errorf("Sine law is violated with epsi: %v", e)
	}
}

func TestPythagorianTheorem(t *testing.T) {
	err := 0.0
	n := 10
	e := 0.000001

	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())

		var a = New(rand.Float64(), rand.Float64(), rand.Float64())
		var b = New(rand.Float64(), rand.Float64(), rand.Float64())

		err = err + 1 - Sin2(a, b) - Cos2(a, b)
	}

	err = err / float64(n)

	if math.Abs(err) > e {
		t.Errorf("Sine law is violated with epsi: %v", e)
	}
}
