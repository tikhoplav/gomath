package vec3

import (
	"math"
)

// Returns squared cos of angle between two specified vectors.
// It is recommended to use this function instead of regular Cos where
// possible due to less number of required calculations. 
func Cos2(a Vec3, b Vec3) float64 {
	var dot = Dot(a, b)
	return dot * dot / ( Len2(a) * Len2(b) )
}

// Returns squared sin of angle between two specified vectors.
// It is recommended to use this function instead of regular Sin where
// possible due to less number of required calculations. 
func Sin2(a Vec3, b Vec3) float64 {
	return 1 - Cos2(a, b)
}

// Returns cos of angle between two specified vectors.
func Cos(a Vec3, b Vec3) float64 {
	return math.Sqrt(Cos2(a, b))
}

// Returns sin of angle between two specified vectors.
func Sin(a Vec3, b Vec3) float64 {
	return math.Sqrt(Sin2(a, b))
}