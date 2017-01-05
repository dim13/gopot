package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Vertex struct{ X, Y, Z float64 }
type Patch []Vertex

func parseVertex(s string) []float64 {
	f := strings.Split(s, ",")
	p := make([]float64, len(f))
	for i, v := range f {
		p[i], _ = strconv.ParseFloat(v, 64)
	}
	return p
}

func parseIndex(s string) []int {
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

	// first part, patch indices
	if !scan.Scan() {
		return nil
	}
	n, _ := strconv.Atoi(scan.Text())

	indices := make([][]int, n)
	for i := range indices {
		if !scan.Scan() {
			return nil
		}
		indices[i] = parseIndex(scan.Text())
	}

	// second part, vertices
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

	// populate patches with vertices
	patches := make([][][]float64, n)
	for u, i := range indices {
		patches[u] = make([][]float64, len(i))
		for v, x := range i {
			patches[u][v] = vertices[x]
		}
	}

	return patches
}
