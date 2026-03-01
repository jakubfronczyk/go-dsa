package arrays

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, name string, got, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("%s failed: got %v, expected %v", name, got, expected)
	}
}

func TestReverseCopy(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}

	t.Run("ReverseCopy", func(t *testing.T) {
		result := ReverseCopy(input)
		assertEqual(t, "Reverse Copy", result, expected)
	})
	t.Run("ReverseInPlace", func(t *testing.T) {
		// Need a fresh copy because this modifies the slice directly
		data := make([]int, len(input))
		copy(data, input)

		ReverseInPlace(data)
		assertEqual(t, "ReverseInPlace", data, expected)
	})
	t.Run("ReverseIdiomatic", func(t *testing.T) {
		data := make([]int, len(input))
		copy(data, input)

		ReverseIdiomatic(data)
		assertEqual(t, "ReverseIdiomatic", data, expected)
	})
}
