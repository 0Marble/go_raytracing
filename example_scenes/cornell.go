package examplescenes

import (
	"math"
	"raytracing/camera"
	"raytracing/lights"
	"raytracing/linal"
	"raytracing/materials"
	"raytracing/scene"
	"raytracing/shapes"
)

func CornellScene() (scene.Scene, camera.Camera) {
	red := materials.InitSimpleMaterial(materials.Color{R: 1})
	green := materials.InitSimpleMaterial(materials.Color{G: 1})
	blue := materials.InitSimpleMaterial(materials.Color{B: 1})
	white := materials.InitSimpleMaterial(materials.Color{R: 1, G: 1, B: 1})
	outside := materials.InitSimpleMaterial(materials.Color{})

	left := shapes.InitRect(
		linal.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.Vec3{Y: -math.Pi * 0.5},
			Translation: linal.Vec3{X: -2.5},
		},
		&red,
	)
	right := shapes.InitRect(
		linal.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.Vec3{Y: math.Pi * 0.5},
			Translation: linal.Vec3{X: 2.5},
		},
		&green,
	)
	bottom := shapes.InitRect(
		linal.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.Vec3{X: math.Pi * 0.5},
			Translation: linal.Vec3{Y: -2.5},
		},
		&white,
	)
	top := shapes.InitRect(
		linal.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.Vec3{X: -math.Pi * 0.5},
			Translation: linal.Vec3{Y: 2.5},
		},
		&white,
	)
	back := shapes.InitRect(
		linal.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.Vec3{Y: math.Pi},
			Translation: linal.Vec3{Z: -2.5},
		},
		&white,
	)
	front := shapes.InitRect(
		linal.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Translation: linal.Vec3{Z: 2.5},
		},
		&white,
	)
	ball := shapes.InitSphere(
		linal.Transform{
			Scale: linal.Vec3{X: 0.75, Y: 0.75, Z: 0.75},
		},
		&blue,
	)
	light := lights.InitPointLight(linal.Vec3{Y: 2.49}, materials.Color{R: 1, G: 1, B: 0.8})

	cam := camera.InitSimpleCamera(
		linal.Transform{
			Scale:       linal.Vec3{X: 1, Y: 1, Z: 1},
			Translation: linal.Vec3{Z: -2},
		}, 500, 500)
	s := scene.InitScene(
		[]shapes.Object{&top, &bottom, &left, &right, &back, &front, &ball},
		[]lights.Light{&light},
		&outside)

	return s, &cam
}
