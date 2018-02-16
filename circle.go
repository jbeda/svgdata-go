package svgdata

import (
	"github.com/jbeda/geom"
)

type Circle struct {
	Center geom.Coord
	Radius float64
}

func (me *Circle) Draw(svg *SVGWriter, s ...string) {
	svg.Circle(me.Center, me.Radius, s...)
}
