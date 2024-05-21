package raytracing

import (
	"raytracing/linal"
	"raytracing/materials"
)

type Raytracer interface {
	Sample(ray linal.Ray, time float32) materials.Color
}
