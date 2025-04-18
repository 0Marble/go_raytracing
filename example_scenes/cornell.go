package examplescenes

import (
	"math"
	"raytracing/lights"
	"raytracing/linal"
	"raytracing/materials"
	"raytracing/scene"
	"raytracing/shapes"
)

func CornellScene() (scene.Scene, linal.Transform) {
	red := materials.InitBlinnPhong(materials.Color{R: 1}, materials.Color{R: 0.2, G: 0.1, B: 0}, 1)
	green := materials.InitBlinnPhong(materials.Color{G: 1}, materials.Color{R: 0, G: 0.2, B: 0.1}, 10)
	blue := materials.InitBlinnPhong(materials.Color{B: 1}, materials.Color{R: 0.4, G: 0.6, B: 1}, 1000)
	white := materials.InitBlinnPhong(materials.Color{R: 1, G: 1, B: 1}, materials.Color{R: 0, G: 0, B: 0}, 0)
	outside := materials.Color{}

	left := shapes.InitRect(
		linal.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.QuatFromRot(linal.Vec3{Y: 1}, -math.Pi*0.5),
			Translation: linal.Vec3{X: -2.5},
		},
		&red,
	)
	right := shapes.InitRect(
		linal.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.QuatFromRot(linal.Vec3{Y: 1}, math.Pi*0.5),
			Translation: linal.Vec3{X: 2.5},
		},
		&green,
	)
	bottom := shapes.InitRect(
		linal.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.QuatFromRot(linal.Vec3{X: 1}, math.Pi*0.5),
			Translation: linal.Vec3{Y: -2.5},
		},
		&white,
	)
	top := shapes.InitRect(
		linal.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.QuatFromRot(linal.Vec3{X: 1}, -math.Pi*0.5),
			Translation: linal.Vec3{Y: 2.5},
		},
		&white,
	)
	back := shapes.InitRect(
		linal.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.QuatFromRot(linal.Vec3{Y: 1}, math.Pi),
			Translation: linal.Vec3{Z: -2.5},
		},
		&white,
	)
	front := shapes.InitRect(
		linal.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.QuatIdentity(),
			Translation: linal.Vec3{Z: 2.5},
		},
		&white,
	)
	ball := shapes.InitSphere(
		linal.Transform{
			Scale:       linal.Vec3{X: 0.5, Y: 0.5, Z: 0.5},
			Rotation:    linal.QuatIdentity(),
			Translation: linal.Vec3{X: -1.3, Y: -2.0, Z: 1.3},
		},
		&blue,
	)
	light := lights.InitSpotlight(linal.Vec3{Y: 2.49}, linal.Vec3{Y: -1}, materials.Color{R: 1, G: 1, B: 0.8})

	camTransform := linal.Transform{
		Scale:       linal.Vec3{X: 1, Y: 1, Z: 1},
		Rotation:    linal.QuatIdentity(),
		Translation: linal.Vec3{Z: -2},
	}
	s := scene.InitScene(
		[]shapes.Object{&top, &bottom, &left, &right, &back, &front, &ball},
		[]lights.Light{&light},
		&outside)

	return s, camTransform
}
