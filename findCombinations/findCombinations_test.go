package findCombinations

import (
	"reflect"
	"testing"
)

func TestFindCombinations(t *testing.T) {
	var tests = []struct {
		str  string
		want map[string]int
	}{
		{"1X", map[string]int{"10": 1, "11": 1}},
		{"XX", map[string]int{"10": 1, "11": 1, "01": 1, "00": 1}},
		{"1", map[string]int{"1": 1}},
	}
	for _, test := range tests {
		got := make(map[string]int)
		for s := range FindCombinations(test.str) {
			got[s] += 1
		}
		if reflect.DeepEqual(test.want, got) == false {
			t.Errorf("FindCombinations(%s) = %+v expected %+v.", test.str, got, test.want)
		}
	}

}
