// Copyright 2018 Joe Beda
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package svgdata

import (
	"encoding/xml"
	"fmt"

	"github.com/jbeda/geom"
)

// Circle is an SVG element that has no specialized code or representation.
type Circle struct {
	Attrs  AttrMap
	Center geom.Coord
	Radius float64
}

func init() {
	RegisterNodeCreator("circle", createCircle)
}

func createCircle() Node {
	return &Circle{}
}

func NewCircle(c geom.Coord, r float64) *Circle {
	return &Circle{Center: c, Radius: r}
}

func (c *Circle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error

	if start.Name.Space != SvgNs {
		return fmt.Errorf("Parsing non-SVG element: %v", start.Name)
	}

	c.Attrs = makeAttrMap(start.Attr)

	cx, err := parseValue(c.Attrs["cx"])
	if err != nil {
		return err
	}
	delete(c.Attrs, "cx")

	cy, err := parseValue(c.Attrs["cy"])
	if err != nil {
		return err
	}
	delete(c.Attrs, "cy")

	r, err := parseValue(c.Attrs["r"])
	if err != nil {
		return err
	}
	delete(c.Attrs, "r")

	c.Center = geom.Coord{X: cx, Y: cy}
	c.Radius = r

	_, _, err = readChildren(d, &start)
	return err
}

func (c *Circle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	am := copyAttrMap(c.Attrs)
	am["cx"] = floatToString(c.Center.X)
	am["cy"] = floatToString(c.Center.Y)
	am["r"] = floatToString(c.Radius)

	se := MakeStartElement("circle", am)

	err := e.EncodeToken(se)
	if err != nil {
		return err
	}

	err = e.EncodeToken(se.End())
	if err != nil {
		return err
	}

	return nil
}
