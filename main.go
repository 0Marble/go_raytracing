package main

import (
	"image/png"
	"log"
	"os"
	scenes "raytracing/example_scenes"
	"raytracing/raytracing"
)

func main() {
	log.SetFlags(log.Lshortfile)

	s, cam := scenes.CornellScene()
	rm := raytracing.InitSimpleRaytracer(s, 10)
	img := cam.Shoot(&rm, 500, 500, 0, 500, 500, 0)

	fileName := "images/img.png"
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
