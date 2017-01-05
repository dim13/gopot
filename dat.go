package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Vertex struct {
	X, Y, Z float64
}

type Bezier [][]float64

func parseVertex(s string) []float64 {
	f := strings.Split(s, ",")
	p := make([]float64, len(f))
	for i, v := range f {
		p[i], _ = strconv.ParseFloat(v, 64)
	}
	return p
}

func parsePatch(s string) []int {
	f := strings.Split(s, ",")
	p := make([]int, len(f))
	for i, v := range f {
		x, _ := strconv.ParseInt(v, 10, 64)
		p[i] = int(x - 1)
	}
	return p
}

func Parse(r io.Reader) [][][]float64 {
	scan := bufio.NewScanner(r)

	if !scan.Scan() {
		return nil
	}
	n, _ := strconv.Atoi(scan.Text())

	patches := make([][]int, n)
	for i := range patches {
		if !scan.Scan() {
			return nil
		}
		patches[i] = parsePatch(scan.Text())
	}

	if !scan.Scan() {
		return nil
	}
	m, _ := strconv.Atoi(scan.Text())

	vertices := make([][]float64, m)
	for i := range vertices {
		if !scan.Scan() {
			return nil
		}
		vertices[i] = parseVertex(scan.Text())
	}

	bezier := make([][][]float64, n)
	for u, p := range patches {
		bezier[u] = make([][]float64, len(p))
		for v, i := range p {
			bezier[u][v] = vertices[i]
		}
	}

	return bezier
}
