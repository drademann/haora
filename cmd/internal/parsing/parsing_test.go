//
// Copyright 2024-2025 The Haora Authors
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

package parsing

import (
	"github.com/drademann/fugo/test"
	"github.com/drademann/fugo/test/assert"
	"github.com/drademann/haora/app/datetime"
	"reflect"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Time("14:42"))

	testCases := []struct {
		name       string
		flag       string
		args       []string
		wantHour   int
		wantMinute int
		wantArgs   []string
		wantErr    bool
	}{
		// now
		{"now flag", "now", []string{"tag", "task"}, 14, 42, []string{"tag", "task"}, false},
		{"now as first arg", "", []string{"now", "tag", "task"}, 14, 42, []string{"tag", "task"}, false},
		// times
		{"standard flag", "09:15", []string{"tag", "task"}, 9, 15, []string{"tag", "task"}, false},
		{"standard flag, no leading 0", "9:15", []string{"tag", "task"}, 9, 15, []string{"tag", "task"}, false},
		{"flag, no semicolon", "0915", []string{"tag", "task"}, 9, 15, []string{"tag", "task"}, false},
		{"flag, no semicolon, no leading 0", "915", []string{"tag", "task"}, 9, 15, []string{"tag", "task"}, false},
		{"flag, no semicolon, hour >= 10", "1730", []string{"tag", "task"}, 17, 30, []string{"tag", "task"}, false},
		{"flag, just single digit", "9", []string{"tag", "task"}, 9, 0, []string{"tag", "task"}, false},
		{"as first arg", "", []string{"09:15", "tag", "task"}, 9, 15, []string{"tag", "task"}, false},
		{"as first arg, no leading 0", "", []string{"9:15", "tag", "task"}, 9, 15, []string{"tag", "task"}, false},
		{"as first arg, no semicolon", "", []string{"0915", "tag", "task"}, 9, 15, []string{"tag", "task"}, false},
		{"as first arg, no semicolon, no leading 0", "", []string{"915", "tag", "task"}, 9, 15, []string{"tag", "task"}, false},
		{"as first arg, no semicolon, hour >= 10", "", []string{"1730", "tag", "task"}, 17, 30, []string{"tag", "task"}, false},
		{"as first arg, just single digit", "", []string{"9", "tag", "task"}, 9, 0, []string{"tag", "task"}, false},
		{"early morning, no semicolon", "015", []string{"tag", "task"}, 0, 15, []string{"tag", "task"}, false},
		// errors
		{"hour > 23", "30:15", []string{"tag", "task"}, 0, 0, []string{"tag", "task"}, true},
		{"minute > 59", "12:75", []string{"tag", "task"}, 0, 0, []string{"tag", "task"}, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parsedTime, parsedArgs, err := TimeWithArgs(tc.flag, tc.args)
			if err != nil && !tc.wantErr {
				t.Fatalf("expected no error, got %v", err)
			}
			if err == nil && tc.wantErr {
				t.Fatalf("expected an error, got none")
			}
			if !tc.wantErr {
				if parsedTime.Hour() != tc.wantHour || parsedTime.Minute() != tc.wantMinute {
					t.Errorf("parsed time %v does not match expected hour %d or minute %d", parsedTime, tc.wantHour, tc.wantMinute)
				}
				if !reflect.DeepEqual(parsedArgs, tc.wantArgs) {
					t.Errorf("remaining args %v after parsing do not match expected args %v", parsedArgs, tc.wantArgs)
				}
			}
		})
	}
}

func TestParseWeekday(t *testing.T) {
	testCases := []struct {
		weekdayStr string
		expected   time.Weekday
	}{
		{"Monday", time.Monday},
		{"Mon", time.Monday},
		{"mo", time.Monday},
		{"wedn", time.Wednesday},
	}

	for _, tc := range testCases {
		t.Run(tc.weekdayStr, func(t *testing.T) {
			wd, err := Weekday(tc.weekdayStr)

			assert.NoError(t, err)
			if wd != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, wd)
			}
		})
	}
}
