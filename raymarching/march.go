package raymarching

import (
	"log"
	"math"
	"math/rand"
	"raytracing/linal"
	"raytracing/scene"
)

type Raymarcher struct {
	scene  scene.Scene
	bounce int
}

func InitRaymarcher(scene scene.Scene, bounce int) Raymarcher {
	return Raymarcher{scene, bounce}
}

func (r *Raymarcher) March() bool {
	ray, uv, notEnd := r.scene.Cam.ShootRay()
	if !notEnd {
		return false
	}
	// log.Println("Emitted ray ", ray)

	hits := make([]struct {
		material *scene.Material
		uv       scene.Uv
	}, 0)
	bounce := r.bounce
	for {
		step := r.march(ray, bounce)
		ray = step.nextRay
		bounce = step.bounce
		// log.Println("\t bounce=", bounce, "ray=", ray)

		if step.isHit {
			hits = append(hits, struct {
				material *scene.Material
				uv       scene.Uv
			}{step.material, step.hitUv})
		}

		if step.isEnd || step.bounce < 0 {
			break
		}
	}

	color := scene.Color{}
	for i := len(hits) - 1; i >= 0; i-- {
		c := (*hits[i].material).Color(hits[i].uv)
		if (*hits[i].material).EmitsLight() {
			color.R += c.R
			color.G += c.G
			color.B += c.B
		} else {
			color.R *= c.R
			color.G *= c.G
			color.B *= c.B
		}

		if color.R > 1 {
			color.R = 1
		}
		if color.G > 1 {
			color.G = 1
		}
		if color.B > 1 {
			color.B = 1
		}
	}

	r.scene.Cam.EmitPixel(uv, color)

	return true
}

type marchStep struct {
	nextRay  scene.Ray
	bounce   int
	isHit    bool
	material *scene.Material
	hitUv    scene.Uv
	isEnd    bool
}

func (r *Raymarcher) march(ray scene.Ray, bounce int) marchStep {
	for {
		obj, dist := r.scene.MinDist(ray)

		if obj == nil {
			aabb := r.scene.TotalAabb()
			if !aabb.ContainsPoint(ray.Start) {
				sp := ray.Start.ToSpherical()
				return marchStep{
					material: &r.scene.Outside,
					hitUv:    scene.Uv{U: sp.Y / (2 * math.Pi), V: sp.Z / math.Pi},
					isEnd:    true,
					isHit:    true,
				}
			}
			ray = ray.AdvanceBy(ray.Step)
			continue
		}
		if dist > 1e-5 {
			ray = ray.AdvanceBy(dist)
			continue
		}

		material := (*obj).Material()
		uv := (*obj).ToUv(ray.Start)
		normal := (*obj).Normal(uv)
		reflectedDir := ray.Dir.Sub(normal.Mul(ray.Dir.Dot(normal) * 2.0))

		dir := linal.Vec3{}
		ref := (*material).Reflectiveness(uv)
		if ref == 1.0 {
			dir = reflectedDir
		} else {
			reflectedAngles := reflectedDir.ToSpherical()
			u := reflectedAngles.Y / (2.0 * math.Pi)
			v := reflectedAngles.Z / math.Pi
			s := float32(rand.NormFloat64())*(1-ref) + u
			t := float32(rand.NormFloat64())*(1-ref) + v

			dir = linal.Vec3{X: 1.0, Y: s * 2 * math.Pi, Z: t * math.Pi}.FromSpherical()

			if normal.Dot(dir) < 0.0 {
				dir = dir.Mul(-1)
			}
		}
		if dir.LenSquared() == 0.0 {
			log.Fatal(dir)
		}

		newRay := scene.Ray{Start: ray.Start, Dir: dir, Step: ray.Step}
		return marchStep{nextRay: newRay.AdvanceBy(1e-3), bounce: bounce - 1, material: material, hitUv: uv, isHit: true}
	}
}
