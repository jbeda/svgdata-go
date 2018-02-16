package svgdata

import (
	"math"

	"github.com/jbeda/geom"
)

// Comparing floating point sucks.  This is probably wrong in the general case
// but is good enough for this application.  Don't assume I know what I'm
// doing here.  I'm pulling stuff out of my butt.
const FLOAT_EQUAL_THRESH = 0.00000001

func FloatAlmostEqual(a, b float64) bool {
	return math.Abs(a-b) < FLOAT_EQUAL_THRESH
}

func AlmostEqualsCoord(a, b geom.Coord) bool {
	return FloatAlmostEqual(a.X, b.X) && FloatAlmostEqual(a.Y, b.Y)
}
