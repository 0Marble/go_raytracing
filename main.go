package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	scenes "raytracing/example_scenes"
	"raytracing/raytracing"
)

func main() {
	log.SetFlags(log.Lshortfile)

	s := scenes.CornellScene()
	rm := raytracing.InitRaytracer(s, 10)

	for i := 0; i < 60; i++ {
		log.Println(i)
		for rm.Trace() {
		}
		img := s.Cam.ToImage(500, 500)
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
