package scene

type Color struct {
	R float32
	G float32
	B float32
}

func (c *Color) RGBA() (uint32, uint32, uint32, uint32) {
	return uint32(c.R * 0xFFFF), uint32(c.G * 0xFFFF), uint32(c.B * 0xFFFF), 0xFFFF
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

func InitMaterial(color Color, reflectiveness float32, isLight bool) SimpleMaterial {
	return SimpleMaterial{color, reflectiveness, isLight}
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
