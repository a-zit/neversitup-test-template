package assignments

import "testing"

func TestFindOdd(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Single element 7", []int{7}, 7},
		{"Single element 0", []int{0}, 0},
		{"Two 1s and one 2", []int{1, 1, 2}, 2},
		{"Three 0s and two 1s", []int{0, 1, 0, 1, 0}, 0},
		{"Complex case", []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindOdd(tt.input)
			if result != tt.expected {
				t.Errorf("FindOdd(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
