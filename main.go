package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"raytracing/linal"
	"raytracing/raymarching"
	"raytracing/scene"
	"raytracing/scene/shapes"
	"raytracing/transfrom"
)

func main() {
	log.SetFlags(log.Lshortfile)

	red := scene.InitSimpleMaterial(scene.Color{R: 1, G: 0.2, B: 0.2}, 1.0, false)
	green := scene.InitSimpleMaterial(scene.Color{R: 0, G: 0.8, B: 0.2}, 0.9, false)
	purple := scene.InitSimpleMaterial(scene.Color{R: 0.5, G: 0.0, B: 0.5}, 0.0, false)
	sky := scene.InitSimpleMaterial(scene.Color{R: 0.4, G: 0.7, B: 0.9}, 1, true)
	light := scene.InitSimpleMaterial(scene.Color{R: 1, G: 1, B: 1}, 0, true)
	ball1 := shapes.InitSphere(
		transfrom.Transform{
			Scale:       linal.Vec3{X: 1, Y: 0.4, Z: 1},
			Translation: linal.Vec3{Z: 1, X: -1.5},
		},
		&red,
	)
	ball2 := shapes.InitSphere(
		transfrom.Transform{
			Scale:       linal.Vec3{X: 1, Y: 2, Z: 1},
			Translation: linal.Vec3{Z: 1, X: 1.5},
		},
		&green,
	)
	ground := shapes.InitSphere(
		transfrom.Transform{
			Scale:       linal.Vec3{X: 100, Y: 2, Z: 100},
			Translation: linal.Vec3{Y: -4},
		},
		&purple,
	)
	sun := shapes.InitSphere(
		transfrom.Transform{
			Scale:       linal.Vec3{X: 10, Y: 10, Z: 1},
			Translation: linal.Vec3{Z: -10},
		},
		&light,
	)
	cam := scene.InitSimpleCamera(
		transfrom.Transform{
			Scale: linal.Vec3{X: 1, Y: 1, Z: 1},
		}, 500, 500)
	s := scene.InitScene([]scene.Object{&ground, &ball1, &ball2, &sun}, &cam, &sky)
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
