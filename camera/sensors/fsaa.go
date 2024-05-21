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

func (s *FsaaNByNSensor) GetPixel(rt raytracing.Raytracer, lens camera.Lens, x, y, width, height int, time float32) materials.Color {
	step := 1.0 / float32(s.n)
	color := materials.Color{}
	pWidth := 1.0 / float32(width)
	pHeight := 1.0 / float32(height)

	u := float32(x) / float32(width)
	v := float32(y) / float32(height)
	for i := 0; i < s.n; i++ {
		for j := 0; j < s.n; j++ {
			uv := linal.Uv{
				U: u + step*float32(i)*pWidth + 0.5*pWidth*step,
				V: v + step*float32(j)*pHeight + 0.5*pHeight*step,
			}

			c := rt.Sample(lens.ShootRay(uv), time)
			color = color.Add(c)
		}
	}
	return color.DivNum(float32(s.n * s.n))
}

type FsaaRgssSensor struct{}

func InitFsaaRgssSensor() FsaaRgssSensor {
	return FsaaRgssSensor{}
}

func (s *FsaaRgssSensor) GetPixel(rt raytracing.Raytracer, lens camera.Lens, x, y, width, height int, time float32) materials.Color {
	color := materials.Color{}
	pWidth := 1.0 / float32(width)
	pHeight := 1.0 / float32(height)

	u := float32(x) / float32(width)
	v := float32(y) / float32(height)
	xpts := []int{2, 0, 3, 1}
	step := float32(0.25)
	for j, i := range xpts {
		uv := linal.Uv{
			U: u + step*float32(i)*pWidth + 0.5*pWidth*step,
			V: v + step*float32(j)*pHeight + 0.5*pHeight*step,
		}

		c := rt.Sample(lens.ShootRay(uv), time)
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

func (s *FsaaCheckerboardSensor) GetPixel(rt raytracing.Raytracer, lens camera.Lens, x, y, width, height int, time float32) materials.Color {
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
				U: u + step*float32(i)*pWidth + 0.5*pWidth*step,
				V: v + step*float32(j)*pHeight + 0.5*pHeight*step,
			}

			ray := lens.ShootRay(uv)
			c := rt.Sample(ray, time)
			color = color.Add(c)
		}
	}
	return color.DivNum(float32(cnt))
}

type maybeColor struct {
	ok    bool
	color materials.Color
}

type FlipquadSensor struct {
	samples []maybeColor
	left    int
	right   int
	bottom  int
	top     int
}

func InitFlipquadSensor(left int, right int, bottom int, top int) FlipquadSensor {
	w := right - left
	h := top - bottom
	return FlipquadSensor{make([]maybeColor, w*(h+1)*2), left, right, bottom, top}
}

func (s *FlipquadSensor) GetPixel(rt raytracing.Raytracer, lens camera.Lens, x, y, width, height int, time float32) materials.Color {
	color := materials.Color{}
	pWidth := 1.0 / float32(width)
	pHeight := 1.0 / float32(height)

	u := float32(x)/float32(width) + 0.5*pWidth
	v := float32(y)/float32(height) + 0.5*pHeight
	var xpts [4]int
	if (x+y)%2 == 0 {
		xpts = [4]int{2, 0, 3, 1}
	} else {
		xpts = [4]int{1, 3, 0, 2}
	}
	w := s.right - s.left

	for j, i := range xpts {
		idx := 0
		if j == 0 || j == 1 {
			idx = ((y-s.bottom)*w + (x - s.left)) * 2
		} else {
			idx = ((y-s.bottom+1)*w + (x - s.left)) * 2
		}
		if i == 2 || i == 3 {
			idx += 1
		}
		if s.samples[idx].ok {
			color = color.Add(s.samples[idx].color)
			continue
		}

		du := float32(0.0)
		dv := float32(0.0)
		switch i {
		case 0:
			du = -0.5 * pWidth
		case 1:
			du = -1.0 / 8.0 * pWidth
		case 2:
			du = 1.0 / 8.0 * pWidth
		case 3:
			du = 0.5 * pWidth
		}
		switch j {
		case 0:
			dv = -0.5 * pHeight
		case 1:
			dv = -1.0 / 8.0 * pHeight
		case 2:
			dv = 1.0 / 8.0 * pHeight
		case 3:
			dv = 0.5 * pHeight
		}

		uv := linal.Uv{
			U: u + du,
			V: v + dv,
		}

		c := rt.Sample(lens.ShootRay(uv), time)
		s.samples[idx] = maybeColor{true, c}
		color = color.Add(c)
	}
	return color.DivNum(float32(len(xpts)))
}
