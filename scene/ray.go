package scene

import (
	"raytracing/linal"
)

type Ray struct {
	Dir   linal.Vec3
	Start linal.Vec3
}

func (ray *Ray) Apply(mat *linal.Mat) Ray {
	dir := mat.ApplyToDir(ray.Dir)
	start := mat.ApplyToPoint(ray.Start)

	return Ray{dir, start}
}

func (ray *Ray) GetPoint(t float32) linal.Vec3 {
	return ray.Start.Add(ray.Dir.Mul(t))
}
