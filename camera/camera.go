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

type Sample struct {
	uv    linal.Uv
	color materials.Color
}
type Camera interface {
	Shoot(rt raytracing.Raytracer) Image
}
