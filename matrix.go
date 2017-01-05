package main

type Matrix []Vector
type Vector []float64

func (M Matrix) Rank() (u, v int) {
	u = len(M)
	if u > 0 {
		v = len(M[0])
	}
	return
}

func (M Matrix) Row(n int) Vector {
	if u, _ := M.Rank(); n > u {
		return nil
	}
	return M[n]
}

func (M Matrix) Column(n int) Vector {
	if _, v := M.Rank(); n > v {
		return nil
	}
	return M.Transpose().Row(n)
}

func (M Matrix) Transpose() Matrix {
	u, v := M.Rank()
	m := NewMatrix(v, u)
	for i, V := range M {
		for j, val := range V {
			m[j][i] = val
		}
	}
	return m
}

func NewMatrix(u, v int) Matrix {
	m := make(Matrix, u)
	for i := range m {
		m[i] = make(Vector, v)
	}
	return m
}

func Rank(m Matrix) (c, l int) {
	c = len(m)
	if len(m) > 0 {
		l = len(m[0])
	}
	return
}

func PowerN(u float64, n int) []float64 {
	if n < 1 {
		return []float64{}
	}
	v := make([]float64, n)
	v[n-1] = 1
	for i := n - 1; i > 0; i-- {
		v[i-1] = u * v[i]
	}
	return v
}

// TODO
// x(u, v) = U * Mb * Gbx * T(Mb) * T(V)
// y(u, v) = U * Mb * Gby * T(Mb) * T(V)
// z(u, v) = U * Mb * Gbz * T(Mb) * T(V)
