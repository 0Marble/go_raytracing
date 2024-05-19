package materials

import "raytracing/linal"

type Color struct {
	R float32
	G float32
	B float32
}

func (c *Color) RGBA() (uint32, uint32, uint32, uint32) {
	return uint32(c.R * 0xFFFF), uint32(c.G * 0xFFFF), uint32(c.B * 0xFFFF), 0xFFFF
}

type Material interface {
	Color(pt linal.Uv) Color
	Reflectiveness(pt linal.Uv) float32
}
