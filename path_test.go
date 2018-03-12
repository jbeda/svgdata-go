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

func TestLoadSavePath(t *testing.T) {
	assert := assert.New(t)

	data0 := []byte(`<svg xmlns="http://www.w3.org/2000/svg"><path d="M0,0 Z" style="color:red" ></path></svg>`)
	r0, err := Unmarshal(data0)
	assert.NoError(err)

	data1, err := Marshal(r0)
	assert.NoError(err)

	r1, err := Unmarshal(data1)
	assert.NoError(err)

	assert.Equal(r0, r1)
}
