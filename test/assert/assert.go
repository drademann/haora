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

package assert

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func Output(t *testing.T, out *bytes.Buffer, expected string) {
	t.Helper()
	if strings.HasPrefix(expected, "\n") {
		expected = expected[1:]
	}
	expected = strings.ReplaceAll(expected, "\t", "")
	if out.String() != expected {
		t.Errorf("expected output \n%q, but got \n%q", expected, out.String())
	}
}

func Duration(t *testing.T, name string, d, expected time.Duration) {
	t.Helper()
	if d != expected {
		t.Errorf("expected %s duration of %v, but got %v", name, expected, d)
	}
}
