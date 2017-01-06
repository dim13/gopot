package main

type Vertex struct{ X, Y, Z float64 }

type Patch []Vertex

func (p Patch) X() []float64 {
	f := make([]float64, len(p))
	for i, v := range p {
		f[i] = v.X
	}
	return f
}

func (p Patch) Y() []float64 {
	f := make([]float64, len(p))
	for i, v := range p {
		f[i] = v.Y
	}
	return f
}

func (p Patch) Z() []float64 {
	f := make([]float64, len(p))
	for i, v := range p {
		f[i] = v.Z
	}
	return f
}

func (v Vertex) RotX() Vertex {
	v.Y = 0.5*v.Y + 0.866*v.Z
	v.Z = 0.5*v.Z - 0.866*v.Y
	return v
}

func (v Vertex) Project() (int, int) {
	dist := 1000.0
	zoom := 1000.0
	v = v.RotX()
	v.X *= dist / (2*dist - v.Z)
	v.Y *= dist / (2*dist - v.Z)
	v.X *= zoom
	v.Y *= zoom
	v.X += width / 2
	v.Y += height / 3
	return int(v.X), int(v.Y)
}
