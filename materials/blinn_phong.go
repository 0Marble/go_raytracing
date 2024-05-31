package materials

import (
	"log"
	"math"
	"raytracing/linal"
)

type BlinnPhong struct {
	diffuseColor  Color
	specularColor Color
	shininess     float32
}

func InitBlinnPhong(diffuseColor Color, specularColor Color, shininess float32) BlinnPhong {
	return BlinnPhong{diffuseColor, specularColor, shininess}
}

func (m *BlinnPhong) Color(toLight linal.Vec3, toView linal.Vec3, normal linal.Vec3) Color {
	specular := float32(0)
	lambertian := max(toLight.Dot(normal), 0.0)

	if lambertian > 0 {
		halfDir, ok := toLight.Add(toView).Normalize()
		if !ok {
			log.Println("Totally flat ray", toLight, toView)
			return Color{R: 1, G: 1, B: 1}
		}
		specAngle := max(halfDir.Dot(normal), 0.0)
		specular = float32(math.Pow(float64(specAngle), float64(m.shininess)))
	}

	return m.diffuseColor.MulNum(lambertian).Add(m.specularColor.MulNum(specular))
}
