package sensors

import (
	"raytracing/camera"
	"raytracing/linal"
	"raytracing/materials"
	"raytracing/raytracing"
)

type FsaaNByNSensor struct {
	n int
}

func InitFsaaNByNSensor(n int) FsaaNByNSensor {
	return FsaaNByNSensor{n}
}

func (s *FsaaNByNSensor) GetPixel(rt raytracing.Raytracer, lens camera.Lens, x, y, width, height int) materials.Color {
	step := 1.0 / float32(s.n)
	color := materials.Color{}
	pWidth := 1.0 / float32(width)
	pHeight := 1.0 / float32(height)

	u := float32(x) / float32(width)
	v := float32(y) / float32(height)
	for i := 0; i < s.n; i++ {
		for j := 0; j < s.n; j++ {
			uv := linal.Uv{
				U: u + step*float32(i)*pWidth,
				V: v + step*float32(j)*pHeight,
			}

			c := rt.Sample(lens.ShootRay(uv))
			color = color.Add(c)
		}
	}
	return color.DivNum(float32(s.n * s.n))
}

type FsaaRgssSensor struct{}

func InitFsaaRgssSensor() FsaaRgssSensor {
	return FsaaRgssSensor{}
}

func (s *FsaaRgssSensor) GetPixel(rt raytracing.Raytracer, lens camera.Lens, x, y, width, height int) materials.Color {
	color := materials.Color{}
	pWidth := 1.0 / float32(width)
	pHeight := 1.0 / float32(height)

	u := float32(x) / float32(width)
	v := float32(y) / float32(height)
	xpts := []int{2, 0, 3, 1}
	step := float32(0.25)
	for j, i := range xpts {
		uv := linal.Uv{
			U: u + step*float32(i)*pWidth + 0.5*pWidth,
			V: v + step*float32(j)*pHeight + 0.5*pHeight,
		}

		c := rt.Sample(lens.ShootRay(uv))
		color = color.Add(c)
	}
	return color.DivNum(float32(len(xpts)))
}

type FsaaCheckerboardSensor struct {
	n int
}

func InitFsaaCheckerboardSensor(n int) FsaaCheckerboardSensor {
	return FsaaCheckerboardSensor{n}
}

func (s *FsaaCheckerboardSensor) GetPixel(rt raytracing.Raytracer, lens camera.Lens, x, y, width, height int) materials.Color {
	step := 1.0 / float32(s.n)
	color := materials.Color{}
	pWidth := 1.0 / float32(width)
	pHeight := 1.0 / float32(height)

	u := float32(x) / float32(width)
	v := float32(y) / float32(height)
	cnt := 0
	for i := 0; i < s.n; i++ {
		for j := 0; j < s.n; j++ {
			if (i+j)%2 == 0 {
				continue
			}
			cnt++
			uv := linal.Uv{
				U: u + step*float32(i)*pWidth,
				V: v + step*float32(j)*pHeight,
			}

			ray := lens.ShootRay(uv)
			c := rt.Sample(ray)
			color = color.Add(c)
		}
	}
	return color.DivNum(float32(cnt))
}
