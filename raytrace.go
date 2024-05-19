package main

import (
	"math"
	"math/rand"
	"raytracing/linal"
	"raytracing/materials"
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

	hits := make([]struct {
		material *materials.Material
		uv       linal.Uv
		incoming linal.Vec3
		normal   linal.Vec3
	}, 0)
	bounce := 0
	for {
		step := r.march(ray)
		ray = step.nextRay

		if step.isHit {
			bounce += 1
			hits = append(hits, struct {
				material *materials.Material
				uv       linal.Uv
				incoming linal.Vec3
				normal   linal.Vec3
			}{step.material, step.hitUv, step.incomingDir, step.normal})
		}

		if step.isEnd || bounce > r.bounce {
			break
		}
	}

	lightColor := materials.Color{}
	for i := len(hits) - 1; i >= 0; i-- {
		matColor := (*hits[i].material).Color(hits[i].uv)
		n := hits[i].normal
		l := hits[i].incoming
		dot := -n.Dot(l)
		if false {
			lightColor.R = matColor.R + lightColor.R*matColor.R*dot
			lightColor.G = matColor.G + lightColor.G*matColor.G*dot
			lightColor.B = matColor.B + lightColor.B*matColor.B*dot
		} else {
			lightColor.R = lightColor.R * matColor.R * dot
			lightColor.G = lightColor.G * matColor.G * dot
			lightColor.B = lightColor.B * matColor.B * dot
		}

		if lightColor.R > 1 {
			lightColor.R = 1
		}
		if lightColor.G > 1 {
			lightColor.G = 1
		}
		if lightColor.B > 1 {
			lightColor.B = 1
		}

	}

	r.scene.Cam.EmitPixel(uv, lightColor)

	return true
}

type marchStep struct {
	nextRay     linal.Ray
	bounce      int
	isHit       bool
	material    *materials.Material
	hitUv       linal.Uv
	incomingDir linal.Vec3
	normal      linal.Vec3
	isEnd       bool
}

func (r *Raymarcher) march(ray linal.Ray) marchStep {
	obj, intersection := r.scene.Intersect(ray)

	if obj == nil {
		sp := ray.Start.ToSpherical()
		return marchStep{
			material:    &r.scene.Outside,
			hitUv:       linal.Uv{U: sp.Y / (2 * math.Pi), V: sp.Z / math.Pi},
			isEnd:       true,
			isHit:       true,
			normal:      ray.Dir.Mul(-1),
			incomingDir: ray.Dir,
		}
	}

	material := (*obj).Material()
	normal := (*obj).Normal(intersection.Uv)
	reflectedDir := ray.Dir.Sub(normal.Mul(ray.Dir.Dot(normal) * 2.0))
	hitPos := (*obj).FromUv(intersection.Uv)

	dir := linal.Vec3{}
	ref := (*material).Reflectiveness(intersection.Uv)
	if ref == 1.0 {
		dir = reflectedDir
	} else {
		reflectedAngles := reflectedDir.ToSpherical()
		u := reflectedAngles.Y / (2.0 * math.Pi)
		v := reflectedAngles.Z / math.Pi
		for {
			s := float32(rand.NormFloat64())*(1-ref) + u
			t := float32(rand.NormFloat64())*(1-ref) + v

			dir = linal.Vec3{X: 1.0, Y: s * 2 * math.Pi, Z: t * math.Pi}.FromSpherical()

			if normal.Dot(dir) >= 0.0 {
				break
			}
		}
	}

	newRay := linal.Ray{Start: hitPos.Add(dir.Mul(1e-3)), Dir: dir}
	return marchStep{
		nextRay:     newRay,
		material:    material,
		hitUv:       intersection.Uv,
		isHit:       true,
		incomingDir: ray.Dir,
		normal:      (*obj).Normal(intersection.Uv),
	}
}
