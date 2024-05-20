package materials

import (
	"raytracing/linal"
)

type SimpleMaterial struct {
	color Color
}

func InitSimpleMaterial(color Color) SimpleMaterial {
	return SimpleMaterial{color}
}

func (m *SimpleMaterial) Reflect(incoming linal.Vec3, normal linal.Vec3) linal.Vec3 {
	reflectedDir := incoming.Sub(normal.Mul(incoming.Dot(normal) * 2.0))
	return reflectedDir
}

func (m *SimpleMaterial) Lit(incoming linal.Vec3, normal linal.Vec3, toLight linal.Vec3) Color {
	return m.color
}
