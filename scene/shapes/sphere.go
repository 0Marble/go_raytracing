package shapes

import (
	"log"
	"math"
	"raytracing/linal"
	"raytracing/scene"
	"raytracing/transfrom"
)

type Sphere struct {
	material scene.Material
	mat      linal.Mat
	inv      linal.Mat
}

func InitSphere(transform transfrom.Transform, material scene.Material) Sphere {
	mat := transform.ToMat()
	inv, ok := mat.Inverse()
	if !ok {
		log.Fatal("Sphere : transform non-invertable")
	}
	return Sphere{mat: mat, inv: inv, material: material}
}

func (s *Sphere) Intersect(ray scene.Ray) scene.Intersection {
	localRay := ray.Apply(&s.inv)
	a := localRay.Dir.LenSquared()
	b := 2.0 * localRay.Dir.Dot(localRay.Start)
	c := localRay.Start.LenSquared() - 1

	d := b*b - 4*a*c
	if d < 0.0 {
		return scene.Intersection{IsHit: false}
	}
	sqrtD := float32(math.Sqrt(float64(d)))

	t1 := (-b - sqrtD) / (2 * a)
	t2 := (-b + sqrtD) / (2 * a)

	if t1 < 0.0 && t2 < 0.0 {
		return scene.Intersection{IsHit: false}
	}
	t := t1
	if t1 < 0.0 || (t1 > t2 && t2 >= 0.0) {
		t = t2
	}

	pt := localRay.Start.Add(localRay.Dir.Mul(t))

	return scene.Intersection{Uv: s.ToUv(pt), IsHit: true}
}

func (s *Sphere) Normal(uv scene.Uv) linal.Vec3 {
	spLoc := linal.Vec3{X: 1, Y: uv.U * 2 * math.Pi, Z: uv.V * math.Pi}
	loc := spLoc.FromSpherical()
	trans := s.mat.Transpose()
	trans, _ = trans.Inverse()
	normal := trans.ApplyToDir(loc)
	res, _ := normal.Normalize()
	return res
}

func (s *Sphere) FromUv(uv scene.Uv) linal.Vec3 {
	theta := uv.U * 2 * math.Pi
	phi := uv.V * math.Pi
	return s.mat.ApplyToPoint(linal.Vec3{X: 1.0, Y: theta, Z: phi}.FromSpherical())
}

func (s *Sphere) ToUv(pt linal.Vec3) scene.Uv {
	local := s.inv.ApplyToPoint(pt).ToSpherical()

	return scene.Uv{U: local.Y / (2 * math.Pi), V: local.Z / math.Pi}
}

func (s *Sphere) TransformMat() linal.Mat {
	return s.mat
}
func (s *Sphere) InverseTransformMat() linal.Mat {
	return s.inv
}

func (s *Sphere) Material() *scene.Material {
	return &s.material
}

func (s *Sphere) Aabb() linal.Aabb {
	min := linal.Vec3{}
	max := linal.Vec3{}

	center := linal.Vec3{}
	for dx := float32(-1.0); dx <= 1; dx++ {
		for dy := float32(-1.0); dy <= 1; dy++ {
			for dz := float32(-1.0); dz <= 1; dz++ {
				p := center.Add(linal.Vec3{X: dx, Y: dy, Z: dz})
				p = s.mat.ApplyToPoint(p)
				min = min.Min(p)
				max = max.Max(p)

			}
		}
	}

	return linal.Aabb{Min: min, Max: max}
}
