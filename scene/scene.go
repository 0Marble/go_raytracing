package scene

import (
	"raytracing/lights"
	"raytracing/linal"
	"raytracing/materials"
	"raytracing/shapes"
)

type Scene struct {
	objs    []shapes.Object
	lights  []lights.Light
	Outside materials.Material
}

func InitScene(objs []shapes.Object, lights []lights.Light, outside materials.Material) Scene {
	return Scene{objs, lights, outside}
}

func (s *Scene) Intersect(ray linal.Ray) (*shapes.Object, shapes.Intersection) {
	var resObj *shapes.Object = nil
	res := shapes.Intersection{}
	minDist := float32(0.0)

	for i, obj := range s.objs {
		intersection := obj.Intersect(ray)
		if !intersection.IsHit {
			continue
		}

		pt := obj.FromUv(intersection.Uv)
		dist := pt.Sub(ray.Start).Len()

		if dist < minDist || resObj == nil {
			minDist = dist
			resObj = &s.objs[i]
			res = intersection
		}
	}
	return resObj, res
}

func (s *Scene) Lights() []lights.Light {
	return s.lights
}

func (s *Scene) TotalAabb() linal.Aabb {
	aabb := linal.Aabb{}

	for _, obj := range s.objs {
		box := obj.Aabb()
		aabb = aabb.Merge(&box)
	}

	return aabb

}
