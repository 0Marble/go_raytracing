package lenses

import "raytracing/linal"

type ProjectiveLens struct {
	mat linal.Mat
}

func InitProjectiveLens(t linal.Transform) ProjectiveLens {
	return ProjectiveLens{t.ToMat()}
}

func (l *ProjectiveLens) ShootRay(uv linal.Uv) linal.Ray {
	d, _ := linal.Vec3{X: uv.U*2 - 1, Y: uv.V*2 - 1, Z: 1}.Normalize()
	return l.mat.ApplyToRay(linal.Ray{Dir: d})
}
