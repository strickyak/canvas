/*
	Canvas is an opaque rectangular pixmap.
	It starts out black.
	You can set points.
	You can file it out as a png.
	Because we are RGB & opaque, we have our own Color type.

	Optimized for simplicity, not speed.
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

// RGB constructs a Color object from given red, green, and blue bytes.
func RGB(r, g, b byte) Color {
	return Color{
		N: color.NRGBA{R: r, G: g, B: b, A: 255},
	}
}

// RGB decodes a Color object into red, green, and blue bytes.
func (o Color) RGB() (byte, byte, byte) {
	return o.N.R, o.N.G, o.N.B
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

// Dup returns a duplicate of the receiver.
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
func (o Canvas) FSet(x, y float64, clr Color) {
	o.Set(int(x*o.FWidth), int(y*o.FHeight), clr)
}

// Includes tells if the canvas includes the given coordinate.
func (o Canvas) Includes(x, y int) bool {
	return x >= 0 && y >= 0 && x < o.Width && y < o.Height
}

// Get a point, as red, green, and blue bytes.
func (o Canvas) Get(x, y int) (r byte, g byte, b byte) {
	if o.Includes(x, y) {
		r16, g16, b16, _ := o.Pixels.At(x, o.Height-1-y).RGBA()
		r, g, b = byte(r16>>8), byte(g16>>8), byte(b16>>8)
	}
	return
}

// Set one point to given color.
func (o Canvas) Set(x, y int, clr Color) {
	if o.Includes(x, y) {
		o.Pixels.SetNRGBA(x, o.Height-1-y, clr.N)
	}
}

// Fill a rectangle with a color.
func (o Canvas) Fill(x1, y1, x2, y2 int, clr Color) {
	bounds := image.Rect(x1, o.Height-1-y1, x2, o.Height-1-y2)
	// Say("Min", bounds.Min)
	// Say("Max", bounds.Max)
	// Say("Color", clr)
	draw.Draw(o.Pixels, bounds, &image.Uniform{clr.N}, image.ZP, draw.Src)
}

// Grid draws a debugging grid on a canvas.
func (o Canvas) Grid(gap int, clr Color) {
	// If final gap would land on o.Width, draw it at o.Width-1.
	for x := 0; x <= o.Width; x += gap {
		if x == o.Width && gap > 1 {
			x--
		}
		o.paintVertLine(x, 0, o.Height-1, clr)
	}
	for y := 0; y <= o.Height; y += gap {
		if y == o.Height && gap > 1 {
			y--
		}
		o.paintHorzLine(y, 0, o.Width-1, clr)
	}
}

// WritePng writes the contents of the canvas as a .png file to the given writer.
func (o Canvas) WritePng(w io.Writer) {
	e := png.Encode(w, o.Pixels)
	if e != nil {
		panic(e)
	}
}

// Say writes debugging info to Stderr.  It takes arbitrary strings, objects, or values, and uses %v on them.
func Say(aa ...interface{}) {
	for _, a := range aa {
		Fprintf(os.Stderr, "%v ", a)
	}
	Fprintf(os.Stderr, "\n")
}

// paintHorzLine expects x1 < x2.
func (o *Canvas) paintHorzLine(y, x1, x2 int, clr Color) {
	// Say("paintHorzLine", y, x1, x2)
	if y < 0 || y >= o.Height {
		return
	}
	if x1 < 0 {
		x1 = 0
	}
	if x2 >= o.Width {
		x2 = o.Width - 1
	}
	for i := x1; i <= x2; i++ {
		o.Set(i, y, clr)
	}
}

// paintHorzLine expects y1 < y2.
func (o *Canvas) paintVertLine(x, y1, y2 int, clr Color) {
	// Say("paintVertLine", x, y1, y2)
	if x < 0 || x >= o.Width {
		return
	}
	if y1 < 0 {
		y1 = 0
	}
	if y2 >= o.Width {
		y2 = o.Width - 1
	}
	for i := y1; i <= y2; i++ {
		o.Set(x, i, clr)
	}
}
