package linal

import "math"

type Vec3 struct {
	x float32
	y float32
	z float32
}

func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{a.x + b.x, a.y + b.y, a.z + b.z}
}
func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{a.x - b.x, a.y - b.y, a.z - b.z}
}
func (a Vec3) Mul(t float32) Vec3 {
	return Vec3{a.x * t, a.y * t, a.z * t}
}
func (a Vec3) Div(t float32) Vec3 {
	return Vec3{a.x / t, a.y / t, a.z / t}
}

func (a Vec3) Dot(b Vec3) float32 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{a.y*b.z - a.z*b.y, -(a.x*b.z - a.z*b.x), a.x*b.y - a.y*b.x}
}

func (a Vec3) LenSquared() float32 {
	return a.Dot(a)
}

func (a Vec3) Len() float32 {
	return float32(math.Sqrt(float64(a.LenSquared())))
}

func (a Vec3) Normalize() (Vec3, bool) {
	if a.x == 0.0 && a.y == 0.0 && a.z == 0.0 {
		return Vec3{}, false
	}
	len := a.Len()
	return Vec3{a.x / len, a.y / len, a.z / len}, true
}
