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

func (s *Scene) MinDist(ray Ray) (*Object, float32) {
	min := float32(ray.Step)
	var resObj *Object = nil

	for i, obj := range s.objs {
		dist := obj.Distance(ray.Start)
		if dist > ray.Step {
			continue
		}

		if dist < min || resObj == nil {
			min = dist
			resObj = &s.objs[i]
		}
	}
	return resObj, min
}

func (s *Scene) TotalAabb() linal.Aabb {
	aabb := linal.Aabb{}

	for _, obj := range s.objs {
		box := obj.Aabb()
		aabb = aabb.Merge(&box)
	}

	return aabb

}
