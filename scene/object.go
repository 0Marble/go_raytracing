package scene

import (
	"raytracing/linal"
	"raytracing/transfrom"
)

type Uv struct {
	U float32
	V float32
}

type Object interface {
	Distance(pt linal.Vec3) float32
	Normal(pt Uv) linal.Vec3
	FromUv(pt Uv) linal.Vec3
	ToUv(pt linal.Vec3) Uv
	Transform() *transfrom.Transform
	Material() *Material
	Aabb() linal.Aabb
}
