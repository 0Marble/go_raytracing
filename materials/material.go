package materials

import (
	"raytracing/linal"
)

type Material interface {
	Color(toLight linal.Vec3, toView linal.Vec3, normal linal.Vec3) Color
}
