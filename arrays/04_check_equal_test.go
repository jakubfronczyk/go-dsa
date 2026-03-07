package arrays

import "testing"

func TestEqual(t *testing.T) {
	type testCase struct {
		name     string
		sliceA   []int
		sliceB   []int
		expected bool
	}

	tests := []testCase{
		{
			name:     "Both empty",
			sliceA:   []int{},
			sliceB:   []int{},
			expected: true,
		},
		{
			name:     "Different lengths",
			sliceA:   []int{1, 2},
			sliceB:   []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "Same length, different values",
			sliceA:   []int{1, 2, 3},
			sliceB:   []int{1, 9, 3},
			expected: false,
		},
		{
			name:     "Same values, different order",
			sliceA:   []int{1, 2, 3},
			sliceB:   []int{3, 2, 1},
			expected: false,
		},
		{
			name:     "Exactly the same",
			sliceA:   []int{10, 20, 30},
			sliceB:   []int{10, 20, 30},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Equal(tc.sliceA, tc.sliceB)
			if result != tc.expected {
				t.Errorf("%s failed: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
