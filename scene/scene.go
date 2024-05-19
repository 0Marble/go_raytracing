package scene

import (
	"raytracing/linal"
)

type Scene struct {
	objs    []Object
	Cam     Camera
	Outside Material
}

func InitScene(objs []Object, cam Camera, outside Material) Scene {
	return Scene{objs, cam, outside}
}

func (s *Scene) Intersect(ray Ray) (*Object, Intersection) {
	var resObj *Object = nil
	var res Intersection
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

func (s *Scene) TotalAabb() linal.Aabb {
	aabb := linal.Aabb{}

	for _, obj := range s.objs {
		box := obj.Aabb()
		aabb = aabb.Merge(&box)
	}

	return aabb

}
