package raymarching

import (
	"math"
	"math/rand"
	"raytracing/linal"
	"raytracing/scene"
)

type Raymarcher struct {
	scene  scene.Scene
	bounce int
}

func (r *Raymarcher) March() bool {
	ray, uv, end := r.scene.Cam.ShootRay()
	if end {
		return true
	}

	color := r.march(ray, r.bounce)
	r.scene.Cam.EmitPixel(uv, color)

	return false
}

func (r *Raymarcher) march(ray scene.Ray, bounce int) scene.Color {
	if bounce == -1 {
		return scene.Color{}
	}

	obj, dist := r.scene.MaxDist(ray)
	if obj == nil {
		aabb := r.scene.TotalAabb()
		if !aabb.ContainsPoint(ray.Start) {
			sp := ray.Start.ToSpherical()
			return r.scene.Outside.Color(scene.Uv{U: sp.Y, V: sp.Z})
		}
		return r.march(ray.AdvanceBy(ray.Step), bounce)
	}
	if dist > 0.0 {
		return r.march(ray.AdvanceBy(ray.Step), bounce)
	}

	material := (*obj).Material()
	uv := (*obj).ToUv(ray.Start)
	c1 := (*material).Color(uv)

	normal := (*obj).Normal(uv)
	reflectedDir := ray.Dir.Sub(normal.Mul(ray.Dir.Dot(normal) * 2.0))

	dir := linal.Vec3{}
	if (*material).Reflectiveness(uv) == 1.0 {
		dir = reflectedDir
	} else {
		reflectedAngles := reflectedDir.ToSpherical()
		theta := float32(rand.NormFloat64()*float64(1.0-(*material).Reflectiveness(uv)) + float64(reflectedAngles.Y))
		phi := float32(rand.NormFloat64()*float64(1.0-(*material).Reflectiveness(uv)) + float64(reflectedAngles.Z))
		theta = float32(math.Remainder(float64(theta), 2.0*math.Pi))
		phi = float32(math.Remainder(float64(phi), math.Pi))
		dir = linal.Vec3{X: 1.0, Y: theta, Z: phi}.FromSpherical()

		if normal.Dot(dir) < 0.0 {
			dir = dir.Mul(-1)
		}
	}

	c2 := r.march(scene.Ray{Dir: dir, Start: ray.Start, Step: ray.Step}, bounce-1)
	res := c1
	if (*material).EmitsLight() {
		res.R += c2.R
		res.G += c2.G
		res.B += c2.B
	} else {
		res.R *= c2.R
		res.G *= c2.G
		res.B *= c2.B
	}

	return res
}
