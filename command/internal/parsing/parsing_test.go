package parsing

import (
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/test"
	"reflect"
	"testing"
)

func TestTime(t *testing.T) {
	datetime.MockNowAt(t, test.MockTime("14:42"))

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
		{"as first arg", "", []string{"09:15", "tag", "task"}, 9, 15, []string{"tag", "task"}, false},
		{"as first arg, no leading 0", "", []string{"9:15", "tag", "task"}, 9, 15, []string{"tag", "task"}, false},
		{"as first arg, no semicolon", "", []string{"0915", "tag", "task"}, 9, 15, []string{"tag", "task"}, false},
		{"as first arg, no semicolon, no leading 0", "", []string{"915", "tag", "task"}, 9, 15, []string{"tag", "task"}, false},
		{"as first arg, no semicolon, hour >= 10", "", []string{"1730", "tag", "task"}, 17, 30, []string{"tag", "task"}, false},
		// errors
		{"hour > 23", "30:15", []string{"tag", "task"}, 0, 0, []string{"tag", "task"}, true},
		{"minute > 59", "12:75", []string{"tag", "task"}, 0, 0, []string{"tag", "task"}, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parsedTime, parsedArgs, err := Time(tc.flag, tc.args)
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
