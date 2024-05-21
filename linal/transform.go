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

func (t Transform) ToStaticTransform() *StaticTransform {
	mat := t.ToMat()
	inv, ok := mat.Inverse()
	if !ok {
		log.Fatal("Could not invert matrix: ", mat, t)
	}

	return &StaticTransform{mat, inv}
}

func translationMat(trans Vec3) Mat {
	return MatFromVals(4, []float32{1, 0, 0, trans.X, 0, 1, 0, trans.Y, 0, 0, 1, trans.Z, 0, 0, 0, 1})
}

func scaleMat(scale Vec3) Mat {
	return MatFromVals(4, []float32{scale.X, 0, 0, 0, 0, scale.Y, 0, 0, 0, 0, scale.Z, 0, 0, 0, 0, 1})
}

type StaticTransform struct {
	mat Mat
	inv Mat
}

func (s *StaticTransform) ToMat(time float32) Mat {
	return s.mat
}
func (s *StaticTransform) ToInv(time float32) Mat {
	return s.inv
}

type TimedTransform interface {
	ToMat(time float32) Mat
	ToInv(time float32) Mat
}

type InterpolatedTransform struct {
	Start Transform
	End   Transform
}

func (tf *InterpolatedTransform) Lerp(t float32) Transform {
	trans := tf.Start.Translation.Lerp(tf.End.Translation, t)
	scale := tf.Start.Scale.Lerp(tf.End.Scale, t)
	rot := tf.Start.Rotation.Slerp(tf.End.Rotation, t)
	return Transform{trans, scale, rot}
}

func (tf *InterpolatedTransform) ToMat(time float32) Mat {
	t := tf.Lerp(time)
	mat := t.ToMat()
	return mat
}
func (tf *InterpolatedTransform) ToInv(time float32) Mat {
	t := tf.Lerp(time)
	mat := t.ToMat()
	inv, ok := mat.Inverse()
	if !ok {
		log.Fatal("Could not invert matrix: ", mat, time, tf)
	}
	return inv
}
