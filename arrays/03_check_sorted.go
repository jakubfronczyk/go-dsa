package arrays

import "errors"

func IsSorted(nums []int) (bool, error) {
	if len(nums) == 0 {
		return false, errors.New("Array cannot be empty")
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false, nil
		} else {
			continue
		}
	}
	return true, nil
}
