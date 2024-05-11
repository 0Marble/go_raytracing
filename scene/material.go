package scene

type Color struct {
	R float32
	G float32
	B float32
}

type Material interface {
	Color(pt Uv) Color
	Reflectiveness(pt Uv) float32
	EmitsLight() bool
}

type SimpleMaterial struct {
	color          Color
	reflectiveness float32
	isLight        bool
}

func (m *SimpleMaterial) Color(pt Uv) Color {
	return m.color
}
func (m *SimpleMaterial) Reflectiveness(pt Uv) float32 {
	return m.reflectiveness
}
func (m *SimpleMaterial) EmitsLight() bool {
	return m.isLight
}
