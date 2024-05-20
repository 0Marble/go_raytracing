package lenses

import (
	"math"
	"raytracing/linal"
)

type SphericalLens struct {
	mat linal.Mat
}

func InitSphericalLens(t linal.Transform) SphericalLens {
	return SphericalLens{t.ToMat()}
}

func (l *SphericalLens) ShootRay(uv linal.Uv) linal.Ray {
	d := linal.Vec3{X: 1, Y: uv.U * 2 * math.Pi, Z: uv.V * math.Pi}.FromSpherical()
	return l.mat.ApplyToRay(linal.Ray{Dir: d})
}
