/*
	Draw triangles given three integer points.
*/
package canvas

func (o *Canvas) PaintTriangle(x1, y1, x2, y2, x3, y3 int, clr Color) {
	// Bubble sort on Y.
	if y1 > y2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}
	if y2 > y3 {
		x2, y2, x3, y3 = x3, y3, x2, y2
	}
	if y1 > y2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	if y1 == y2 {
		o.paintTriangleSameYY(y1, x1, x2 /**/, x3, y3, clr)
	} else if y2 == y3 {
		o.paintTriangleSameYY(y3, x2, x3 /**/, x1, y1, clr)
	} else {
		o.paintTriangleDiffYs(x1, y1, x2, y2, x3, y3, clr)
	}
}

// paintTriangleDiffYs expects y1 < y2 < y3.
func (o *Canvas) paintTriangleDiffYs(x1, y1, x2, y2, x3, y3 int, clr Color) {
	Say("paintTriangleDiffYs", x1, y1, x2, y2, x3, y3)
	// x4 is on the line from (x1, y1) to (x3, y3) where y = y2.
	x4 := x1 + int(float64(y2-y1)*float64(x3-x1)/float64(y3-y1))

	// Now draw two triangles, using horizontal line at y2.
	o.paintTriangleSameYY(y2, x2, x4 /**/, x1, y1, clr)
	o.paintTriangleSameYY(y2, x2, x4 /**/, x3, y3, clr)
}

// paintTriangleSameYY does not know y <=> y3, nor anything about xs.
func (o *Canvas) paintTriangleSameYY(y, x1, x2 int, x3, y3 int, clr Color) {
	Say("paintTriangleSameYY", y, x1, x2, x3, y3)
	if y == y3 {
		// Avoid DivisionByZero by handling y==y3 here.
		o.paintTriangleSameYYY(y, x1, x2, x3, clr)
		return
	}

	// Sort x1 & x2.
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	if y < y3 {
		dx1 := float64(x3-x1) / float64(y3-y)
		Say("A: y<y3", "num", x3-x1, "den", y3-y, "dx1", dx1)
		dx2 := float64(x3-x2) / float64(y3-y)
		Say("A: y<y3", "num", x3-x2, "den", y3-y, "dx2", dx2)
		for i := y; i <= y3; i++ {
			o.paintHorzLine(i, x1+int(float64(i-y)*dx1), x2+int(float64(i-y)*dx2), clr)
		}
	} else {
		dx1 := float64(x3-x1) / float64(y-y3)
		Say("B: y>y3", "num", x3-x1, "den", y-y3, "dx1", dx1)
		dx2 := float64(x3-x2) / float64(y-y3)
		Say("B: y>y3", "num", x3-x2, "den", y-y3, "dx2", dx2)
		for i := y3; i <= y; i++ {
			o.paintHorzLine(i, x3-int(float64(i-y3)*dx1), x3-int(float64(i-y3)*dx2), clr)
		}
	}
}

func (o *Canvas) paintTriangleSameYYY(y, x1, x2, x3 int, clr Color) {
	// Bubble sort on X.
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if x2 > x3 {
		x2, x3 = x3, x2
	}
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	o.paintHorzLine(y, x1, x3, clr)
}
