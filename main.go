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
	out.Clear()
	out.Enable()
	defer out.Disable()
	step := 1.0 / float64(5)
	for _, p := range patches {
		for u := 0.0; u <= 1.0; u += step {
			out.PenDown()
			for v := 0.0; v <= 1.0; v += step {
				vertex := Calc(u, v, p)
				x, y := vertex.Project()
				out.Draw(x, y)
			}
			out.PenUp()

			out.PenDown()
			for v := 0.0; v <= 1.0; v += step {
				vertex := Calc(v, u, p)
				x, y := vertex.Project()
				out.Draw(x, y)
			}
			out.PenUp()
		}
	}
}
