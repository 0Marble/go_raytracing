package shapes

import (
	"math"
	"raytracing/linal"
	"raytracing/scene"
	"raytracing/transfrom"
	"testing"
)

func uvAlmostEqual(a scene.Uv, b scene.Uv, t *testing.T) {
	if math.Abs(float64(a.U-b.U)) > 1e-3 || math.Abs(float64(a.V-b.V)) > 1e-3 {
		t.Fatal(a, b)
	}
}

func TestRectIntersect1(t *testing.T) {
	m := scene.InitSimpleMaterial(scene.Color{}, 0.0, false)
	rect := InitRect(transfrom.Transform{
		Scale: linal.Vec3{X: 1, Y: 1, Z: 1}},
		&m)

	ray := scene.Ray{Start: linal.Vec3{Z: -1}, Dir: linal.Vec3{Z: 1}}
	intersection := rect.Intersect(ray)

	if !intersection.IsHit {
		t.Fatal(intersection)
	}
	uvAlmostEqual(intersection.Uv, scene.Uv{U: 0.5, V: 0.5}, t)

	pt := rect.FromUv(intersection.Uv)
	vecAlmostEqual(pt, linal.Vec3{}, t)

	n := rect.Normal(intersection.Uv)
	vecAlmostEqual(n, linal.Vec3{Z: -1}, t)
}

func TestRectIntersect2(t *testing.T) {
	m := scene.InitSimpleMaterial(scene.Color{}, 0.0, false)
	rect := InitRect(transfrom.Transform{
		Scale:    linal.Vec3{X: 2, Y: 2, Z: 1},
		Rotation: linal.Vec3{X: math.Pi * 0.25},
	}, &m)

	ray := scene.Ray{Start: linal.Vec3{Z: -1}, Dir: linal.Vec3{Z: 1}}
	intersection := rect.Intersect(ray)

	if !intersection.IsHit {
		t.Fatal(intersection)
	}
	uvAlmostEqual(intersection.Uv, scene.Uv{U: 0.5, V: 0.5}, t)

	pt := rect.FromUv(intersection.Uv)
	vecAlmostEqual(pt, linal.Vec3{}, t)

	n := rect.Normal(intersection.Uv)
	vecAlmostEqual(n, linal.Vec3{Y: math.Sqrt2 * 0.5, Z: -math.Sqrt2 * 0.5}, t)
}
