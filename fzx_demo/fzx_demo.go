package main

import (
	"flag"
	"github.com/strickyak/canvas"
	"os"
)

var F = flag.String("f", "", "FZX font pathname")

var M = flag.String("m", "The quick brown fox jumped over the lazy dog's back.", "message")

func main() {
	flag.Parse()
	f := canvas.ReadFZX(*F)
	for i := 32; i <= f.LastChar(); i++ {
		f.Demo(i)
	}

	z := make([]byte, len(*M))
	for i := 0; i < len(*M); i++ {
		z[i] = 32
		if (*M)[i] > 32 {
			z[i] = 128 | (*M)[i]
		}
	}

	c := canvas.NewCanvas(400, 100)
	x, y := 10, 90
	x, y = c.Print(x, y, f, canvas.White, *M)
	x, y = 10, 60
	x, y = c.Print(x, y, f, canvas.White, string(z))
	c.WritePng(os.Stdout)

}
