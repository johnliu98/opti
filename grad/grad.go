package grad

import (
	"gonum.org/v1/gonum/diff/fd"
	"gonum.org/v1/gonum/mat"
)

func Gradient(dst *mat.Dense, f func([]float64) []float64, x []float64, settings *fd.Settings) {
	in := len(x)
	out := len(f(x))

	if dst.IsEmpty() {
        *dst = *(dst.Grow(in, out).(*mat.Dense))
	} else if i, o := dst.Dims(); i != in || o != out {
		panic("grad: matrix dimension mismatch")
	}
    dst.Zero()

	if settings == nil {
		settings = &fd.Settings{Formula: fd.Central}
	}

	for j := 0; j < out; j++ {
		fj := func(x []float64) float64 { return f(x)[j] }
		dst.SetCol(j, fd.Gradient(nil, fj, x, settings))
	}
}
