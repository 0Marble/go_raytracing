package scene

import (
	"image"
	"image/color"
	"math/rand"
	"raytracing/linal"
	"raytracing/materials"
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

type Camera interface {
	Transform() *linal.Transform
	ShootRay() (linal.Ray, linal.Uv, bool)
	EmitPixel(uv linal.Uv, color materials.Color)
	ToImage(width int, height int) Image
}

type SimpleCamera struct {
	transform linal.Transform
	x         int
	y         int
	xSamples  int
	ySamples  int
	samples   []struct {
		color materials.Color
		uv    linal.Uv
	}
}

func InitSimpleCamera(transform linal.Transform, xSamples int, ySamples int) SimpleCamera {
	return SimpleCamera{transform: transform, x: 0, y: 0, xSamples: xSamples, ySamples: ySamples, samples: make([]struct {
		color materials.Color
		uv    linal.Uv
	}, 0)}
}

func (c *SimpleCamera) Transform() *linal.Transform {
	return &c.transform
}

func (c *SimpleCamera) ShootRay() (linal.Ray, linal.Uv, bool) {
	if c.y >= c.ySamples {
		c.y = 0
		return linal.Ray{}, linal.Uv{}, false
	}
	xStep := 1.0 / float32(c.xSamples)
	yStep := 1.0 / float32(c.ySamples)

	s := rand.Float32()
	t := rand.Float32()
	uv := linal.Uv{U: (float32(c.x) + s) * xStep, V: (float32(c.y) + t) * yStep}
	c.x++
	if c.x >= c.xSamples {
		c.x = 0
		c.y++
	}

	origin := linal.Vec3{}
	p1 := linal.Vec3{X: -1, Y: -1, Z: 1}
	p2 := linal.Vec3{X: 1, Y: -1, Z: 1}
	p3 := linal.Vec3{X: -1, Y: 1, Z: 1}
	p4 := linal.Vec3{X: 1, Y: 1, Z: 1}

	p, _ := p1.Lerp(p2, uv.U).Lerp(p3.Lerp(p4, uv.U), uv.V).Normalize()

	mat := c.transform.ToMat()
	origin = mat.ApplyToPoint(origin)
	dir := mat.ApplyToDir(p)

	return linal.Ray{Start: origin, Dir: dir}, uv, true
}

func (c *SimpleCamera) EmitPixel(uv linal.Uv, color materials.Color) {
	c.samples = append(c.samples, struct {
		color materials.Color
		uv    linal.Uv
	}{color, uv})
}

func (c *SimpleCamera) ToImage(width int, height int) Image {
	pixels := make([]materials.Color, width*height)
	xStep := float32(width)
	yStep := float32(height)

	counts := make([]int, width*height)
	for _, sample := range c.samples {
		row := int(sample.uv.V * yStep)
		col := int(sample.uv.U * xStep)
		if col >= width || row >= height || col < 0 || row < 0 {
			continue
		}
		pixels[row*width+col].R += sample.color.R
		pixels[row*width+col].G += sample.color.G
		pixels[row*width+col].B += sample.color.B
		counts[row*width+col]++
	}

	for i, count := range counts {
		if count == 0 {
			continue
		}
		pixels[i].R /= float32(count)
		pixels[i].G /= float32(count)
		pixels[i].B /= float32(count)
	}

	return Image{width, height, pixels}
}
