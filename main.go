package main

import (
	"fmt"
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
	camLens := lenses.InitProjectiveLens(camTransform)
	pp := postprocess.InitGammaCorrection(1.0/2.2, &postprocess.NoProcessing{})

	normalSensor := sensors.InitFsaaNByNSensor(1)
	fsaa2x2 := sensors.InitFsaaNByNSensor(2)
	fsaa4x4 := sensors.InitFsaaNByNSensor(4)
	rgss := sensors.InitFsaaRgssSensor()
	checker2x2 := sensors.InitFsaaCheckerboardSensor(2)
	checker4x4 := sensors.InitFsaaCheckerboardSensor(4)

	for i, sensor := range []camera.Sensor{&normalSensor, &fsaa2x2, &fsaa4x4, &rgss, &checker2x2, &checker4x4} {
		log.Println("Using sensor #", i)
		cam := camera.InitCamera(&camLens, sensor, &pp)
		rm := raytracing.InitSimpleRaytracer(s, 4)
		img := cam.Shoot(&rm, 500, 500, 0, 500, 0, 500)

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
