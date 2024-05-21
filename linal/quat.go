package linal

import "math"

type Quat struct {
	X float32
	Y float32
	Z float32
	W float32
}

func InitQuatFromImRe(im Vec3, re float32) Quat {
	return Quat{im.X, im.Y, im.Z, re}
}
func (a Quat) V() Vec3 {
	return Vec3{a.X, a.Y, a.Z}
}
func (a Quat) Add(b Quat) Quat {
	return Quat{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}
func (a Quat) Sub(b Quat) Quat {
	return Quat{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}
}
func (a Quat) Mul(b Quat) Quat {
	x := a.V()
	y := b.V()
	v := x.Cross(y).Add(x.Mul(b.W)).Add(y.Mul(a.W))
	u := a.W*b.W - x.Dot(y)
	return InitQuatFromImRe(v, u)
}
func (a Quat) Dot(b Quat) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}
func (a Quat) MulNum(t float32) Quat {
	return InitQuatFromImRe(a.V().Mul(t), a.W*t)
}
func (a Quat) Conjugate() Quat {
	return InitQuatFromImRe(a.V().Mul(-1), a.W)
}
func (a Quat) NormSquared() float32 {
	sum := a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W
	return sum
}
func (a Quat) Norm() float32 {
	return float32(math.Sqrt(float64(a.NormSquared())))
}
func QuatIdentity() Quat {
	return InitQuatFromImRe(Vec3{}, 1.0)
}
func (a Quat) Inverse() (Quat, bool) {
	ns := a.NormSquared()
	if ns == 0.0 {
		return Quat{}, false
	}

	return a.Conjugate().MulNum(1.0 / ns), true
}
func (a Quat) Normalize() (Quat, bool) {
	norm := a.Norm()
	if norm == 0.0 {
		return Quat{}, false
	}
	return a.MulNum(1.0 / norm), true
}

func (a Quat) ToMat() (Mat, bool) {
	ns := a.NormSquared()
	if ns == 0.0 {
		return Mat{}, false
	}

	s := 2.0 / ns
	mat := MatIdent(4)
	mat.Set(0, 0, 1.0-s*(a.Y*a.Y+a.Z*a.Z))
	mat.Set(1, 1, 1.0-s*(a.X*a.X+a.Z*a.Z))
	mat.Set(2, 2, 1.0-s*(a.X*a.X+a.Y*a.Y))
	mat.Set(0, 1, s*(a.X*a.Y-a.W*a.Z))
	mat.Set(0, 2, s*(a.X*a.Z+a.W*a.Y))
	mat.Set(1, 0, s*(a.X*a.Y+a.W*a.Z))
	mat.Set(1, 2, s*(a.Y*a.Z-a.W*a.X))
	mat.Set(2, 0, s*(a.X*a.Z-a.W*a.Y))
	mat.Set(2, 1, s*(a.Y*a.Z+a.W*a.X))

	return mat, true
}

func (a Quat) Slerp(b Quat, t float32) Quat {
	phi := float32(math.Acos(float64(a.Dot(b))))

	sinPhi := float32(math.Sin(float64(phi)))
	return a.MulNum(float32(math.Sin(float64(phi*(1-t)))) / sinPhi).Add(b.MulNum(float32(math.Sin(float64(phi*t))) / sinPhi))
}
func QuatFromRot(axis Vec3, angle float32) Quat {
	sin, cos := math.Sincos(float64(angle / 2.0))
	a := InitQuatFromImRe(axis.Mul(float32(sin)), float32(cos))
	return a
}
