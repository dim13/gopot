package main

import (
	"bufio"
	"io"
)

const (
	height = 3072
	width  = 4096
)

type Out struct {
	*bufio.Writer
	hix, hiy, lox, loy, eb int
}

func NewOut(w io.Writer) *Out {
	return &Out{Writer: bufio.NewWriter(w)}
}

func (o Out) Enable() {
	o.WriteByte(0x1b)
	o.WriteString("[?38h")
	o.Flush()
}

func (o Out) Disable() {
	o.WriteByte(0x1b)
	o.WriteByte(0x03)
	o.Flush()
}

func (o Out) Clear() {
	o.Enable()
	o.WriteByte(0x1b)
	o.WriteByte(0x0c)
	o.Disable()
	o.Flush()
}

func (o Out) PenUp() {
	o.WriteByte(0x1d)
	o.WriteByte(0x07)
	o.Flush()
}

func (o Out) PenDown() {
	o.WriteByte(0x1d)
	o.Flush()
}

func limit(val, max int) int {
	if val < 0 {
		return 0
	}
	if val >= max {
		return max - 1
	}
	return val
}

func (o *Out) Plot(x, y int) {
	x = limit(x, width)
	y = limit(y, height)

	hix := (x >> 7) & 0x1f
	hiy := (y >> 7) & 0x1f
	lox := (x >> 2) & 0x1f
	loy := (y >> 2) & 0x1f
	eb := (x & 3) | ((y & 3) << 2)

	if hiy != o.hiy {
		o.WriteByte(byte(hiy | 0x20))
	}
	if eb != o.eb {
		o.WriteByte(byte(eb | 0x60))
	}
	if eb != o.eb || loy != o.loy || hix != o.hix {
		o.WriteByte(byte(loy | 0x60))
	}
	if hix != o.hix {
		o.WriteByte(byte(hix | 0x20))
	}
	o.WriteByte(byte(lox | 0x40))
	o.hix = hix
	o.hiy = hiy
	o.lox = lox
	o.loy = loy
	o.eb = eb
	o.Flush()
}
