package materials

import "raytracing/linal"

type SimpleMaterial struct {
	color          Color
	reflectiveness float32
}

func InitSimpleMaterial(color Color, reflectiveness float32) SimpleMaterial {
	return SimpleMaterial{color, reflectiveness}
}

func (m *SimpleMaterial) Color(pt linal.Uv) Color {
	return m.color
}
func (m *SimpleMaterial) Reflectiveness(pt linal.Uv) float32 {
	return m.reflectiveness
}
