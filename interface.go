package svgdata

import "github.com/jbeda/geom"

type Element interface {
	Draw(w *SVGWriter, s ...string)
}

type PathSegment interface {
	P1() *geom.Coord
	P2() *geom.Coord
	Reverse()
	PathDraw(w *SVGWriter)
}
