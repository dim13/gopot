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

func steps(n int) []float64 {
	st := make([]float64, n+1)
	for i := 0; i <= n; i++ {
		st[i] = float64(i) * 1.0 / float64(n)
	}
	return st
}

func plot(out Plotter, patches []Patch) {
	st := steps(5)
	for _, p := range patches {
		for _, u := range st {
			out.Pen()
			for _, v := range st {
				p.Calc(u, v).Project(out, -60, 0, -15)
			}
		}
		for _, u := range st {
			out.Pen()
			for _, v := range st {
				p.Calc(v, u).Project(out, -60, 0, -15)
			}
		}
	}
}
