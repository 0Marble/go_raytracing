package examplescenes

import (
	"raytracing/lights"
	"raytracing/linal"
	"raytracing/materials"
	"raytracing/scene"
	"raytracing/shapes"
)

func RedRectScene() (scene.Scene, linal.Transform) {
	bg := materials.InitSimpleMaterial(materials.Color{})
	m := materials.InitSimpleMaterial(materials.Color{R: 1})
	r := shapes.InitRect(
		linal.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1},
			Rotation: linal.QuatIdentity(),
		}.ToStaticTransform(), &m)
	light := lights.InitDirectionalLight(linal.Vec3{Z: 1}, materials.Color{R: 1, G: 1, B: 1})
	scene := scene.InitScene([]shapes.Object{&r}, []lights.Light{&light}, &bg)

	camTransform := linal.Transform{
		Scale:       linal.Vec3{X: 1, Y: 1, Z: 1},
		Translation: linal.Vec3{Z: -1},
		Rotation:    linal.QuatIdentity(),
	}
	return scene, camTransform
}
