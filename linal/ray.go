package linal

type Ray struct {
	Dir   Vec3
	Start Vec3
}

func (ray *Ray) Apply(mat *Mat) Ray {
	dir := mat.ApplyToDir(ray.Dir)
	start := mat.ApplyToPoint(ray.Start)

	return Ray{dir, start}
}

func (ray *Ray) GetPoint(t float32) Vec3 {
	return ray.Start.Add(ray.Dir.Mul(t))
}
