package main

import (
	"image/png"
	"log"
	"os"
	"raytracing/camera"
	"raytracing/camera/lenses"
	"raytracing/camera/postprocess"
	"raytracing/camera/sensors"
	scenes "raytracing/example_scenes"
	"raytracing/raytracing"
)

func main() {
	log.SetFlags(log.Lshortfile)

	s, camTransform := scenes.CornellScene()
	camSensor := sensors.InitFsaaCheckerboardSensor(4)
	camLens := lenses.InitSphericalLens(camTransform)
	cam := camera.InitCamera(&camLens, &camSensor, &postprocess.NoProcessing{})
	rm := raytracing.InitSimpleRaytracer(s, 2)
	img := cam.Shoot(&rm, 500, 500, 0, 500, 0, 500)

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
