package arrays

import (
	"testing"
)

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"Unsorted", []int{5, 1, 3, 2}, false},
		{"Sorted", []int{1, 2, 3, 4}, true},
		{"Empty", []int{}, false}, // Assuming error handling for empty
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsSorted(tt.input)

			// NOTE: Since empty array returns an error in logic
			// we need to adjuts error checking
			if err != nil && len(tt.input) > 0 {
				t.Errorf("Unexpexted error %v", err)
			}

			if got != tt.expected {
				t.Errorf("got %v, expected %v", got, tt.expected)
			}
		})
	}
}
