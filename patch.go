package main

import "math"

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

func rot(deg int) (float64, float64) {
	rad := float64(deg) * math.Pi / 180.0
	return math.Sin(rad), math.Cos(rad)
}

func (v Vertex) RotX(deg int) Vertex {
	sin, cos := rot(deg)
	v.Y = cos*v.Y - sin*v.Z
	v.Z = sin*v.Y + cos*v.Z
	return v
}

func (v Vertex) RotY(deg int) Vertex {
	sin, cos := rot(deg)
	v.X = cos*v.X + sin*v.Z
	v.Z = -sin*v.X + cos*v.Z
	return v
}

func (v Vertex) RotZ(deg int) Vertex {
	sin, cos := rot(deg)
	v.X = cos*v.X - sin*v.Y
	v.Y = sin*v.X + cos*v.Y
	return v
}

func (v Vertex) Zoom(zoom float64) Vertex {
	v.X *= zoom
	v.Y *= zoom
	v.Z *= zoom
	return v
}

func (v Vertex) Project(p Plotter, rx, ry, rz int) {
	dist := 100000.0
	v = v.Zoom(1000).RotZ(rz).RotY(ry).RotX(rx)
	w, h := p.Dim()

	v.X *= dist / (2*dist - v.Z)
	v.Y *= dist / (2*dist - v.Z)
	v.X += float64(w) / 2
	v.Y += float64(h) / 3

	p.Plot(int(v.X), int(v.Y))
}
