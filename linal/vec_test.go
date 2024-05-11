package linal

import (
	"math"
	"testing"
)

func vecAlmostEqual(a Vec3, b Vec3, t *testing.T) {
	if math.Abs(float64(a.X-b.X)) > 1e-3 || math.Abs(float64(a.Y-b.Y)) > 1e-3 || math.Abs(float64(a.Z-b.Z)) > 1e-3 {
		t.Fatal(a, b)
	}
}

func TestSpherical1(t *testing.T) {
	a := Vec3{1, 0, 0}
	asp := a.ToSpherical()
	b := asp.FromSpherical()
	vecAlmostEqual(asp, Vec3{1, math.Pi * 0.5, 0}, t)
	vecAlmostEqual(a, b, t)
}
func TestSpherical2(t *testing.T) {
	a := Vec3{1, 1, 0}
	asp := a.ToSpherical()
	b := asp.FromSpherical()
	vecAlmostEqual(asp, Vec3{math.Sqrt2, math.Pi * 0.5, 0.7854}, t)
	vecAlmostEqual(a, b, t)
}
func TestSpherical3(t *testing.T) {
	a := Vec3{1, 1, 1}
	asp := a.ToSpherical()
	b := asp.FromSpherical()
	vecAlmostEqual(asp, Vec3{1.732, 0.9553, 0.7854}, t)
	vecAlmostEqual(a, b, t)
}
