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

// Root represents the root <svg> element.
type Root struct {
	Viewbox  *geom.Rect
	Attrs    AttrMap
	Children []Node
}

var _ Node = (*Root)(nil)

func CreateRoot() Node {
	return &Root{}
}

func init() {
	RegisterNodeCreator("svg", CreateRoot)
}

func (r *Root) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != SvgNs || start.Name.Local != "svg" {
		return fmt.Errorf("Unexpected element for SVG Root: %v", start.Name)
	}

	r.Attrs = makeAttrMap(start.Attr)
	// XML Namespaces are totally borked in encoding/xml.
	delete(r.Attrs, "xmlns")

	var err error
	r.Children, _, err = readChildren(d, &start)
	if err != nil {
		return err
	}
	return nil
}

func (r *Root) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	se := MakeStartElement("svg", r.Attrs)
	// XML Namespaces are totally borked in encoding/xml.
	se.Name.Space = SvgNs

	err := e.EncodeToken(se)
	if err != nil {
		return err
	}

	for _, c := range r.Children {
		err = e.Encode(c)
		if err != nil {
			return err
		}
	}

	err = e.EncodeToken(se.End())
	if err != nil {
		return err
	}

	return nil
}
