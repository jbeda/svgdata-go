package svgdata

// OptimizedPathCollection takes a set of Paths and PathSegments and constructs
// continuous paths. This is sometimes referred to as "chains".  After adding
// all of the Paths and PathSegments, call Optimize.
type OptimizedPathCollection struct {
	Paths []*Path
}

func (opc *OptimizedPathCollection) Draw(svg *SVGWriter, s ...string) {
	for _, path := range opc.Paths {
		path.Draw(svg, s...)
	}
}

func (opc *OptimizedPathCollection) NumPaths() int {
	return len(opc.Paths)
}

func (opc *OptimizedPathCollection) AddSegment(p PathSegment) {
	path := new(Path)
	path.PushFront(p)
	opc.AddPath(path)
}

func (opc *OptimizedPathCollection) AddPath(np *Path) {
	npP1 := np.Front().P1()
	npP2 := np.Back().P2()
	for _, path := range opc.Paths {
		if AlmostEqualsCoord(*npP2, *path.Front().P1()) {
			path.PushPathFront(np)
			return
		}
		if AlmostEqualsCoord(*npP1, *path.Back().P2()) {
			path.PushPathBack(np)
			return
		}
		if AlmostEqualsCoord(*npP1, *path.Front().P1()) {
			np.Reverse()
			path.PushPathFront(np)
			return
		}
		if AlmostEqualsCoord(*npP2, *path.Back().P2()) {
			np.Reverse()
			path.PushPathBack(np)
			return
		}
	}

	opc.Paths = append(opc.Paths, np)
}

func (opc *OptimizedPathCollection) Optimize() {
	// Loop through until the number of paths stabilizes
	for i := 0; ; i++ {
		prevNumPaths := len(opc.Paths)

		oldPaths := opc.Paths
		opc.Paths = nil
		for _, p := range oldPaths {
			opc.AddPath(p)
		}

		if prevNumPaths == len(opc.Paths) {
			break
		}
	}

	for _, path := range opc.Paths {
		if AlmostEqualsCoord(*path.Front().P1(), *path.Back().P2()) {
			path.Closed = true
		}
	}
}
