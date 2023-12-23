package main

import (
	"fmt"
)

func twoSums(nums []int, target int) []int {

	// Initialize two pointers, ptr1 and ptr2, at the beginning and end of the array, respectively.
	var ptr1, ptr2 int = 0, len(nums) - 1

	for ptr1 < ptr2 {
		// Calculate the sum of elements at ptr1 and ptr2
		current := nums[ptr2] + nums[ptr1]

		if current == target {
			return []int{ptr1, ptr2}
		} else if current > target {
			// If the current sum is greater than the target, decrement ptr2
			ptr2--
		} else {
			// If the current sum is less than the target, increment ptr1
			ptr1++
		}
	}
	return []int{}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 19

	result := twoSums(nums, target)
	if len(result) == 2 {
		fmt.Printf("Indices of elements that add up to %d: %v\n", target, result)
	} else {
		fmt.Printf("No such indices found for the target %d\n", target)
	}
}

// this function is not fulfilled case 2
