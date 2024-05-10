package linal

import "math"

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

func (a Vec3) Lerp(b Vec3, t float32) Vec3 {
	return a.Mul(1.0 - t).Add(b.Mul(t))
}

func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}
func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}
func (a Vec3) Mul(t float32) Vec3 {
	return Vec3{a.X * t, a.Y * t, a.Z * t}
}
func (a Vec3) Div(t float32) Vec3 {
	return Vec3{a.X / t, a.Y / t, a.Z / t}
}

func (a Vec3) Dot(b Vec3) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{a.Y*b.Z - a.Z*b.Y, -(a.X*b.Z - a.Z*b.X), a.X*b.Y - a.Y*b.X}
}

func (a Vec3) LenSquared() float32 {
	return a.Dot(a)
}

func (a Vec3) Len() float32 {
	return float32(math.Sqrt(float64(a.LenSquared())))
}

func (a Vec3) Normalize() (Vec3, bool) {
	if a.X == 0.0 && a.Y == 0.0 && a.Z == 0.0 {
		return Vec3{}, false
	}
	len := a.Len()
	return Vec3{a.X / len, a.Y / len, a.Z / len}, true
}
