package camera

import (
	"raytracing/linal"
	"raytracing/materials"
	"raytracing/raytracing"
)

type Tile struct {
	min     linal.Uv
	max     linal.Uv
	samples []Sample
}

type TiledCamera struct {
	tiles       []Tile
	tileSampler TileSampler
	rayProducer RayProducer
}

type TileSampler interface {
	ColorAt(tile Tile, x int, y int) materials.Color
}

type RayProducer interface {
	ShootRay(min linal.Uv, max linal.Uv, plannedCount int, curCount int) linal.Ray
}

func (c *TiledCamera) Shoot(rt raytracing.Raytracer) Image {
	return Image{}
}
