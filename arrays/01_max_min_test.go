package arrays

import "testing"

func TestFindMax(t *testing.T) {
	nums := []int{1, 2, 3}

	got, err := findMax(nums)
	if err != nil {
		t.Errorf("Unexpexted error: %v", err)
	}

	if got != 3 {
		t.Errorf("Expected 5 but got %d", got)
	}
}

func TestFindMin(t *testing.T) {
	nums := []int{1, 2, 3}

	got, err := findMin(nums)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if got != 1 {
		t.Errorf("Expected 1 but got %d", got)
	}
}
