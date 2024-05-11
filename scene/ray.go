package scene

import (
	"raytracing/linal"
)

type Ray struct {
	Dir   linal.Vec3
	Start linal.Vec3
	Step  float32
}

func (ray *Ray) Apply(mat *linal.Mat) Ray {
	dir := mat.ApplyToDir(ray.Dir)
	start := mat.ApplyToPoint(ray.Start)

	l := dir.Len()
	return Ray{dir, start, ray.Step * l}
}

func (ray *Ray) GetPoint(t float32) linal.Vec3 {
	return ray.Start.Add(ray.Dir.Mul(t))
}

func (ray *Ray) AdvanceBy(dist float32) Ray {
	return Ray{Start: ray.Start.Add(ray.Dir.Mul(dist)), Dir: ray.Dir, Step: ray.Step}
}
