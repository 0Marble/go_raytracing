package raytracing

import (
	"math"
	"raytracing/linal"
	"raytracing/materials"
	"raytracing/scene"
)

type SimpleRaytracer struct {
	scene  scene.Scene
	bounce int
}

func InitSimpleRaytracer(scene scene.Scene, bounce int) SimpleRaytracer {
	return SimpleRaytracer{scene, bounce}
}

func (r *SimpleRaytracer) Sample(ray linal.Ray) materials.Color {
	hits := make([]singleStep, 0)
	for bounce := 0; bounce < r.bounce; bounce++ {
		step := r.trace(ray)
		hits = append(hits, step)

		if step.hitOutside {
			break
		}
		ray = step.nextRay
	}

	res := materials.Color{}
	for i := len(hits) - 1; i >= 0; i-- {
		step := hits[i]
		color := materials.Color{}

		cnt := 1
		for _, light := range r.scene.Lights() {
			ray, uv := light.RayToLight(step.hitPos)
			ray.Start = ray.Start.Add(ray.Dir.Mul(1e-3))
			obj, intersection := r.scene.Intersect(ray)
			if obj != nil && intersection.T < light.TDist(ray) && intersection.T > 0 {
				continue
			}

			lightColor := light.GetColor(uv, step.hitPos)
			lit := (*step.material).Color(ray.Dir, step.incoming.Mul(-1), step.normal)
			color = color.Add(lightColor.Mul(lit))
			cnt += 1
		}

		lit := (*step.material).Color(step.nextRay.Dir, step.incoming.Mul(-1), step.normal)
		color = color.Add(res.Mul(lit))

		res = color.DivNum(float32(cnt)).Clamp(materials.Color{}, materials.Color{R: 1, G: 1, B: 1})
	}

	return res
}

type singleStep struct {
	material   *materials.Material
	uv         linal.Uv
	hitPos     linal.Vec3
	incoming   linal.Vec3
	normal     linal.Vec3
	hitOutside bool
	nextRay    linal.Ray
}

func (r *SimpleRaytracer) trace(ray linal.Ray) singleStep {
	obj, intersection := r.scene.Intersect(ray)

	if obj == nil {
		sp := ray.Start.ToSpherical()
		return singleStep{
			hitOutside: true,
			incoming:   ray.Dir,
			normal:     ray.Dir.Mul(-1),
			material:   &r.scene.Outside,
			uv:         linal.Uv{U: sp.Y / (2 * math.Pi), V: sp.Z / math.Pi},
		}
	}

	normal := (*obj).Normal(intersection.Uv)
	nextDir := ray.Dir.Sub(normal.Mul(2 * ray.Dir.Dot(normal)))
	pos := (*obj).FromUv(intersection.Uv)
	return singleStep{
		hitOutside: false,
		incoming:   ray.Dir,
		material:   (*obj).Material(),
		normal:     normal,
		uv:         intersection.Uv,
		hitPos:     pos,
		nextRay:    linal.Ray{Start: pos.Add(nextDir.Mul(1e-3)), Dir: nextDir},
	}
}
