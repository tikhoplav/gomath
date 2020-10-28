package govector

type Vec3 [3]float64

func (a Vec3) Equals(b Vec3) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2]
}

func Cross(a Vec3, b Vec3) Vec3 {
	return Vec3{
		a[1] * b[2] - a[2] * b[1],
		a[2] * b[0] - a[0] * b[2],
		a[0] * b[1] - a[1] * b[0],
	}
}