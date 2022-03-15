package canvas

import (
	"io/ioutil"
	"log"
)

type FZX struct {
	b []byte
}

func ReadFZX(filename string) *FZX {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("ReadFZX: cannot ReadFile(%q): %v", filename, err)
	}
	return &FZX{b}
}

func (f *FZX) Height() int   { return int(f.b[0]) }
func (f *FZX) Tracking() int { return int(f.b[1]) }
func (f *FZX) LastChar() int { return int(f.b[2]) }
func (f *FZX) CharIndex(ch int) int {
	assert(ch >= 32)
	assert(ch <= f.LastChar())
	return 3 + 3*(ch-32)
}
func (f *FZX) CharOffset(ch int) int {
	i := f.CharIndex(ch)
	return i + int(f.b[i]) + 256*int(f.b[i+1]&0x3F)
}
func (f *FZX) CharLen(ch int) int {
	i := f.CharIndex(ch)
	j := 6 + 3*(ch-32)
	x := i + int(f.b[i]) + 256*int(f.b[i+1]&0x3F)
	y := j + int(f.b[j]) + 256*int(f.b[j+1]&0x3F)
	return y - x
}
func (f *FZX) CharKern(ch int) int {
	i := f.CharIndex(ch)
	return int((f.b[i+1] & 0xC0) >> 6)
}
func (f *FZX) CharShift(ch int) int {
	i := f.CharIndex(ch)
	return int((f.b[i+2] & 0xF0) >> 4)
}
func (f *FZX) CharWidth(ch int) int {
	i := f.CharIndex(ch)
	return int(f.b[i+2] & 0x0F)
}
func (f *FZX) Demo(ch int) {
	off, n, krn, sh, wid := f.CharOffset(ch), f.CharLen(ch), f.CharKern(ch), f.CharShift(ch), f.CharWidth(ch)
	x := int('.')
	if ch < 128 {
		x = ch
	}
	log.Printf("[%d]=`%c` off=%d len=%d kern=%d shift=%d wid=%d", ch, x, off, n, krn, sh, wid)
}

func (o *Canvas) Print(cx, cy int, f *FZX, c Color, s string) (int, int) {
	hei, trk := f.Height(), f.Tracking()
	_ = hei
	for i := 0; i < len(s); i++ {
		ch := int(s[i])
		off, n, krn, sh, wid := f.CharOffset(ch), f.CharLen(ch), f.CharKern(ch), f.CharShift(ch), f.CharWidth(ch)

		for x := 0; x <= wid; x++ {
			for y := 0; y < n; y++ {
				if (f.b[off+y]>>(7-(x&7)))&1 != 0 {
					o.Set(cx+x-krn, cy-y-sh, c)
				}
			}
			if x == 7 {
				off++
			}
		}
		cx += trk + wid + 1
	}
	return cx, cy
}

func assert(x bool) {
	if !x {
		panic("assert failed")
	}
}
