package raytracing

import (
	"log"
	"math"
	"raytracing/lights"
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
func colorAlmostEqual(a materials.Color, b materials.Color, t *testing.T) {
	if math.Abs(float64(a.R-b.R)) > 1e-3 || math.Abs(float64(a.G-b.G)) > 1e-3 || math.Abs(float64(a.B-b.B)) > 1e-3 {
		t.Fatal(a, b)
	}
}

func TestTwoSpheres(t *testing.T) {
	m1 := materials.InitSimpleMaterial(materials.Color{R: 1})
	m2 := materials.InitSimpleMaterial(materials.Color{R: 1, G: 1, B: 1})
	sky := materials.InitSimpleMaterial(materials.Color{})
	s1 := shapes.InitSphere(
		linal.Transform{
			Scale:       linal.Vec3{X: 1, Y: 1, Z: 1},
			Translation: linal.Vec3{Z: 3},
			Rotation:    linal.QuatIdentity(),
		}, &m1)
	s2 := shapes.InitSphere(
		linal.Transform{
			Scale:       linal.Vec3{X: 1, Y: 1, Z: 1},
			Translation: linal.Vec3{Z: -3},
			Rotation:    linal.QuatIdentity(),
		}, &m2)
	s := scene.InitScene([]shapes.Object{&s1, &s2}, []lights.Light{}, &sky)
	rm := InitSimpleRaytracer(s, 1)

	ray := linal.Ray{Dir: linal.Vec3{Z: 1}}
	step := rm.trace(ray)

	if step.material != s1.Material() {
		t.Fatal(step)
	}
	vecAlmostEqual(step.nextRay.Dir, ray.Dir.Mul(-1), t)

	ray = step.nextRay
	step = rm.trace(ray)
	if step.material != s2.Material() {
		t.Fatal(step)
	}
	vecAlmostEqual(step.nextRay.Dir, ray.Dir.Mul(-1), t)

	ray = step.nextRay
	step = rm.trace(ray)
	if step.material != s1.Material() {
		t.Fatal(step)
	}
	vecAlmostEqual(step.nextRay.Dir, ray.Dir.Mul(-1), t)
}

func TestLowGround(t *testing.T) {
	log.Println("TestLowGround")
	purple := materials.InitSimpleMaterial(materials.Color{R: 0.5, G: 0, B: 0.5})
	sky := materials.InitSimpleMaterial(materials.Color{R: 1, G: 1, B: 1})
	ground := shapes.InitSphere(
		linal.Transform{
			Scale:       linal.Vec3{X: 100, Y: 2, Z: 100},
			Translation: linal.Vec3{Y: -8},
			Rotation:    linal.QuatIdentity(),
		},
		&purple,
	)
	s := scene.InitScene([]shapes.Object{&ground}, []lights.Light{}, &sky)
	rm := InitSimpleRaytracer(s, 10)

	dir := linal.Vec3{Y: -1, Z: 1}
	dir, _ = dir.Normalize()
	ray := linal.Ray{Dir: dir}

	step := rm.trace(ray)
	if step.material != ground.Material() {
		t.Fatal(step)
	}
	log.Println("Next ray: ", step.nextRay)
	step = rm.trace(step.nextRay)
	if step.material != &rm.scene.Outside {
		t.Fatal(step)
	}

}

func TestLight(t *testing.T) {
	bg := materials.InitSimpleMaterial(materials.Color{})
	m := materials.InitSimpleMaterial(materials.Color{R: 1})
	r := shapes.InitRect(linal.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}, Rotation: linal.QuatIdentity()}, &m)
	light := lights.InitDirectionalLight(linal.Vec3{Z: 1}, materials.Color{R: 1, G: 1, B: 1})
	scene := scene.InitScene([]shapes.Object{&r}, []lights.Light{&light}, &bg)
	raytracer := InitSimpleRaytracer(scene, 1)

	n := 50
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			u := float32(i)/float32(n) - 0.5
			v := float32(j)/float32(n) - 0.5
			dir, _ := linal.Vec3{X: u, Y: v, Z: 1}.Normalize()
			color := raytracer.Sample(linal.Ray{Start: linal.Vec3{Z: -1}, Dir: dir})
			colorAlmostEqual(color, materials.Color{R: 1}, t)
		}
	}
}
