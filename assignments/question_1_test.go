package assignments

import (
	"reflect"
	"sort"
	"testing"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"aabb", []string{"aabb", "baab", "abab", "abba", "baba", "bbaa"}},
	}

	for _, test := range tests {
		result := Permutations(test.input)

		sort.Strings(result)
		sort.Strings(test.expected)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Permutations(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
