package raymarching

import (
	"log"
	"math"
	"raytracing/linal"
	"raytracing/scene"
	"raytracing/scene/shapes"
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

func TestTwoSpheres(t *testing.T) {
	m1 := scene.InitMaterial(scene.Color{R: 1}, 1.0, false)
	m2 := scene.InitMaterial(scene.Color{R: 1, G: 1, B: 1}, 1.0, true)
	sky := scene.InitMaterial(scene.Color{}, 1, false)
	s1 := shapes.InitSphere(transfrom.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}, Translation: linal.Vec3{Z: 3}}, &m1)
	s2 := shapes.InitSphere(transfrom.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}, Translation: linal.Vec3{Z: -3}}, &m2)
	cam := scene.InitSimpleCamera(transfrom.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}}, 400, 400)
	s := scene.InitScene([]scene.Object{&s1, &s2}, &cam, &sky)
	rm := InitRaymarcher(s, 1)

	ray := scene.Ray{Dir: linal.Vec3{Z: 1}, Step: 10}
	step := rm.march(ray, 1)

	if !step.isHit {
		t.Fatal(step)
	}
	if step.material != s1.Material() {
		t.Fatal(step)
	}
	if step.bounce != 0 {
		t.Fatal(step)
	}
	vecAlmostEqual(step.nextRay.Dir, ray.Dir.Mul(-1), t)

	ray = step.nextRay
	log.Println("Next ray: ", ray)
	step = rm.march(ray, 0)
	if !step.isHit {
		t.Fatal(step)
	}
	if step.material != s2.Material() {
		t.Fatal(step)
	}
	if step.bounce != -1 {
		t.Fatal(step)
	}
	vecAlmostEqual(step.nextRay.Dir, ray.Dir.Mul(-1), t)
}
