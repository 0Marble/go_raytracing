package camera

import (
	"image"
	"image/color"
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

func (c *Camera) Shoot(rt raytracing.Raytracer, width, height, left, right, bottom, top int) Image {
	w := right - left
	h := top - bottom
	img := Image{w, h, make([]materials.Color, w*h)}

	for y := bottom; y < top; y++ {
		for x := left; x < right; x++ {
			color := c.sensor.GetPixel(rt, c.lens, x, y, width, height)
			img.Pixels[(y-bottom)*w+(x-left)] = color
		}
	}

	return c.pp.Process(img)
}

type Lens interface {
	ShootRay(uv linal.Uv) linal.Ray
}

type Sensor interface {
	GetPixel(rt raytracing.Raytracer, lens Lens, x, y, width, height int) materials.Color
}

type PostProcess interface {
	Process(img Image) Image
}
