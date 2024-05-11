package shapes

import (
	"log"
	"math"
	"raytracing/linal"
	"raytracing/scene"
	"raytracing/transfrom"
)

type Sphere struct {
	transform transfrom.Transform
	material  scene.Material
}

func InitSphere(transform transfrom.Transform, material scene.Material) Sphere {
	return Sphere{transform, material}
}

func (s *Sphere) intersect(ray scene.Ray) (float32, bool) {
	mat := s.transform.ToMat()
	inv, ok := mat.Inverse()
	if !ok {
		log.Fatalf("Sphere %p : transform non-invertable", s)
	}

	localRay := ray.Apply(&inv)
	a := localRay.Dir.LenSquared()
	b := localRay.Dir.Dot(localRay.Start)
	c := localRay.Start.LenSquared()

	d := b*b - 4*a*c
	if d < 0.0 {
		return 0.0, false
	}
	sqrtD := float32(math.Sqrt(float64(d)))

	t1 := (-b - sqrtD) / (2 * a)
	t2 := (-b + sqrtD) / (2 * a)

	if t1 < 0.0 && t2 < 0.0 {
		return 0.0, false
	}
	t := t1
	if t1 < 0.0 || (t1 > t2 && t2 >= 0.0) {
		t = t2
	}

	return t, true
}

func (s *Sphere) Distance(pt linal.Vec3) float32 {
	center := s.transform.Translation
	return pt.Sub(center).Len()
}

func (s *Sphere) Normal(uv scene.Uv) linal.Vec3 {
	center := linal.Vec3{}.Add(s.transform.Translation)
	res, _ := s.FromUv(uv).Sub(center).Normalize()
	return res
}

func (s *Sphere) FromUv(uv scene.Uv) linal.Vec3 {
	theta := uv.U * 2 * math.Pi
	phi := uv.V * math.Pi
	mat := s.transform.ToMat()
	return mat.ApplyToPoint(linal.Vec3{X: 1.0, Y: theta, Z: phi}.FromSpherical())
}

func (s *Sphere) ToUv(pt linal.Vec3) scene.Uv {
	mat := s.transform.ToMat()
	inv, ok := mat.Inverse()
	if !ok {
		log.Fatalf("Sphere %p : transform not invertable", s)
	}
	local := inv.ApplyToPoint(pt).ToSpherical()

	return scene.Uv{U: local.Y / (2 * math.Pi), V: local.Z / math.Pi}
}

func (s *Sphere) Transform() *transfrom.Transform {
	return &s.transform
}

func (s *Sphere) Material() *scene.Material {
	return &s.material
}

func (s *Sphere) Aabb() linal.Aabb {
	min := linal.Vec3{}
	max := linal.Vec3{}

	mat := s.transform.ToMat()

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
