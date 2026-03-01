package arrays

func ReverseCopy(nums []int) []int {
	var result []int
	for i := len(nums) - 1; i >= 0; i-- {
		result = append(result, nums[i])
	}
	return result
}

func ReverseInPlace(nums []int) {
	i := 0             // Left pointer
	j := len(nums) - 1 // Right pointer

	// Loop is going until pointers meet in the middle
	for i < j {
		// Swap the values at i and j
		nums[i], nums[j] = nums[j], nums[i]
		// Move pointers towards each other
		i++
		j--
	}
}

func ReverseIdiomatic(nums []int) {
	// Init: i and j
	// Condition: i < j
	// Post: increment i and decrement j simultaneously
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
