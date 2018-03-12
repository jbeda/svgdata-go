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
	"regexp"
	"strconv"
	"strings"
)

type Path struct {
	Attrs    AttrMap
	SubPaths []SubPath
}

type SubPath struct {
	Commands []PathCommand
}

type PathCommand struct {
	Command byte
	Params  []float64
}

func init() {
	unknownCreator = createPath
}

func createPath() Node {
	return &Path{}
}

func (p *Path) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error

	if start.Name.Space != SvgNs {
		return fmt.Errorf("Parsing non-SVG element: %v", start.Name)
	}
	p.Attrs = makeAttrMap(start.Attr)
	_, _, err = readChildren(d, &start)
	return err
}

func (p *Path) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	se := MakeStartElement("path", p.Attrs)

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

// Parse the path string. Each command represents a single instance of the
// command. This is to enable easier further processing.
func ParsePathString(d string) ([]SubPath, error) {
	sp := []SubPath{}

	d = skipWsp(d)

	return sp, nil
}

func skipWsp(d string) string {
	return strings.TrimLeft(d, " \t\r\n")
}

func skipCommaWsp(d string) string {
	d = skipWsp(d)
	if len(d) > 0 && d[0] == ',' {
		d = d[1:]
	}
	d = skipWsp(d)

	return d
}

var floatRE *regexp.Regexp = regexp.MustCompile(`^[-+]?[0-9]*\.?[0-9]+(?:[eE][-+]?[0-9]+)?`)

// readNumber reads the next number from the string.  It returns the rest of the
// string, the parsed number and an ok bool.  If ok is false then there is no
// number and d is returned.
func readNumber(d string) (string, float64, bool) {
	loc := valueRE.FindStringIndex(d)
	if loc == nil {
		return d, 0, false
	}

	v, err := strconv.ParseFloat(d[loc[0]:loc[1]], 64)
	if err != nil {
		return d, 0, false
	}

	return d[loc[1]:], v, true
}

// readByte reads the next byte from the string. It returns the remaining
// string, the byte and an ok bool. If bool is false then the string is empty.
func readByte(d string) (string, byte, bool) {
	if len(d) == 0 {
		return d, 0, false
	}
	return d[1:], d[0], true
}
