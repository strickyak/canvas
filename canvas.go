/*
	Canvas is square, for now.  Might change that later.
	It starts out black and opaque.
	You can set points.
	You can file it out as a png.
*/
package canvas

import (
	"image"
	"image/color"
	"image/png"
	"io"
)

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
	black := color.NRGBA{R: 1, G: 1, B: 1, A: 255}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			z.Pixels.SetNRGBA(x, y, black)
		}
	}
	return z
}

// FSet takes float64 x and y, which are visible on the unit square [0, 1) x [0, 1).
func (o Canvas) FSet(x, y float64, c Color) {
	o.Set(int(x*o.FWidth), int(y*o.FHeight), c)
}

func (o Canvas) Set(x, y int, c Color) {
	if x >= 0 && y >= 0 && x < o.Width && y < o.Height {
		o.Pixels.SetNRGBA(x, y, c.N)
	}
}

func (o Canvas) WritePng(w io.Writer) {
	e := png.Encode(w, o.Pixels)
	if e != nil {
		panic(e)
	}
}
