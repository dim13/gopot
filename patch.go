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
