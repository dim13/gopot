package main

type M [4][4]float64
type V [4]float64

func (s Surface) X() (m M) {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			m[x][y] = s[x][y].X
		}
	}
	return
}

func (s Surface) Y() (m M) {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			m[x][y] = s[x][y].Y
		}
	}
	return
}

func (s Surface) Z() (m M) {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			m[x][y] = s[x][y].Z
		}
	}
	return
}

func (m M) T() (r M) {
	for x, l := range m {
		for y, v := range l {
			r[y][x] = v
		}
	}
	return
}

func Power(u float64) (v V) {
	v[3] = 1
	v[2] = u * v[3]
	v[1] = u * v[2]
	v[0] = u * v[1]
	return
}

// TODO
// x(u, v) = U * Mb * Gbx * T(Mb) * T(V)

// Left multiplication
func (v V) LMul(m M) (r M) {
	return
}

// Matrix Multiplication
func (m M) MMul(mm M) (r M) {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
		}
	}
	return
}

// Right multiplication
func (m M) RMul(v V) (r V) {
	return
}

func (v V) Sum() (r float64) {
	for i := range v {
		r += v[i]
	}
	return
}
