package main

import (
	"fmt"
	"image/png"
	"log"
	"math"
	"os"
	"raytracing/linal"
	"raytracing/raymarching"
	"raytracing/scene"
	"raytracing/scene/shapes"
	"raytracing/transfrom"
)

func main() {
	log.SetFlags(log.Lshortfile)

	red := scene.InitSimpleMaterial(scene.Color{R: 1}, 0.0, false)
	green := scene.InitSimpleMaterial(scene.Color{G: 1}, 0.0, false)
	blue := scene.InitSimpleMaterial(scene.Color{B: 1}, 1.0, false)
	white := scene.InitSimpleMaterial(scene.Color{R: 1, G: 1, B: 1}, 0.0, false)
	light := scene.InitSimpleMaterial(scene.Color{R: 1, G: 1, B: 0.8}, 0.0, true)
	outside := scene.InitSimpleMaterial(scene.Color{}, 0, false)

	left := shapes.InitRect(
		transfrom.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.Vec3{Y: -math.Pi * 0.5},
			Translation: linal.Vec3{X: -2.5},
		},
		&red,
	)
	right := shapes.InitRect(
		transfrom.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.Vec3{Y: math.Pi * 0.5},
			Translation: linal.Vec3{X: 2.5},
		},
		&green,
	)
	bottom := shapes.InitRect(
		transfrom.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.Vec3{X: math.Pi * 0.5},
			Translation: linal.Vec3{Y: -2.5},
		},
		&white,
	)
	top := shapes.InitRect(
		transfrom.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.Vec3{X: math.Pi * 0.5},
			Translation: linal.Vec3{Y: 2.5},
		},
		&white,
	)
	lightSource := shapes.InitRect(
		transfrom.Transform{
			Scale:       linal.Vec3{X: 1.5, Y: 1.5, Z: 1},
			Rotation:    linal.Vec3{X: math.Pi * 0.5},
			Translation: linal.Vec3{Y: 2.49},
		},
		&light,
	)
	back := shapes.InitRect(
		transfrom.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Rotation:    linal.Vec3{Y: math.Pi},
			Translation: linal.Vec3{Z: -2.5},
		},
		&white,
	)
	front := shapes.InitRect(
		transfrom.Transform{
			Scale:       linal.Vec3{X: 5, Y: 5, Z: 1},
			Translation: linal.Vec3{Z: 2.5},
		},
		&white,
	)
	ball := shapes.InitSphere(
		transfrom.Transform{
			Scale: linal.Vec3{X: 0.75, Y: 0.75, Z: 0.75},
		},
		&blue,
	)

	cam := scene.InitSimpleCamera(
		transfrom.Transform{
			Scale:       linal.Vec3{X: 1, Y: 1, Z: 1},
			Translation: linal.Vec3{Z: -2},
		}, 500, 500)
	s := scene.InitScene([]scene.Object{&top, &bottom, &left, &right, &back, &front, &ball, &lightSource}, &cam, &outside)
	rm := raymarching.InitRaymarcher(s, 10)

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
