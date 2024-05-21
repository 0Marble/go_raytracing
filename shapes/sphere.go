package shapes

import (
	"math"
	"raytracing/linal"
	"raytracing/materials"
)

type Sphere struct {
	material  materials.Material
	transform linal.TimedTransform
}

func InitSphere(transform linal.TimedTransform, material materials.Material) Sphere {
	return Sphere{material, transform}
}

func (s *Sphere) Intersect(ray linal.Ray, time float32) Intersection {
	mat := s.transform.ToMat(time)
	inv := s.transform.ToInv(time)

	localRay := inv.ApplyToRay(ray)
	a := localRay.Dir.LenSquared()
	b := 2.0 * localRay.Dir.Dot(localRay.Start)
	c := localRay.Start.LenSquared() - 1

	d := b*b - 4*a*c
	if d < 0.0 {
		return Intersection{IsHit: false}
	}
	sqrtD := float32(math.Sqrt(float64(d)))

	t1 := (-b - sqrtD) / (2 * a)
	t2 := (-b + sqrtD) / (2 * a)

	if t1 < 0.0 && t2 < 0.0 {
		return Intersection{IsHit: false}
	}
	t := t1
	if t1 < 0.0 || (t1 > t2 && t2 >= 0.0) {
		t = t2
	}

	pt := localRay.Start.Add(localRay.Dir.Mul(t))

	return Intersection{Uv: s.ToUv(pt, time), IsHit: true, T: mat.ApplyToPoint(pt).Sub(ray.Start).Len() / ray.Dir.Len()}
}

func (s *Sphere) Normal(uv linal.Uv, time float32) linal.Vec3 {
	mat := s.transform.ToMat(time)
	spLoc := linal.Vec3{X: 1, Y: uv.U * 2 * math.Pi, Z: uv.V * math.Pi}
	loc := spLoc.FromSpherical()
	trans := mat.Transpose()
	trans, _ = trans.Inverse()
	normal := trans.ApplyToDir(loc)
	res, _ := normal.Normalize()
	return res
}

func (s *Sphere) FromUv(uv linal.Uv, time float32) linal.Vec3 {
	mat := s.transform.ToMat(time)
	theta := uv.U * 2 * math.Pi
	phi := uv.V * math.Pi
	return mat.ApplyToPoint(linal.Vec3{X: 1.0, Y: theta, Z: phi}.FromSpherical())
}

func (s *Sphere) ToUv(pt linal.Vec3, time float32) linal.Uv {
	inv := s.transform.ToInv(time)
	local := inv.ApplyToPoint(pt).ToSpherical()
	return linal.Uv{U: local.Y / (2 * math.Pi), V: local.Z / math.Pi}
}

func (s *Sphere) TransformMat(time float32) linal.Mat {
	return s.transform.ToMat(time)
}
func (s *Sphere) InverseTransformMat(time float32) linal.Mat {
	return s.transform.ToInv(time)
}

func (s *Sphere) Material() *materials.Material {
	return &s.material
}

func (s *Sphere) Aabb(time float32) linal.Aabb {
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
