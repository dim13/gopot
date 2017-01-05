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

type Bezier []Vertex

func ParseVertex(s string) (v Vertex) {
	ss := strings.Split(s, ",")
	if len(ss) != 3 {
		return
	}
	v.X, _ = strconv.ParseFloat(ss[0], 64)
	v.Y, _ = strconv.ParseFloat(ss[1], 64)
	v.Z, _ = strconv.ParseFloat(ss[2], 64)
	return
}

func ParsePatch(s string) (p []int) {
	f := strings.Split(s, ",")
	for _, v := range f {
		i, _ := strconv.ParseInt(v, 10, 64)
		p = append(p, int(i-1))
	}
	return
}

func Parse(r io.Reader) (b []Bezier) {
	scan := bufio.NewScanner(r)

	if !scan.Scan() {
		return
	}
	n, _ := strconv.Atoi(scan.Text())

	p := make([][]int, n)
	for i := 0; i < n; i++ {
		if !scan.Scan() {
			return
		}
		p[i] = ParsePatch(scan.Text())
	}

	if !scan.Scan() {
		return
	}
	m, _ := strconv.Atoi(scan.Text())

	v := make([]Vertex, m)
	for i := 0; i < m; i++ {
		if !scan.Scan() {
			return
		}
		v[i] = ParseVertex(scan.Text())
	}

	b = make([]Bezier, n)
	for i, patch := range p {
		b[i] = make([]Vertex, n)
		for j, x := range patch {
			b[i][j] = v[x]
		}
	}

	return
}
