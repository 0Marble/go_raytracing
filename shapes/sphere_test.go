package shapes

import (
	"math"
	"raytracing/linal"
	"raytracing/materials"
	"testing"
)

func almostEqual(x float32, y float32, t *testing.T) {
	if math.Abs(float64(x-y)) > 1e-3 {
		t.Fatal(x, y)
	}
}
func vecAlmostEqual(a linal.Vec3, b linal.Vec3, t *testing.T) {
	if math.Abs(float64(a.X-b.X)) > 1e-3 || math.Abs(float64(a.Y-b.Y)) > 1e-3 || math.Abs(float64(a.Z-b.Z)) > 1e-3 {
		t.Fatal(a, b)
	}
}

func TestSphereIntersect(t *testing.T) {
	m := materials.InitSimpleMaterial(materials.Color{})
	s := InitSphere(linal.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}, Rotation: linal.QuatIdentity()}, &m)
	ray := linal.Ray{Dir: linal.Vec3{Z: 1}, Start: linal.Vec3{Z: -2}}
	intersection := s.Intersect(ray)

	if !intersection.IsHit {
		t.Fatal(intersection)
	}
	pt := s.FromUv(intersection.Uv)
	vecAlmostEqual(pt, linal.Vec3{Z: -1}, t)
}

func TestSphereUvs(t *testing.T) {
	m := materials.InitSimpleMaterial(materials.Color{})
	s := InitSphere(linal.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}, Rotation: linal.QuatIdentity()}, &m)

	pt := linal.Vec3{X: 1}
	uv := s.ToUv(pt)
	fromUv := s.FromUv(uv)
	vecAlmostEqual(pt, fromUv, t)

	pt = linal.Vec3{Z: 1}
	uv = s.ToUv(pt)
	fromUv = s.FromUv(uv)
	vecAlmostEqual(pt, fromUv, t)
}

func TestSphereNormal1(t *testing.T) {
	m := materials.InitSimpleMaterial(materials.Color{})
	s := InitSphere(linal.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}, Rotation: linal.QuatIdentity()}, &m)
	pt := linal.Vec3{Z: 1}
	uv := s.ToUv(pt)
	n := s.Normal(uv)
	pt, _ = pt.Normalize()
	vecAlmostEqual(n, pt, t)

	pt = linal.Vec3{Z: -1}
	uv = s.ToUv(pt)
	n = s.Normal(uv)
	pt, _ = pt.Normalize()
	vecAlmostEqual(n, pt, t)

	pt = linal.Vec3{Y: 1}
	uv = s.ToUv(pt)
	n = s.Normal(uv)
	pt, _ = pt.Normalize()
	vecAlmostEqual(n, pt, t)

	pt = linal.Vec3{X: math.Sqrt2, Z: math.Sqrt2}
	uv = s.ToUv(pt)
	n = s.Normal(uv)
	pt, _ = pt.Normalize()
	vecAlmostEqual(n, pt, t)
}
