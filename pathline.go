package svgdata

import (
	"github.com/jbeda/geom"
)

type PathLine struct {
	A, B geom.Coord
}

var _ PathSegment = (*PathLine)(nil)

func NewPathLine(a, b geom.Coord) *PathLine {
	return &PathLine{A: a, B: b}
}

func AlmostEqualsPathLines(a, b *PathLine) bool {
	return (AlmostEqualsCoord(a.A, b.A) && AlmostEqualsCoord(a.B, b.B)) ||
		(AlmostEqualsCoord(a.A, b.B) && AlmostEqualsCoord(a.B, b.A))
}

func (cl *PathLine) Equals(oi interface{}) bool {
	ocl, ok := oi.(*PathLine)
	return ok && AlmostEqualsPathLines(cl, ocl)
}

func (cl *PathLine) Bounds() geom.Rect {
	r := geom.Rect{cl.A, cl.A}
	r.ExpandToContainCoord(cl.B)
	return r
}

func (cl *PathLine) P1() *geom.Coord { return &cl.A }
func (cl *PathLine) P2() *geom.Coord { return &cl.B }
func (cl *PathLine) PathDraw(svg *SVGWriter) {
	svg.PathLineTo(cl.B)
}
func (cl *PathLine) Reverse() {
	cl.A, cl.B = cl.B, cl.A
}
