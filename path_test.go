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
	"testing"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"
)

func TestLoadSavePath(t *testing.T) {
	assert := assert.New(t)
	l := litter.Options{HidePrivateFields: false}

	data0 := []byte(`<svg xmlns="http://www.w3.org/2000/svg"><path d="M0,0 Z" style="color:red" ></path></svg>`)
	r0, err := Unmarshal(data0)
	assert.NoError(err)

	data1, err := Marshal(r0, false)
	assert.NoError(err)

	r1, err := Unmarshal(data1)
	assert.NoError(err)

	assert.Equal(l.Sdump(r0), l.Sdump(r1))
}

func TestPathTokenize(t *testing.T) {
	assert := assert.New(t)

	pt, err := pathTokenize("m0,2 ,  3.4 4.4 ")
	assert.NoError(err)
	expect := []pathToken{
		commandToken{'m'},
		numberToken{0},
		numberToken{2},
		numberToken{3.4},
		numberToken{4.4},
		eodToken{},
	}
	assert.Equal(expect, pt)

	pt, err = pathTokenize(" w ")
	assert.Error(err)
	assert.Nil(pt)
}

func NewPathCommand(c byte, p []float64) PathCommand {
	return PathCommand{Command: c, Params: p}
}

func TestParsePathCommands(t *testing.T) {
	assert := assert.New(t)

	cs, err := parsePathCommands("")
	assert.NoError(err)
	assert.Equal([]PathCommand{}, cs)

	cs, err = parsePathCommands("M1 2Z")
	assert.NoError(err)
	assert.Equal([]PathCommand{
		NewPathCommand('M', []float64{1, 2}),
		NewPathCommand('Z', nil),
	}, cs)

	cs, err = parsePathCommands("1 2Z")
	assert.NoError(err)
	assert.Equal([]PathCommand{
		NewPathCommand('L', []float64{1, 2}),
		NewPathCommand('Z', nil),
	}, cs)

	cs, err = parsePathCommands("1 2 3 4")
	assert.NoError(err)
	assert.Equal([]PathCommand{
		NewPathCommand('L', []float64{1, 2}),
		NewPathCommand('L', []float64{3, 4}),
	}, cs)

	cs, err = parsePathCommands("m1 2 3 4")
	assert.NoError(err)
	assert.Equal([]PathCommand{
		NewPathCommand('m', []float64{1, 2}),
		NewPathCommand('l', []float64{3, 4}),
	}, cs)

	_, err = parsePathCommands("m1")
	assert.Error(err)

	_, err = parsePathCommands("m1z")
	assert.Error(err)

	_, err = parsePathCommands("w")
	assert.Error(err)

}

func TestPathParseSave(t *testing.T) {
	assert := assert.New(t)

	sps, err := ParsePathString("M0,0 L10,10 Z")
	assert.NoError(err)
	assert.Len(sps, 1)
	out := SavePathString(sps)
	assert.Equal("M0 0L10 10Z", out)

	// Test missing M at start
	sps, err = ParsePathString("L10,10 z")
	assert.NoError(err)
	assert.Len(sps, 1)
	out = SavePathString(sps)
	assert.Equal("M0 0L10 10z", out)

	// Test 2 subpaths with implied start
	sps, err = ParsePathString("L10,20 Z L20,10")
	assert.NoError(err)
	assert.Len(sps, 2)
	assert.Equal(0.0, sps[0].startX)
	assert.Equal(0.0, sps[0].startY)
	assert.Equal(0.0, sps[0].endX)
	assert.Equal(0.0, sps[0].endY)
	assert.Equal(0.0, sps[1].startX)
	assert.Equal(0.0, sps[1].startY)
	assert.Equal(20.0, sps[1].endX)
	assert.Equal(10.0, sps[1].endY)
	out = SavePathString(sps)
	assert.Equal("M0 0L10 20ZM0 0L20 10", out)

	// Test some relative movement
	sps, err = ParsePathString("m1,2 l1,2 h1 v1 c1,2 2,3 4,5s1,2 3,4q1,2 3,4t1,2a1,2 20 1 0 3,4")
	assert.NoError(err)
	assert.Len(sps, 1)
	assert.Equal(SubPath{
		Commands: []PathCommand{
			{'m', []float64{1, 2}, 0, 0, 1, 2},
			{'l', []float64{1, 2}, 1, 2, 2, 4},
			{'h', []float64{1}, 2, 4, 3, 4},
			{'v', []float64{1}, 3, 4, 3, 5},
			{'c', []float64{1, 2, 2, 3, 4, 5}, 3, 5, 7, 10},
			{'s', []float64{1, 2, 3, 4}, 7, 10, 10, 14},
			{'q', []float64{1, 2, 3, 4}, 10, 14, 13, 18},
			{'t', []float64{1, 2}, 13, 18, 14, 20},
			{'a', []float64{1, 2, 20, 1, 0, 3, 4}, 14, 20, 17, 24},
		},
		startX: 1, startY: 2,
		endX: 17, endY: 24,
	}, sps[0])
	out = SavePathString(sps)
	assert.Equal("m1 2l1 2h1v1c1 2 2 3 4 5s1 2 3 4q1 2 3 4t1 2a1 2 20 1 0 3 4", out)

	// Test absolute movement
	sps, err = ParsePathString("M1,2 L1,2 H1 V1 C1,2 2,3 4,5S1,2 3,4Q1,2 3,4T1,2A1,2 20 1 0 3,4")
	assert.NoError(err)
	assert.Len(sps, 1)
	assert.Equal(SubPath{
		Commands: []PathCommand{
			{'M', []float64{1, 2}, 0, 0, 1, 2},
			{'L', []float64{1, 2}, 1, 2, 1, 2},
			{'H', []float64{1}, 1, 2, 1, 2},
			{'V', []float64{1}, 1, 2, 1, 1},
			{'C', []float64{1, 2, 2, 3, 4, 5}, 1, 1, 4, 5},
			{'S', []float64{1, 2, 3, 4}, 4, 5, 3, 4},
			{'Q', []float64{1, 2, 3, 4}, 3, 4, 3, 4},
			{'T', []float64{1, 2}, 3, 4, 1, 2},
			{'A', []float64{1, 2, 20, 1, 0, 3, 4}, 1, 2, 3, 4},
		},
		startX: 1, startY: 2,
		endX: 3, endY: 4,
	}, sps[0])
	out = SavePathString(sps)
	assert.Equal("M1 2L1 2H1V1C1 2 2 3 4 5S1 2 3 4Q1 2 3 4T1 2A1 2 20 1 0 3 4", out)

	// Test bad string
	sps, err = ParsePathString("hello!")
	assert.Error(err)
	assert.Nil(sps)
}
