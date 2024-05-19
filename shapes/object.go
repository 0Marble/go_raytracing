package shapes

import (
	"raytracing/linal"
	"raytracing/materials"
)

type Intersection struct {
	Uv    linal.Uv
	T     float32
	IsHit bool
}

type Object interface {
	Intersect(ray linal.Ray) Intersection
	Normal(pt linal.Uv) linal.Vec3
	FromUv(pt linal.Uv) linal.Vec3
	ToUv(pt linal.Vec3) linal.Uv
	Material() *materials.Material
	Aabb() linal.Aabb
}
