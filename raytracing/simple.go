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
		for _, light := range r.scene.Lights() {
			ray, uv := light.RayToLight(step.hitPos)
			ray.Start = ray.Start.Add(ray.Dir.Mul(1e-3))
			obj, intersection := r.scene.Intersect(ray)
			if obj != nil && intersection.T < light.TDist(ray) && intersection.T > 0 {
				continue
			}

			lightColor := light.GetColor(uv, step.hitPos)
			lit := (*step.material).Lit(step.incoming, step.normal, ray.Dir)

			color.R += step.normal.Dot(ray.Dir) * lightColor.R * lit.R
			color.G += step.normal.Dot(ray.Dir) * lightColor.G * lit.G
			color.B += step.normal.Dot(ray.Dir) * lightColor.B * lit.B
		}

		lit := (*step.material).Lit(step.incoming, step.normal, step.nextRay.Dir)
		color.R += step.normal.Dot(step.nextRay.Dir) * res.R * lit.R
		color.G += step.normal.Dot(step.nextRay.Dir) * res.G * lit.G
		color.B += step.normal.Dot(step.nextRay.Dir) * res.B * lit.B

		res = color.Clamp(materials.Color{}, materials.Color{R: 1, G: 1, B: 1})
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

	mat := (*obj).Material()
	normal := (*obj).Normal(intersection.Uv)
	nextDir := (*mat).Reflect(ray.Dir, normal)
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
