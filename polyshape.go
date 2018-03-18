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

// Polyshape is an SVG element is a shape specified with a list of straight
// lines.
type Polyshape struct {
	Attrs  AttrMap
	Points []geom.Coord
	Closed bool
}

func init() {
	RegisterNodeCreator("polygon", createPolygon)
	RegisterNodeCreator("polyline", createPolyline)
}

func createPolygon() Node {
	return &Polyshape{Closed: true}
}

func createPolyline() Node {
	return &Polyshape{Closed: false}
}

func (p *Polyshape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error

	if start.Name.Space != SvgNs {
		return fmt.Errorf("Parsing non-SVG element: %v", start.Name)
	}

	if start.Name.Local == "polygon" {
		p.Closed = true
	}

	p.Attrs = makeAttrMap(start.Attr)

	_, _, err = readChildren(d, &start)
	return err
}

func (p *Polyshape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	am := copyAttrMap(p.Attrs)

	var se xml.StartElement
	if p.Closed {
		se = MakeStartElement("polygon", am)
	} else {
		se = MakeStartElement("polyline", am)
	}

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
