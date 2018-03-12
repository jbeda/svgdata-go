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

	"github.com/stretchr/testify/assert"
)

func TestParseValue(t *testing.T) {
	assert := assert.New(t)
	p, err := parseValue("1")
	assert.NoError(err)
	assert.InEpsilon(1, p, 0.01)

	p, err = parseValue("1.1px")
	assert.NoError(err)
	assert.InEpsilon(1.1, p, 0.01)

	p, err = parseValue("1.1in")
	assert.NoError(err)
	assert.InEpsilon(105.6, p, 0.01)

	p, err = parseValue("25.4mm")
	assert.NoError(err)
	assert.InEpsilon(96, p, 0.01)

	p, err = parseValue("2.54cm")
	assert.NoError(err)
	assert.InEpsilon(96, p, 0.01)

	p, err = parseValue("72pt")
	assert.NoError(err)
	assert.InEpsilon(96, p, 0.01)

	p, err = parseValue("6pc")
	assert.NoError(err)
	assert.InEpsilon(96, p, 0.01)

	p, err = parseValue("")
	assert.Error(err)

	p, err = parseValue(" 1")
	assert.Error(err)

	p, err = parseValue("1 ")
	assert.Error(err)

	p, err = parseValue("1%")
	assert.Error(err)

	p, err = parseValue("1 in")
	assert.Error(err)
}
