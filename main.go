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
	"raytracing/scene"
)

func RenderAndSave(s scene.Scene, cam camera.Camera, fileName string) {
	log.Println("Saving to ", fileName)
	rm := raytracing.InitSimpleRaytracer(s, 4)
	img := cam.Shoot(&rm, 500, 500, 0, 500, 0, 500)

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

func main() {
	log.SetFlags(log.Lshortfile)

	s, camTransform := scenes.CornellScene()
	camLens := lenses.InitProjectiveLens(camTransform)
	pp := postprocess.InitGammaCorrection(1.0/2.2, &postprocess.NoProcessing{})
	sensor := sensors.InitFsaaRgssSensor()
	RenderAndSave(s, camera.InitCamera(&camLens, &sensor, &pp), "images/img.png")
}
