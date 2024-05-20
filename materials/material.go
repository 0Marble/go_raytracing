package materials

import (
	"raytracing/linal"
)

type Color struct {
	R float32
	G float32
	B float32
}

func (c *Color) RGBA() (uint32, uint32, uint32, uint32) {
	return uint32(c.R * 0xFFFF), uint32(c.G * 0xFFFF), uint32(c.B * 0xFFFF), 0xFFFF
}

func (c Color) Clamp(min Color, max Color) Color {
	res := c
	if res.R < min.R {
		res.R = min.R
	}
	if res.R > max.R {
		res.R = max.R
	}
	if res.G < min.G {
		res.G = min.G
	}
	if res.G > max.G {
		res.G = max.G
	}
	if res.B < min.B {
		res.B = min.B
	}
	if res.B > max.B {
		res.B = max.B
	}

	return res
}

type Material interface {
	Reflect(incoming linal.Vec3, normal linal.Vec3) linal.Vec3
	Lit(incoming linal.Vec3, normal linal.Vec3, toLight linal.Vec3) Color
}
