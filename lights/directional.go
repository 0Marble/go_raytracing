package lights

import (
	"log"
	"raytracing/linal"
	"raytracing/materials"
)

type DirectionalLight struct {
	dir   linal.Vec3
	color materials.Color
}

func InitDirectionalLight(dir linal.Vec3, color materials.Color) DirectionalLight {
	d, ok := dir.Normalize()
	if !ok {
		log.Fatal("Direction set to zeros!", dir)
	}
	return DirectionalLight{d, color}
}

func (d *DirectionalLight) RayToLight(pos linal.Vec3) (linal.Ray, linal.Uv) {
	return linal.Ray{Start: pos, Dir: d.dir.Mul(-1)}, linal.Uv{}
}
func (d *DirectionalLight) GetColor(uv linal.Uv, pos linal.Vec3) materials.Color {
	return d.color
}
func (d *DirectionalLight) TDist(ray linal.Ray) float32 {
	return 0.0
}
