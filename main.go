package main

import (
	"log"
	"os"
)

func main() {
	fd, err := os.Open("teapot.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	patches := Parse(fd)

	out := NewOut(os.Stdout)
	out.Enable()
	defer out.Disable()
	plot(out, patches)
}

func plot(out Plotter, patches []Patch) {
	out.Clear()
	step := 1.0 / float64(5)
	for _, p := range patches {
		for u := 0.0; u <= 1.0; u += step {
			out.Pen()
			for v := 0.0; v <= 1.0; v += step {
				p.Calc(u, v).Project(out, -60, 0, -15)
			}
		}
		for u := 0.0; u <= 1.0; u += step {
			out.Pen()
			for v := 0.0; v <= 1.0; v += step {
				p.Calc(v, u).Project(out, -60, 0, -15)
			}
		}
	}
}
