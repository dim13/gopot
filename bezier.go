package main

import "github.com/gonum/matrix/mat64"

var mb = []float64{-1, 3, -3, 1, 3, -6, 3, 0, -3, 3, 0, 0, 1, 0, 0, 0}

func newVector(x float64) []float64 {
	v := make([]float64, 4)
	v[3] = 1
	v[2] = x * v[3]
	v[1] = x * v[2]
	v[0] = x * v[1]
	return v
}

func Vector(x float64) *mat64.Vector {
	return mat64.NewVector(4, newVector(x))
}

func Matrix(m []float64) *mat64.Dense {
	return mat64.NewDense(4, 4, m)
}

// x(u, v) = U * Mb * Gbx * T(Mb) * T(V)
// y(u, v) = U * Mb * Gby * T(Mb) * T(V)
// z(u, v) = U * Mb * Gbz * T(Mb) * T(V)

func mult(U, V *mat64.Vector, Mb, Gb *mat64.Dense) float64 {
	m1 := new(mat64.Dense)
	m2 := new(mat64.Dense)
	m1.Mul(Mb, Gb)
	m2.Mul(m1, Mb)
	return mat64.Inner(U, m2, V)
}

func Calc(u, v float64, p Patch) (r Vertex) {
	U := Vector(u)
	V := Vector(v)
	Mb := Matrix(mb)
	Gbx := Matrix(p.X())
	Gby := Matrix(p.Y())
	Gbz := Matrix(p.Z())
	r.X = mult(U, V, Mb, Gbx)
	r.Y = mult(U, V, Mb, Gby)
	r.Z = mult(U, V, Mb, Gbz)
	return
}
