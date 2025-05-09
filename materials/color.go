package materials

import (
	"math"
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

func (c Color) Add(a Color) Color {
	return Color{R: c.R + a.R, G: c.G + a.G, B: c.B + a.B}
}
func (c Color) Sub(a Color) Color {
	return Color{R: c.R - a.R, G: c.G - a.G, B: c.B - a.B}
}
func (c Color) Mul(a Color) Color {
	return Color{R: c.R * a.R, G: c.G * a.G, B: c.B * a.B}
}
func (c Color) Div(a Color) Color {
	return Color{R: c.R / a.R, G: c.G / a.G, B: c.B / a.B}
}
func (c Color) MulNum(a float32) Color {
	return Color{R: c.R * a, G: c.G * a, B: c.B * a}
}
func (c Color) DivNum(a float32) Color {
	return Color{R: c.R / a, G: c.G / a, B: c.B / a}
}
func (c Color) PowNum(t float32) Color {
	return Color{
		R: float32(math.Pow(float64(c.R), float64(t))),
		G: float32(math.Pow(float64(c.G), float64(t))),
		B: float32(math.Pow(float64(c.B), float64(t))),
	}
}

func (c Color) Color(toLight linal.Vec3, toView linal.Vec3, normal linal.Vec3) Color {
	return c
}
