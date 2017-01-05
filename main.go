package main

import (
	"fmt"
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
	for _, p := range patches {
		for u := 0.0; u <= 1.0; u += 1.0 {
			for v := 0.0; v <= 1.0; v += 1.0 {
				fmt.Print(Calc(u, v, p))
			}
		}
		fmt.Println("")
	}
}
