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
	Intersect(ray linal.Ray, time float32) Intersection
	Normal(pt linal.Uv, time float32) linal.Vec3
	FromUv(pt linal.Uv, time float32) linal.Vec3
	ToUv(pt linal.Vec3, time float32) linal.Uv
	Material() *materials.Material
	Aabb(time float32) linal.Aabb
}
