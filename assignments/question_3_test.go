package assignments

import (
	"reflect"
	"testing"
)

func TestCountSmileys(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{":)", ";(", ";}", ":-D"}, 2},
		{[]string{";D", ":-(", ":-)", ";~)"}, 3},
		{[]string{";]", ":[", ";*", ":$", ";-D"}, 1},
		{[]string{}, 0},
		{[]string{":)", ":)", ":)", ":)"}, 4},
		{[]string{";(", ";(", ";(", ";("}, 0},
		{[]string{}, 0},
	}

	for _, test := range tests {
		result := countSmileys(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("countSmileys(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
