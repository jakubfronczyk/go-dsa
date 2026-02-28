package arrays

import (
	"errors"
)

// Find the maximut and minimun elemetns in the array.
func findMax(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("array cannot be empty")
	}

	max := nums[0]

	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func findMin(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("array cannot be empty")
	}

	min := nums[0]

	for _, v := range nums {
		if v < min {
			min = v
		}
	}

	return min, nil
}
