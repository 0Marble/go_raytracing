package lights

import (
	"raytracing/linal"
	"raytracing/materials"
)

type Light interface {
	RayToLight(pos linal.Vec3) (linal.Ray, linal.Uv)
	GetColor(uv linal.Uv, pos linal.Vec3) materials.Color
	TDist(ray linal.Ray) float32
}
