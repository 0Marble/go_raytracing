package scene

import "raytracing/linal"

type Scene struct {
	objs    []Object
	Cam     Camera
	Outside Material
}

func (s *Scene) MaxDist(ray Ray) (*Object, float32) {
	max := float32(0.0)
	var resObj *Object = nil

	for i, obj := range s.objs {
		dist := obj.Distance(ray.Start)
		if dist > ray.Step {
			continue
		}

		if dist > max {
			max = dist
			resObj = &s.objs[i]
		}
	}

	return resObj, max
}

func (s *Scene) TotalAabb() linal.Aabb {
	aabb := linal.Aabb{}

	for _, obj := range s.objs {
		box := obj.Aabb()
		aabb = aabb.Merge(&box)
	}

	return aabb

}
