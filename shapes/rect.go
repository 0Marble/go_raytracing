package shapes

import (
	"raytracing/linal"
	"raytracing/materials"
)

type Rect struct {
	material  materials.Material
	transform linal.TimedTransform
}

func InitRect(transform linal.TimedTransform, material materials.Material) Rect {
	return Rect{material, transform}
}

func (s *Rect) Intersect(ray linal.Ray, time float32) Intersection {
	mat := s.transform.ToMat(time)
	inv := s.transform.ToInv(time)

	origin := inv.ApplyToPoint(ray.Start)
	dir := inv.ApplyToDir(ray.Dir)
	if origin.Z == 0 {
		uv := linal.Uv{U: origin.X + 0.5, V: origin.Y + 0.5}
		if uv.U > 1 || uv.U < 0 || uv.V > 1 || uv.V < 0 {
			return Intersection{IsHit: false}
		}

		return Intersection{IsHit: true, Uv: uv, T: 0.0}
	} else if dir.Z == 0 {
		return Intersection{IsHit: false}
	}

	t := -origin.Z / dir.Z
	if t <= 0.0 {
		return Intersection{IsHit: false}
	}
	pt := origin.Add(dir.Mul(t))
	uv := linal.Uv{U: pt.X + 0.5, V: pt.Y + 0.5}
	if uv.U > 1 || uv.U < 0 || uv.V > 1 || uv.V < 0 {
		return Intersection{IsHit: false}
	}

	globalPt := mat.ApplyToPoint(pt)
	tDist := globalPt.Sub(ray.Start).Len() / ray.Dir.Len()

	return Intersection{IsHit: true, Uv: uv, T: tDist}
}

func (s *Rect) Normal(uv linal.Uv, time float32) linal.Vec3 {
	mat := s.transform.ToMat(time)
	norm := linal.Vec3{Z: -1}
	trans := mat.Transpose()
	trans, _ = trans.Inverse()
	normal := trans.ApplyToDir(norm)
	res, _ := normal.Normalize()
	return res
}

func (s *Rect) FromUv(uv linal.Uv, time float32) linal.Vec3 {
	pt := linal.Vec3{X: uv.U - 0.5, Y: uv.V - 0.5, Z: 0}
	mat := s.transform.ToMat(time)
	return mat.ApplyToPoint(pt)
}

func (s *Rect) ToUv(pt linal.Vec3, time float32) linal.Uv {
	inv := s.transform.ToInv(time)
	uv := inv.ApplyToPoint(pt)
	return linal.Uv{U: uv.X + 0.5, V: uv.Y + 0.5}
}

func (s *Rect) TransformMat(time float32) linal.Mat {
	return s.transform.ToMat(time)
}
func (s *Rect) InverseTransformMat(time float32) linal.Mat {
	return s.transform.ToInv(time)
}

func (s *Rect) Material() *materials.Material {
	return &s.material
}

func (s *Rect) Aabb(time float32) linal.Aabb {
	min := linal.Vec3{}
	max := linal.Vec3{}
	mat := s.transform.ToMat(time)

	center := linal.Vec3{}
	for dx := float32(-1.0); dx <= 1; dx++ {
		for dy := float32(-1.0); dy <= 1; dy++ {
			for dz := float32(-1.0); dz <= 1; dz++ {
				p := center.Add(linal.Vec3{X: dx, Y: dy, Z: dz})
				p = mat.ApplyToPoint(p)
				min = min.Min(p)
				max = max.Max(p)

			}
		}
	}

	return linal.Aabb{Min: min, Max: max}
}
