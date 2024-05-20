package camera

import (
	"log"
	"raytracing/linal"
	"raytracing/materials"
	"raytracing/raytracing"
)

const (
	N_BY_N = iota
	RGSS
	CHECKER
	FLIPQUAD
)

type FsaaCamera struct {
	transform linal.Transform
	kind      int
	n         int
}

func InitFsaaCamera(transform linal.Transform, kind int, n int) FsaaCamera {
	if kind != RGSS && kind != N_BY_N && kind != CHECKER && kind != FLIPQUAD {
		log.Fatal("Unknown FSAA kind", kind)
	}
	return FsaaCamera{transform, kind, n}
}

func (c *FsaaCamera) Shoot(rt raytracing.Raytracer, width int, height int, left int, right int, top int, bottom int) Image {
	mat := c.transform.ToMat()

	w := right - left
	h := top - bottom
	img := Image{w, h, make([]materials.Color, w*h)}
	pWidth := 1.0 / float32(w)
	pHeight := 1.0 / float32(h)

	for row := bottom; row < top; row++ {
		for col := left; col < right; col++ {
			color := materials.Color{}
			u := float32(col)/(0.5*float32(width)) - 1
			v := float32(row)/(0.5*float32(height)) - 1

			switch c.kind {
			case N_BY_N:
				step := 1.0 / float32(c.n)
				for i := 0; i < c.n; i++ {
					for j := 0; j < c.n; j++ {
						d, _ := linal.Vec3{
							X: u + step*float32(i)*pWidth,
							Y: v + step*float32(j)*pHeight,
							Z: 1}.Normalize()
						c := rt.Sample(mat.ApplyToRay(linal.Ray{Dir: d}))
						color = color.Add(c)
					}
				}
				color = color.DivNum(float32(c.n * c.n))
			case RGSS:
				xpts := []int{2, 0, 3, 1}
				step := float32(0.25)
				for j, i := range xpts {
					d, _ := linal.Vec3{
						X: u + step*float32(i)*pWidth,
						Y: v + step*float32(j)*pHeight,
						Z: 1}.Normalize()
					c := rt.Sample(mat.ApplyToRay(linal.Ray{Dir: d}))
					color = color.Add(c)
				}
				color = color.DivNum(float32(len(xpts)))
			case CHECKER:
				step := 1.0 / float32(c.n)
				cnt := 0
				for i := 0; i < c.n; i++ {
					for j := 0; j < c.n; j++ {
						if (i+j)%2 == 0 {
							continue
						}
						cnt += 1
						d, _ := linal.Vec3{
							X: u + step*float32(i)*pWidth,
							Y: v + step*float32(j)*pHeight,
							Z: 1}.Normalize()
						c := rt.Sample(mat.ApplyToRay(linal.Ray{Dir: d}))
						color = color.Add(c)
					}
				}
				color = color.DivNum(float32(cnt))

			}

			img.Pixels[(row-bottom)*w+(col-left)] = color
		}
	}

	return img
}

func (c *FsaaCamera) Transform() linal.Transform {
	return c.transform
}
