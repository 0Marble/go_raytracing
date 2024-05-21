package linal

import (
	"log"
)

type Transform struct {
	Translation Vec3
	Scale       Vec3
	Rotation    Quat
}

func (t *Transform) ToMat() Mat {
	trans := translationMat(t.Translation)
	scale := scaleMat(t.Scale)
	rot, ok := t.Rotation.ToMat()
	if !ok {
		log.Fatal("Quat does not represent a rotation", t.Rotation)
	}

	a := rot.Matmul(&scale)
	return trans.Matmul(&a)
}

func translationMat(trans Vec3) Mat {
	return MatFromVals(4, []float32{1, 0, 0, trans.X, 0, 1, 0, trans.Y, 0, 0, 1, trans.Z, 0, 0, 0, 1})
}

func scaleMat(scale Vec3) Mat {
	return MatFromVals(4, []float32{scale.X, 0, 0, 0, 0, scale.Y, 0, 0, 0, 0, scale.Z, 0, 0, 0, 0, 1})
}
