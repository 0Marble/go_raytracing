package scene

import (
	"raytracing/linal"
)

type Uv struct {
	U float32
	V float32
}

type Intersection struct {
	Uv    Uv
	T     float32
	IsHit bool
}

type Object interface {
	Intersect(ray Ray) Intersection
	Normal(pt Uv) linal.Vec3
	FromUv(pt Uv) linal.Vec3
	ToUv(pt linal.Vec3) Uv
	Material() *Material
	Aabb() linal.Aabb
}
