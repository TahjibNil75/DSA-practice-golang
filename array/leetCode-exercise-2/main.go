package main

import (
	"fmt"
)

// Function to find the index at which the 'target' should be inserted into the sorted array 'nums'.
// Iterates through the array using index 'i' and element 'v'.
func searchItems(nums []int, target int) int {
	// Iterate through the array 'nums'.
	for i, v := range nums {
		// Check if the current element 'v' is greater than or equal to the target.
		if v >= target {
			// If true, return the current index 'i' as the position where 'target' should be inserted.
			return i
		}
	}
	// If no element is greater than or equal to the target, return the length of the array.
	// This indicates that 'target' should be inserted at the end to maintain sorted order.
	return len(nums)
}

func main() {
	// Example usage with a predefined array 'number' and a target value 'target'.
	number := []int{1, 3, 5, 6}
	target := 5

	// Call the searchItems function with the provided numbers and target.
	result := searchItems(number, target)

	// Print the result.
	fmt.Printf("Index of the target %d: %d\n", target, result)
}
