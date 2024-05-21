package camera

import (
	"image"
	"image/color"
	"log"
	"math"
	"raytracing/linal"
	"raytracing/materials"
	"raytracing/raytracing"
)

type Image struct {
	Width  int
	Height int
	Pixels []materials.Color
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{i.Width, i.Height}}
}

func (i *Image) At(x int, y int) color.Color {
	return &i.Pixels[(i.Height-y-1)*i.Width+x]
}

type Camera struct {
	lens   Lens
	sensor Sensor
	pp     PostProcess
}

func InitCamera(lens Lens, sensor Sensor, postprocess PostProcess) Camera {
	return Camera{lens, sensor, postprocess}
}

func (c *Camera) Shoot(rt raytracing.Raytracer, width, height, left, right, bottom, top int, timeSteps int) Image {
	w := right - left
	h := top - bottom
	img := Image{w, h, make([]materials.Color, w*h)}
	step := 1.0 / float32(timeSteps)

	total := w * h * timeSteps
	cnt := 0
	prevPercent := 0
	for y := bottom; y < top; y++ {
		for x := left; x < right; x++ {
			color := materials.Color{}
			for t := 0; t < timeSteps; t++ {
				w := float32(math.Pow(0.5, float64(timeSteps-t)))
				color = c.sensor.GetPixel(rt, c.lens, x, y, width, height, float32(t)*step).MulNum(w).Add(color)
				cnt++
			}
			color = color.Clamp(materials.Color{}, materials.Color{R: 1, G: 1, B: 1})
			img.Pixels[(y-bottom)*w+(x-left)] = color

			percent := int(float32(cnt) / float32(total) * 100.0)
			if percent%10 == 0 && percent != prevPercent {
				log.Println(percent, "% done")
				prevPercent = percent
			}
		}
	}

	return c.pp.Process(img)
}

type Lens interface {
	ShootRay(uv linal.Uv) linal.Ray
}

type Sensor interface {
	GetPixel(rt raytracing.Raytracer, lens Lens, x, y, width, height int, time float32) materials.Color
}

type PostProcess interface {
	Process(img Image) Image
}
