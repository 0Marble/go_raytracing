package scene

import (
	"raytracing/linal"
	"raytracing/transfrom"
)

type Image struct {
	Width  int
	Height int
	Pixels []Color
}

type Camera interface {
	Transform() *transfrom.Transform
	ShootRay() (Ray, Uv, bool)
	EmitPixel(uv Uv, color Color)
	ToImage(width int, height int) Image
}

type SimpleCamera struct {
	transform transfrom.Transform
	x         int
	y         int
	xSamples  int
	ySamples  int
	samples   []struct {
		color Color
		uv    Uv
	}
}

func (c *SimpleCamera) Transform() *transfrom.Transform {
	return &c.transform
}

func (c *SimpleCamera) ShootRay() (Ray, Uv, bool) {
	if c.y >= c.ySamples {
		return Ray{}, Uv{}, false
	}
	xStep := 1.0 / float32(c.xSamples)
	yStep := 1.0 / float32(c.ySamples)

	uv := Uv{U: float32(c.x) * xStep, V: float32(c.y) * yStep}
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

	return Ray{Start: origin, Dir: dir, Step: 1.0}, uv, true
}

func (c *SimpleCamera) EmitPixel(uv Uv, color Color) {
	c.samples = append(c.samples, struct {
		color Color
		uv    Uv
	}{color, uv})
}

func (c *SimpleCamera) ToImage(width int, height int) Image {
	pixels := make([]Color, width*height)
	xStep := 1.0 / float32(c.xSamples)
	yStep := 1.0 / float32(c.ySamples)

	counts := make([]int, width*height)
	for _, sample := range c.samples {
		row := int(sample.uv.V / yStep)
		col := int(sample.uv.U / xStep)
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
