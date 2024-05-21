package lights

import (
	"log"
	"math"
	"raytracing/linal"
	"raytracing/materials"
)

type Spotlight struct {
	pos   linal.Vec3
	dir   linal.Vec3
	color materials.Color
}

func InitSpotlight(pos linal.Vec3, dir linal.Vec3, color materials.Color) Spotlight {
	d, ok := dir.Normalize()
	if !ok {
		log.Fatal("No light direction given!")
	}

	return Spotlight{pos, d, color}
}

func (s *Spotlight) RayToLight(pos linal.Vec3) (linal.Ray, linal.Uv) {
	dir, ok := s.pos.Sub(pos).Normalize()
	if !ok {
		log.Fatal("RayToLight starting from the light!", s)
	}
	sp := dir.Mul(-1).ToSpherical()
	uv := linal.Uv{U: sp.Y / (2.0 * math.Pi), V: sp.Z / (math.Pi)}

	return linal.Ray{Start: pos, Dir: dir}, uv
}
func (s *Spotlight) GetColor(uv linal.Uv, pos linal.Vec3) materials.Color {
	dir, _ := linal.Vec3{X: 1, Y: uv.U * 2 * math.Pi, Z: uv.V * math.Pi}.FromSpherical().Normalize()
	power := dir.Dot(s.dir)
	return s.color.MulNum(power).Clamp(materials.Color{R: 0, G: 0, B: 0}, materials.Color{R: 1, G: 1, B: 1})
}
func (s *Spotlight) TDist(ray linal.Ray) float32 {
	return ray.Start.Sub(s.pos).Len() / ray.Dir.Len()
}
