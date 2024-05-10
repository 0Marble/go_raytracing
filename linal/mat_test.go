package linal

import (
	"math"
	"testing"
)

func almostEqual(a Mat, b Mat, t *testing.T) {
	t.Logf("almostEqual %v\n%v", a, b)
	if a.Size != b.Size {
		t.Fail()
	}

	for row := 0; row < a.Size; row++ {
		for col := 0; col < a.Size; col++ {
			if math.Abs(float64(a.Get(row, col)-b.Get(row, col))) > 1e-3 {
				t.Fail()
			}
		}
	}
}

func TestMatmul1(t *testing.T) {
	a := MatFromVals(2, []float32{0.0, 1.0, 0.0, 0.0})
	b := MatFromVals(2, []float32{0.0, 0.0, 1.0, 0.0})
	c := MatFromVals(2, []float32{1.0, 0.0, 0.0, 0.0})

	almostEqual(a.Matmul(&b), c, t)
}
func TestMatmul2(t *testing.T) {
	a := MatFromVals(3, []float32{1.0, 0.0, 1.0, 2.0, 1.0, 1.0, 0.0, 1.0, 1.0})
	b := MatFromVals(3, []float32{1.0, 2.0, 1.0, 2.0, 3.0, 1.0, 4.0, 2.0, 2.0})
	c := MatFromVals(3, []float32{5.0, 4.0, 3.0, 8.0, 9.0, 5.0, 6.0, 5.0, 3.0})

	almostEqual(a.Matmul(&b), c, t)
}

func testInverseSucc(t *testing.T) {
	for i, mat := range []Mat{
		MatFromVals(2, []float32{-1.0, 1.5, 1.0, -1.0}),
		MatFromVals(3, []float32{4, 3, 8, 6, 2, 5, 1, 5, 9}),
		MatFromVals(3, []float32{2, 3, 1, 1, 1, 2, 2, 3, 4}),
		MatFromVals(3, []float32{6, 2, 3, 0, 0, 4, 2, 0, 0}),
		MatFromVals(3, []float32{1, 2, 3, 0, 1, 4, 0, 0, 1}),
		MatFromVals(3, []float32{1, 2, 3, 2, 1, 4, 3, 4, 1}),
	} {
		t.Logf("cur matrix: %v\n", i)
		inv, ok := mat.Inverse()
		if !ok {
			t.Fail()
		}
		almostEqual(mat.Matmul(&inv), MatIdent(mat.Size), t)
	}
}

func TestInverseFail(t *testing.T) {
	for i, mat := range []Mat{
		MatFromVals(3, []float32{1, 2, 3, 4, 5, 6, 7, 8, 9}),
	} {

		t.Logf("cur matrix: %v\n", i)
		_, ok := mat.Inverse()
		if ok {
			t.Fail()
		}
	}
}
