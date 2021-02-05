package main

import (
	"io"
	"os"
)

const (
	FF  = 12
	ESC = 27
	FS  = 28 // point plot
	GS  = 29 // graph and dark vector
	RS  = 30 // incremental plot
	US  = 31 // alpha mode
)

type Tek struct {
	io.Writer
	hx, hy, lx, ly, eb byte
	xterm              bool
	height, width      int
}

func (t *Tek) escString(s string) {
	t.Write([]byte{ESC})
	t.Write([]byte(s))
}

func (t *Tek) writeByte(b ...byte) {
	t.Write(b)
}

func NewTek(w io.Writer) *Tek {
	return &Tek{
		Writer: w,
		xterm:  os.Getenv("TERM") == "xterm",
		height: 3072,
		width:  4096,
	}
}

func (t Tek) Clear() {
	t.writeByte(ESC, FF) // Tek Page
}

func (t Tek) Enable() {
	if t.xterm {
		t.escString("[?38h")
	}
	t.Clear()
}

func (t Tek) Disable() {
	t.writeByte(US) // Text mode
	if t.xterm {
		t.writeByte(ESC, 3) // VT Page
	}
}

func (t Tek) Pen() {
	t.writeByte(GS)
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

func (t *Tek) Dim() (w, h int) {
	return t.width, t.height
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

func (t *Tek) Plot(x, y int) {
	x = limit(x, t.width)
	y = limit(y, t.height)

	hy := byte(y>>7) & 0x1f
	eb := (byte(y&3) << 2) | byte(x&3)
	ly := byte(y>>2) & 0x1f
	hx := byte(x>>7) & 0x1f
	lx := byte(x>>2) & 0x1f

	if hy != t.hy {
		t.writeByte(0x20 | hy)
	}
	if eb != t.eb {
		t.writeByte(0x60 | eb)
	}
	if ly != t.ly || eb != t.eb || hx != t.hx {
		t.writeByte(0x60 | ly)
	}
	if hx != t.hx {
		t.writeByte(0x20 | hx)
	}
	t.writeByte(0x40 | lx)

	t.hy = hy
	t.eb = eb
	t.ly = ly
	t.hx = hx
	t.lx = lx
}
