package lights

import (
	"log"
	"raytracing/linal"
	"raytracing/materials"
)

type PointLight struct {
	pos   linal.Vec3
	color materials.Color
}

func InitPointLight(pos linal.Vec3, color materials.Color) PointLight {
	return PointLight{pos, color}
}

func (p *PointLight) RayToLight(pos linal.Vec3) (linal.Ray, linal.Uv) {
	dir, ok := p.pos.Sub(pos).Normalize()
	if !ok {
		log.Fatal("RayToLight starting from the light!", p)
	}

	return linal.Ray{Start: pos, Dir: dir}, linal.Uv{}
}

func (p *PointLight) GetColor(uv linal.Uv, pos linal.Vec3) materials.Color {
	return p.color
}
func (p *PointLight) TDist(ray linal.Ray) float32 {
	return ray.Start.Sub(p.pos).Len() / ray.Dir.Len()
}
