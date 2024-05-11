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

	red := scene.InitMaterial(scene.Color{R: 1, G: 0.2, B: 0.2}, 0.9, false)
	green := scene.InitMaterial(scene.Color{R: 0, G: 0.8, B: 0.2}, 0.2, false)
	purple := scene.InitMaterial(scene.Color{R: 0.5, G: 0, B: 0.5}, 0.1, false)
	sky := scene.InitMaterial(scene.Color{R: 1, G: 1, B: 1}, 0, true)
	ball1 := shapes.InitSphere(transfrom.Transform{Scale: linal.Vec3{X: 2, Y: 1, Z: 1}, Translation: linal.Vec3{Z: 3, X: -3}}, &red)
	ball2 := shapes.InitSphere(transfrom.Transform{Scale: linal.Vec3{X: 1, Y: 2, Z: 1}, Translation: linal.Vec3{Z: 3, X: 3}}, &green)
	ground := shapes.InitSphere(transfrom.Transform{Scale: linal.Vec3{X: 100, Y: 1, Z: 100}, Translation: linal.Vec3{Y: -2}}, &purple)
	cam := scene.InitSimpleCamera(transfrom.Transform{Scale: linal.Vec3{X: 1, Y: 1, Z: 1}}, 400, 400)
	s := scene.InitScene([]scene.Object{&ball1, &ball2, &ground}, &cam, &sky)
	rm := raymarching.InitRaymarcher(s, 2)

	for i := 0; i < 60; i++ {
		log.Println(i)
		for rm.March() {
		}
		img := cam.ToImage(200, 200)
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
