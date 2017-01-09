package main

type Plotter interface {
	Clear()
	Dim() (width, height int)
	Plot(x, y int)
	Pen()
}
