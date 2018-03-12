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
)

// Unknown is an SVG element that has no specialized code or representation.
type Unknown struct {
	Name     string
	Attrs    AttrMap
	Children []Node
	Text     string
}

func init() {
	unknownCreator = createUnknown
}

func createUnknown() Node {
	return &Unknown{}
}

func (u *Unknown) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error

	if start.Name.Space != SvgNs {
		return fmt.Errorf("Parsing non-SVG element: %v", start.Name)
	}
	u.Name = start.Name.Local
	u.Attrs = makeAttrMap(start.Attr)
	u.Children, u.Text, err = readChildren(d, &start)
	return err
}

func (u *Unknown) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	se := MakeStartElement(u.Name, u.Attrs)

	err := e.EncodeToken(se)
	if err != nil {
		return err
	}

	if len(u.Children) != 0 {
		for _, c := range u.Children {
			err = e.Encode(c)
			if err != nil {
				return err
			}
		}
	} else if u.Text != "" {
		tt := xml.CharData(u.Text)
		err = e.EncodeToken(tt)
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
