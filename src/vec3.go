package gomath

import (
	"github.com/tikhoplav/gomath/vec3"
	"github.com/tikhoplav/gomath/mat3"
)

func Vec3ApplyMat3(v vec3.Vec3, m mat3.Mat3) vec3.Vec3 {
	return vec3.Vec3 {
		m[0]*v[0] + m[1]*v[1] + m[2]*v[2],
		m[3]*v[0] + m[4]*v[1] + m[5]*v[2],
		m[6]*v[0] + m[7]*v[1] + m[8]*v[2],
	}
}