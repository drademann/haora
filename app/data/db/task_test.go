//
// Copyright 2024-2024 The Haora Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package db

import (
	"github.com/drademann/fugo/test/assert"
	"testing"
)

func TestTagArray_Scan(t *testing.T) {
	ta := TagArray{}

	err := ta.Scan("haora,lunch")

	assert.NoError(t, err)
	assert.ContainsInAnyOrder(t, ta, TagArray{"haora", "lunch"})
}

func TestTagArray_Value(t *testing.T) {
	ta := TagArray{"haora", "lunch"}

	value, err := ta.Value()

	assert.NoError(t, err)
	assert.Equal(t, "haora,lunch", value)
}
