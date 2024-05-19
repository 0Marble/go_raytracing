package main

import (
	"log"
	"math"
	"raytracing/linal"
	"raytracing/materials"
	"raytracing/scene"
	"raytracing/shapes"
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
	m1 := materials.InitSimpleMaterial(materials.Color{R: 1}, 1.0)
	m2 := materials.InitSimpleMaterial(materials.Color{R: 1, G: 1, B: 1}, 1.0)
	sky := materials.InitSimpleMaterial(materials.Color{}, 1)
	s1 := shapes.InitSphere(linal.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}, Translation: linal.Vec3{Z: 3}}, &m1)
	s2 := shapes.InitSphere(linal.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}, Translation: linal.Vec3{Z: -3}}, &m2)
	cam := scene.InitSimpleCamera(linal.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}}, 400, 400)
	s := scene.InitScene([]shapes.Object{&s1, &s2}, &cam, &sky)
	rm := InitRaymarcher(s, 1)

	ray := linal.Ray{Dir: linal.Vec3{Z: 1}}
	step := rm.march(ray)

	if !step.isHit {
		t.Fatal(step)
	}
	if step.material != s1.Material() {
		t.Fatal(step)
	}
	vecAlmostEqual(step.nextRay.Dir, ray.Dir.Mul(-1), t)

	ray = step.nextRay
	step = rm.march(ray)
	if !step.isHit {
		t.Fatal(step)
	}
	if step.material != s2.Material() {
		t.Fatal(step)
	}
	vecAlmostEqual(step.nextRay.Dir, ray.Dir.Mul(-1), t)

	ray = step.nextRay
	step = rm.march(ray)
	if !step.isHit {
		t.Fatal(step)
	}
	if step.material != s1.Material() {
		t.Fatal(step)
	}
	vecAlmostEqual(step.nextRay.Dir, ray.Dir.Mul(-1), t)
}

func TestLowGround(t *testing.T) {
	log.Println("TestLowGround")
	purple := materials.InitSimpleMaterial(materials.Color{R: 0.5, G: 0, B: 0.5}, 1.0)
	sky := materials.InitSimpleMaterial(materials.Color{R: 1, G: 1, B: 1}, 0)
	ground := shapes.InitSphere(
		linal.Transform{
			Scale:       linal.Vec3{X: 100, Y: 2, Z: 100},
			Translation: linal.Vec3{Y: -8},
		},
		&purple,
	)
	cam := scene.InitSimpleCamera(linal.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}, Translation: linal.Vec3{Z: -2}}, 500, 500)
	s := scene.InitScene([]shapes.Object{&ground}, &cam, &sky)
	rm := InitRaymarcher(s, 10)

	dir := linal.Vec3{Y: -1, Z: 1}
	dir, _ = dir.Normalize()
	ray := linal.Ray{Dir: dir}

	step := rm.march(ray)
	if !step.isHit || step.material != ground.Material() {
		t.Fatal(step)
	}
	log.Println("Next ray: ", step.nextRay)
	step = rm.march(step.nextRay)
	if !step.isHit || step.material != &rm.scene.Outside {
		t.Fatal(step)
	}

}
