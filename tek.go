package main

import (
	"io"
	"os"
)

type Out struct {
	io.Writer
	hx, hy, lx, ly, eb byte
	xterm              bool
	height, width      int
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
		height: 3072,
		width:  4096,
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
		//o.escString("[?38l")
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

func (o *Out) Dim() (w, h int) {
	return o.width, o.height
}

// Table 13-4 Bytes Values for Encoding Coordinates
// 		Tag Bits	Address Bits
// Byte Name 	7	6	5	4	3	2	1
// High Y	0	1	5 most significant bits of Y address
// Extra	1	1		Y2	Y1	X2	X1
// Low Y	1	1	5 intermediate bits of Y address
// High X	0	1	5 most significant bits of X address
// Low X	1	0	5 intermediate bits of X address

// Table 13-5 Rules for Sending Short Address
// Bytes Changed	Bytes Sent
// 			High Y	Extra	Low Y	High X	Low X
// High Y		Yes	No	No	No	Yes
// Extra		No	Yes	Yes	No	Yes
// Low Y		No	No	Yes	No	Yes
// High X		No	No	Yes	Yes	Yes
// Low X		No	No	No	No	Yes

// Ref: http://www.vt100.net/docs/vt3xx-gp/chapter13.html

func (o *Out) Plot(x, y int) {
	x = limit(x, o.width)
	y = limit(y, o.height)

	hy := byte(y>>7) & 0x1f
	eb := (byte(y&3) << 2) | byte(x&3)
	ly := byte(y>>2) & 0x1f
	hx := byte(x>>7) & 0x1f
	lx := byte(x>>2) & 0x1f

	if hy != o.hy {
		o.writeByte(0x20 | hy)
	}
	if eb != o.eb {
		o.writeByte(0x60 | eb)
	}
	if ly != o.ly || eb != o.eb || hx != o.hx {
		o.writeByte(0x60 | ly)
	}
	if hx != o.hx {
		o.writeByte(0x20 | hx)
	}
	o.writeByte(0x40 | lx)

	o.hy = hy
	o.eb = eb
	o.ly = ly
	o.hx = hx
	o.lx = lx
}
