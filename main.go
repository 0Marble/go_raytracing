package main

import (
	"fmt"
	"image/png"
	"log"
	"math"
	"os"
	"raytracing/linal"
	"raytracing/materials"
	"raytracing/scene"
	"raytracing/shapes"
)

func main() {
	log.SetFlags(log.Lshortfile)

	red := materials.InitSimpleMaterial(materials.Color{R: 1}, 0.0)
	green := materials.InitSimpleMaterial(materials.Color{G: 1}, 0.0)
	blue := materials.InitSimpleMaterial(materials.Color{B: 1}, 1.0)
	white := materials.InitSimpleMaterial(materials.Color{R: 1, G: 1, B: 1}, 0.0)
	light := materials.InitSimpleMaterial(materials.Color{R: 1, G: 1, B: 0.8}, 0.0)
	outside := materials.InitSimpleMaterial(materials.Color{}, 0)

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
			Rotation:    linal.Vec3{X: math.Pi * 0.5},
			Translation: linal.Vec3{Y: 2.5},
		},
		&white,
	)
	lightSource := shapes.InitRect(
		linal.Transform{
			Scale:       linal.Vec3{X: 1.5, Y: 1.5, Z: 1},
			Rotation:    linal.Vec3{X: math.Pi * 0.5},
			Translation: linal.Vec3{Y: 2.49},
		},
		&light,
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

	cam := scene.InitSimpleCamera(
		linal.Transform{
			Scale:       linal.Vec3{X: 1, Y: 1, Z: 1},
			Translation: linal.Vec3{Z: -2},
		}, 500, 500)
	s := scene.InitScene([]shapes.Object{&top, &bottom, &left, &right, &back, &front, &ball, &lightSource}, &cam, &outside)
	rm := InitRaymarcher(s, 10)

	for i := 0; i < 60; i++ {
		log.Println(i)
		for rm.March() {
		}
		img := cam.ToImage(500, 500)
		fileName := fmt.Sprintf("images/img_%v.png", i)
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal("Could not open ", fileName, " : ", err)
		}

		err = png.Encode(file, &img)
		if err != nil {
			log.Fatal("Could not save as png: ", fileName, " : ", err)
		}
		file.Close()
	}
}
