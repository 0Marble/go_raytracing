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

func (a Vec3) Min(b Vec3) Vec3 {
	if a.X > b.X {
		a.X = b.X
	}
	if a.Y > b.Y {
		a.Y = b.Y
	}
	if a.Z > b.Z {
		a.Z = b.Z
	}
	return a
}
func (a Vec3) Max(b Vec3) Vec3 {
	if a.X < b.X {
		a.X = b.X
	}
	if a.Y < b.Y {
		a.Y = b.Y
	}
	if a.Z < b.Z {
		a.Z = b.Z
	}
	return a
}

func (c Vec3) Clamp(min Vec3, max Vec3) Vec3 {
	res := c
	if res.X < min.X {
		res.X = min.X
	}
	if res.X > max.X {
		res.X = max.X
	}
	if res.Y < min.Y {
		res.Y = min.Y
	}
	if res.Y > max.Y {
		res.Y = max.Y
	}
	if res.Z < min.Z {
		res.Z = min.Z
	}
	if res.Z > max.Z {
		res.Z = max.Z
	}

	return res
}

func (sp Vec3) FromSpherical() Vec3 {
	theta := sp.Y
	phi := sp.Z
	thetaSin, thetaCos := math.Sincos(float64(theta))
	phiSin, phiCos := math.Sincos(float64(phi))
	res := Vec3{X: sp.X * float32(thetaSin*phiCos), Y: sp.X * float32(thetaSin*phiSin), Z: sp.X * float32(thetaCos)}
	return res
}

func (pt Vec3) ToSpherical() Vec3 {
	xy := pt.X*pt.X + pt.Y*pt.Y

	theta := float32(math.Atan2(math.Sqrt(float64(xy)), float64(pt.Z)))
	phi := float32(math.Atan2(float64(pt.Y), float64(pt.X)))

	return Vec3{X: pt.Len(), Y: theta, Z: phi}
}
