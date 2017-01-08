package main

import (
	"io"
	"os"
)

const (
	height = 3072
	width  = 4096
)

type Out struct {
	io.Writer
	hix, hiy, lox, loy, eb byte
	xterm                  bool
}

func (o *Out) escString(s string) {
	o.Write([]byte{27})
	o.Write([]byte(s))
}

func (o *Out) writeByte(b ...byte) {
	o.Write(b)
}

func NewOut(w io.Writer) *Out {
	return &Out{
		Writer: w,
		xterm:  os.Getenv("TERM") == "xterm",
	}
}

func (o Out) Enable() {
	if o.xterm {
		o.escString("[?38h")
		o.writeByte(27, 12) // Tek Page
	}
}

func (o Out) Disable() {
	if o.xterm {
		o.writeByte(31)    // Text mode
		o.writeByte(27, 3) // VT Page
	}
}

func (o Out) PenUp() {
	o.writeByte(29, 7)
}

func (o Out) PenDown() {
	o.writeByte(29)
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

	hiy := byte(y>>7) & 0x1f
	loy := byte(y>>2) & 0x1f
	hix := byte(x>>7) & 0x1f
	lox := byte(x>>2) & 0x1f
	eb := byte(x&3) | (byte(y&3) << 2)

	if hiy != o.hiy {
		o.writeByte(hiy | 0x20)
	}
	if eb != o.eb {
		o.writeByte(eb | 0x60)
	}
	if eb != o.eb || loy != o.loy || hix != o.hix {
		o.writeByte(loy | 0x60)
	}
	if hix != o.hix {
		o.writeByte(hix | 0x20)
	}
	o.writeByte(lox | 0x40)
	o.hix = hix
	o.hiy = hiy
	o.lox = lox
	o.loy = loy
	o.eb = eb
}
