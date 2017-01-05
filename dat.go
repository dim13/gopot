package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func parseVertex(s string) (v Vertex) {
	f := strings.Split(s, ",")
	if len(f) != 3 {
		panic("invalid vertex")
	}
	v.X, _ = strconv.ParseFloat(f[0], 64)
	v.Y, _ = strconv.ParseFloat(f[1], 64)
	v.Z, _ = strconv.ParseFloat(f[2], 64)
	return v
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

func Parse(r io.Reader) []Patch {
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

	vertices := make([]Vertex, m)
	for i := range vertices {
		if !scan.Scan() {
			return nil
		}
		vertices[i] = parseVertex(scan.Text())
	}

	// populate patches with vertices
	patches := make([]Patch, n)
	for u, i := range indices {
		patches[u] = make([]Vertex, len(i))
		for v, j := range i {
			patches[u][v] = vertices[j]
		}
	}

	return patches
}
