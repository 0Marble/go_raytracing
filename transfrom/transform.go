package transfrom

import (
	"math"
	"raytracing/linal"
)

type Transform struct {
	Translation linal.Vec3
	Scale       linal.Vec3
	Rotation    linal.Vec3
}

func (t *Transform) ToMat() linal.Mat {
	trans := translationMat(t.Translation)
	scale := scaleMat(t.Scale)
	rotx := rotMatX(t.Rotation.X)
	roty := rotMatX(t.Rotation.Y)
	rotz := rotMatX(t.Rotation.Z)

	a := roty.Matmul(&rotx)
	b := rotz.Matmul(&a)
	c := scale.Matmul(&b)

	return trans.Matmul(&c)
}

func translationMat(trans linal.Vec3) linal.Mat {
	return linal.MatFromVals(4, []float32{0, 0, 0, trans.X, 0, 0, 0, trans.Y, 0, 0, 0, trans.Z, 0, 0, 0, 1})
}

func scaleMat(scale linal.Vec3) linal.Mat {
	return linal.MatFromVals(4, []float32{scale.X, 0, 0, 0, 0, scale.Y, 0, 0, 0, 0, scale.Z, 0})
}
func rotMatX(rot float32) linal.Mat {
	c := float32(math.Cos(float64(rot)))
	s := float32(math.Sin(float64(rot)))
	return linal.MatFromVals(4, []float32{1, 0, 0, 0, 0, c, -s, 0, 0, s, c, 0, 0, 0, 0, 0})
}
func rotMatY(rot float32) linal.Mat {
	c := float32(math.Cos(float64(rot)))
	s := float32(math.Sin(float64(rot)))
	return linal.MatFromVals(4, []float32{c, 0, s, 0, 0, 1, 0, 0, -s, 0, c, 0, 0, 0, 0, 0})
}
func rotMatZ(rot float32) linal.Mat {
	c := float32(math.Cos(float64(rot)))
	s := float32(math.Sin(float64(rot)))
	return linal.MatFromVals(4, []float32{c, -s, 0, 0, s, c, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0})
}
