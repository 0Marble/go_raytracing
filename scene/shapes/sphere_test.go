package shapes

import (
	"math"
	"raytracing/linal"
	"raytracing/scene"
	"raytracing/transfrom"
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

func TestSphereDist1(t *testing.T) {
	m := scene.InitMaterial(scene.Color{}, 0.0, false)
	s := InitSphere(transfrom.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}}, &m)
	pt := linal.Vec3{X: 0, Y: 0, Z: 10}
	almostEqual(s.Distance(pt), 9, t)

	pt = linal.Vec3{X: 10, Y: 0, Z: 10}
	almostEqual(s.Distance(pt), 10*math.Sqrt2-1, t)

	pt = linal.Vec3{X: 10, Y: 10, Z: 10}
	almostEqual(s.Distance(pt), pt.Len()-1, t)
}
func TestSphereDist2(t *testing.T) {
	m := scene.InitMaterial(scene.Color{}, 0.0, false)
	s := InitSphere(transfrom.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 2}}, &m)
	pt := linal.Vec3{X: 0, Y: 0, Z: 10}
	almostEqual(s.Distance(pt), 8, t)

	pt = linal.Vec3{X: 10, Y: 0, Z: 0}
	almostEqual(s.Distance(pt), 9, t)
}
func TestSphereDist3(t *testing.T) {
	m := scene.InitMaterial(scene.Color{}, 0.0, false)
	s := InitSphere(transfrom.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}, Translation: linal.Vec3{Z: -1}}, &m)
	pt := linal.Vec3{X: 0, Y: 0, Z: 10}
	almostEqual(s.Distance(pt), 10, t)
}
func TestSphereDist4(t *testing.T) {
	m := scene.InitMaterial(scene.Color{}, 0.0, false)
	s := InitSphere(transfrom.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}, Translation: linal.Vec3{Z: 5}}, &m)
	pt := linal.Vec3{X: 0, Y: 0, Z: 3}
	almostEqual(s.Distance(pt), 1, t)
	pt = linal.Vec3{X: 0, Y: 0, Z: 4}
	almostEqual(s.Distance(pt), 0, t)
	pt = linal.Vec3{X: 0, Y: 0, Z: 5}
	almostEqual(s.Distance(pt), -1, t)
	pt = linal.Vec3{X: 0, Y: 0, Z: 6}
	almostEqual(s.Distance(pt), 0, t)
	pt = linal.Vec3{X: 0, Y: 0, Z: 7}
	almostEqual(s.Distance(pt), 1, t)
}

func TestSphereUvs(t *testing.T) {
	m := scene.InitMaterial(scene.Color{}, 0.0, false)
	s := InitSphere(transfrom.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}}, &m)

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
	m := scene.InitMaterial(scene.Color{}, 0.0, false)
	s := InitSphere(transfrom.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}}, &m)
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
