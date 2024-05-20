package camera

import (
	"raytracing/linal"
	"raytracing/materials"
	"raytracing/raytracing"
)

type SimpleCamera struct {
	transform linal.Transform
	img       Image
}

func InitSimpleCamera(transform linal.Transform, width int, height int) SimpleCamera {
	return SimpleCamera{transform, Image{Width: width, Height: height, Pixels: make([]materials.Color, width*height)}}
}

func (s *SimpleCamera) Shoot(rt raytracing.Raytracer) Image {
	mat := s.transform.ToMat()

	for row := 0; row < s.img.Height; row++ {
		for col := 0; col < s.img.Width; col++ {
			u := float32(col)/(0.5*float32(s.img.Width)) - 1
			v := float32(row)/(0.5*float32(s.img.Height)) - 1
			dir, _ := linal.Vec3{X: u, Y: v, Z: 1}.Normalize()
			localRay := linal.Ray{Dir: dir}
			ray := mat.ApplyToRay(localRay)
			color := rt.Sample(ray)
			s.img.Pixels[row*s.img.Width+col] = color
		}
	}

	return s.img
}
