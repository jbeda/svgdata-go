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

func TestLoadSaveCircle(t *testing.T) {
	assert := assert.New(t)
	l := litter.Options{HidePrivateFields: false}

	data0 := []byte(`<svg xmlns="http://www.w3.org/2000/svg"><circle cx="3px" cy="4in" r="20" ></circle></svg>`)
	r0, err := Unmarshal(data0)
	assert.NoError(err)

	data1, err := Marshal(r0, false)
	assert.NoError(err)

	r1, err := Unmarshal(data1)
	assert.NoError(err)

	d0 := l.Sdump(r0)
	d1 := l.Sdump(r1)

	assert.Equal(d0, d1)
}
