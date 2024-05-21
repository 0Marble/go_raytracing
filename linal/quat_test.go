package linal

import (
	"math"
	"testing"
)

func quatAlmostEqual(a Quat, b Quat, t *testing.T) {
	if math.Abs(float64(a.X-b.X)) > 1e-3 ||
		math.Abs(float64(a.Y-b.Y)) > 1e-3 ||
		math.Abs(float64(a.Z-b.Z)) > 1e-3 ||
		math.Abs(float64(a.W-b.W)) > 1e-3 {
		t.Fatal(a, b)
	}
}

func TestQuatRotation1(t *testing.T) {
	q := QuatIdentity()
	mat, ok := q.ToMat()
	if !ok {
		t.Fatal()
	}
	x := Vec3{X: 1}
	y := mat.ApplyToPoint(x)
	vecAlmostEqual(y, x, t)
}
func TestQuatRotation2(t *testing.T) {
	q := QuatFromRot(Vec3{Z: 1}, math.Pi*0.5)
	mat, ok := q.ToMat()
	if !ok {
		t.Fatal()
	}
	x := Vec3{X: 1}
	y := mat.ApplyToPoint(x)
	vecAlmostEqual(y, Vec3{Y: 1}, t)
}
func TestQuatRotation3(t *testing.T) {
	q := QuatFromRot(Vec3{Z: 1}, math.Pi*0.25)
	mat, ok := q.ToMat()
	if !ok {
		t.Fatal()
	}
	x := Vec3{X: 1}
	y := mat.ApplyToPoint(x)
	vecAlmostEqual(y, Vec3{X: math.Sqrt2, Y: math.Sqrt2}.Div(2.0), t)
}
func TestQuatRotation4(t *testing.T) {
	q := QuatFromRot(Vec3{Z: 1}, math.Pi*0.25)
	q = QuatFromRot(Vec3{Y: 1}, math.Pi*0.5).Mul(q)
	mat, ok := q.ToMat()
	if !ok {
		t.Fatal()
	}
	x := Vec3{X: 1}
	y := mat.ApplyToPoint(x)
	vecAlmostEqual(y, Vec3{Z: -math.Sqrt2, Y: math.Sqrt2}.Div(2.0), t)
}

func TestQuatInterpolation(t *testing.T) {
	q := QuatIdentity()
	p := QuatFromRot(Vec3{Z: 1}, math.Pi*0.5)
	v := q.Slerp(p, 0.5)
	quatAlmostEqual(v, QuatFromRot(Vec3{Z: 1}, math.Pi*0.25), t)
}
