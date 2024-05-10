package linal

import (
	"fmt"
	"math"
)

type Mat struct {
	Size int
	vals []float32
}

func MatIdent(size int) Mat {
	vals := make([]float32, size*size)

	for i := 0; i < size; i++ {
		vals[i*size+i] = 1.0
	}

	return Mat{size, vals}
}

func MatZeros(size int) Mat {
	vals := make([]float32, size*size)
	return Mat{size, vals}
}

func MatFromVals(size int, vals []float32) Mat {
	return Mat{size, vals}
}

func (mat *Mat) index(row int, col int) int {
	return row*mat.Size + col
}
func (mat *Mat) Set(row int, col int, val float32) {
	mat.vals[mat.index(row, col)] = val
}
func (mat *Mat) Get(row int, col int) float32 {
	return mat.vals[mat.index(row, col)]
}

func (a *Mat) Matmul(b *Mat) Mat {
	c := MatZeros(a.Size)

	for i := 0; i < a.Size; i++ {
		for j := 0; j < a.Size; j++ {
			for k := 0; k < a.Size; k++ {
				c.vals[c.index(i, j)] += a.vals[a.index(i, k)] * b.vals[b.index(k, j)]
			}
		}
	}

	return c
}

func (a *Mat) Mul(t float32) Mat {
	res := MatZeros(a.Size)
	for row := 0; row < a.Size; row++ {
		for col := 0; col < a.Size; col++ {
			res.Set(row, col, a.Get(row, col)*t)
		}
	}
	return res
}

func (mat *Mat) Transpose() Mat {
	res := MatZeros(mat.Size)

	for i := 0; i < mat.Size; i++ {
		for j := 0; j < mat.Size; j++ {
			res.Set(j, i, mat.Get(i, j))
		}
	}

	return res
}

func (a *Mat) Inverse() (Mat, bool) {
	mat := MatZeros(a.Size)
	copy(mat.vals, a.vals)
	res := MatIdent(mat.Size)

	// upper-triangular
	h := 0
	k := 0
	for h < mat.Size && k < mat.Size {
		imax := h
		for i := h; i < mat.Size; i++ {
			if math.Abs(float64(mat.Get(i, k))) > math.Abs(float64(mat.Get(imax, k))) {
				imax = i
			}
		}

		if mat.Get(imax, k) == 0.0 {
			k++
		} else {
			mat.swapRows(h, imax)
			res.swapRows(h, imax)

			for i := h + 1; i < mat.Size; i++ {
				t := mat.Get(i, k) / mat.Get(h, k)
				mat.addRow(h, i, -t)
				res.addRow(h, i, -t)
			}
			k++
			h++
		}
	}

	// lower-triangular
	h = mat.Size - 1
	k = mat.Size - 1
	for h >= 0 && k >= 0 {
		if mat.Get(h, k) == 0.0 {
			k--
		} else {
			for i := 0; i < h; i++ {
				t := mat.Get(i, k) / mat.Get(h, k)
				mat.addRow(h, i, -t)
				res.addRow(h, i, -t)
			}
			k--
			h--
		}
	}

	for i := 0; i < mat.Size; i++ {
		if mat.Get(i, i) == 0.0 {
			return res, false
		}
		t := 1.0 / mat.Get(i, i)
		mat.mulRow(i, t)
		res.mulRow(i, t)
	}

	return res, true
}

func (mat *Mat) addRow(from int, to int, t float32) {
	for i := 0; i < mat.Size; i++ {
		mat.vals[mat.index(to, i)] += t * mat.vals[mat.index(from, i)]
	}
}

func (mat *Mat) mulRow(row int, t float32) {
	for i := 0; i < mat.Size; i++ {
		mat.vals[mat.index(row, i)] *= t
	}
}

func (mat *Mat) swapRows(a int, b int) {
	for i := 0; i < mat.Size; i++ {
		t := mat.vals[mat.index(a, i)]
		mat.vals[mat.index(a, i)] = mat.vals[mat.index(b, i)]
		mat.vals[mat.index(b, i)] = t
	}
}

func (mat *Mat) String() string {
	res := ""

	for row := 0; row < mat.Size; row++ {
		res += "|"
		for col := 0; col < mat.Size; col++ {
			res += fmt.Sprintf("%4.4f ", mat.Get(row, col))
		}
		res += "|\n"
	}

	return res
}
