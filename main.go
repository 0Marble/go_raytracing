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
	"raytracing/materials"
	"raytracing/raytracing"
	"raytracing/scene"
)

func RenderAndSave(s scene.Scene, cam camera.Camera, fileName string) {
	w := 500
	h := 500
	n := 10
	tileWidth := w / n
	tileHeight := h / n
	tiles := make([]chan camera.Image, n*n)
	img := camera.Image{Width: w, Height: h, Pixels: make([]materials.Color, w*h)}

	log.Println("Rendering ", fileName)
	rm := raytracing.InitSimpleRaytracer(s, 4)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c := make(chan camera.Image)
			left := i * tileWidth
			right := (i + 1) * tileWidth
			bottom := j * tileHeight
			top := (j + 1) * tileHeight
			go func() {
				img := cam.Shoot(&rm, w, h, left, right, bottom, top)
				c <- img
			}()
			tiles[i*n+j] = c
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c := tiles[i*n+j]
			tile := <-c

			left := i * tileWidth
			right := (i + 1) * tileWidth
			bottom := j * tileHeight
			top := (j + 1) * tileHeight
			for y := bottom; y < top; y++ {
				for x := left; x < right; x++ {
					img.Pixels[y*w+x] = tile.Pixels[(y-bottom)*tileWidth+(x-left)]
				}
			}
		}
	}

	log.Println("Saving to ", fileName)

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
	sensor := sensors.InitFsaaCheckerboardSensor(4)
	RenderAndSave(s, camera.InitCamera(&camLens, &sensor, &pp), "images/img.png")
}
