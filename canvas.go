/*
	Canvas is rectangular pixmap.
	It starts out black and opaque.
	You can set points.
	You can file it out as a png.
*/
package canvas

import (
	. "fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
)

var Black = RGB(0, 0, 0)
var Red = RGB(255, 0, 0)
var Green = RGB(0, 255, 0)
var Blue = RGB(0, 0, 255)
var White = RGB(255, 255, 255)

type Color struct {
	N color.NRGBA
}

type Canvas struct {
	Pixels          *image.NRGBA
	Width, Height   int
	FWidth, FHeight float64
}

func RGB(r, g, b byte) Color {
	return Color{
		N: color.NRGBA{R: r, G: g, B: b, A: 255},
	}
}

func NewCanvas(w, h int) *Canvas {
	z := &Canvas{
		Pixels: image.NewNRGBA(image.Rect(0, 0, w, h)),
		Width:  w, Height: h,
		FWidth: float64(w), FHeight: float64(h),
	}
	z.Fill(0, 0, w, h, RGB(0, 0, 0))
	return z
}

func (o Canvas) Dup() *Canvas {
	w, h := o.Width, o.Height
	bounds := image.Rect(0, 0, w, h)
	z := &Canvas{
		Pixels: image.NewNRGBA(bounds),
		Width:  w, Height: h,
		FWidth: float64(w), FHeight: float64(h),
	}
	draw.Draw(z.Pixels, bounds, o.Pixels, image.ZP, draw.Src)
	return z
}

// FSet takes float64 x and y, which are visible on the unit square [0, 1) x [0, 1).
// Might generalize from Unit Square later.
func (o Canvas) FSet(x, y float64, c Color) {
	o.Set(int(x*o.FWidth), int(y*o.FHeight), c)
}

func (o Canvas) Includes(x, y int) bool {
	return x >= 0 && y >= 0 && x < o.Width && y < o.Height
}

func (o Canvas) Get(x, y int) (r byte, g byte, b byte) {
	if o.Includes(x, y) {
		r16, g16, b16, _ := o.Pixels.At(x, y).RGBA()
		r, g, b = byte(r16>>8), byte(g16>>8), byte(b16>>8)
	}
	return
}

func (o Canvas) Set(x, y int, c Color) {
	if o.Includes(x, y) {
		o.Pixels.SetNRGBA(x, y, c.N)
	}
}

// Fill a rectangle with a color.
func (o Canvas) Fill(x1, y1, x2, y2 int, c Color) {
	bounds := image.Rect(x1, y1, x2, y2)
	Say("Min", bounds.Min)
	Say("Max", bounds.Max)
	Say("Color", c)
	draw.Draw(o.Pixels, bounds, &image.Uniform{c.N}, image.ZP, draw.Src)
}

func (o Canvas) WritePng(w io.Writer) {
	e := png.Encode(w, o.Pixels)
	if e != nil {
		panic(e)
	}
}

func Say(aa ...interface{}) {
	for _, a := range aa {
		Fprintf(os.Stderr, "%v ", a)
	}
	Fprintf(os.Stderr, "\n")
}
