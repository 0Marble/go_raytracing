package camera

import (
	"raytracing/linal"
	"raytracing/materials"
	"raytracing/raytracing"
)

type SimpleCamera struct {
	transform linal.Transform
}

func InitSimpleCamera(transform linal.Transform) SimpleCamera {
	return SimpleCamera{transform}
}

func (s *SimpleCamera) Shoot(rt raytracing.Raytracer, width int, height int, left int, right int, top int, bottom int) Image {
	mat := s.transform.ToMat()

	w := right - left
	h := top - bottom
	img := Image{w, h, make([]materials.Color, w*h)}

	for row := bottom; row < top; row++ {
		for col := left; col < right; col++ {
			u := float32(col)/(0.5*float32(width)) - 1
			v := float32(row)/(0.5*float32(height)) - 1
			dir, _ := linal.Vec3{X: u, Y: v, Z: 1}.Normalize()
			localRay := linal.Ray{Dir: dir}
			ray := mat.ApplyToRay(localRay)
			color := rt.Sample(ray)
			img.Pixels[(row-bottom)*w+(col-left)] = color
		}
	}

	return img
}
func (c *SimpleCamera) Transform() linal.Transform {
	return c.transform
}
