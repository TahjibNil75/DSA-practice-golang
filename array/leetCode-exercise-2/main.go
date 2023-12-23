package main

import (
	"fmt"
)

func searchItems(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}

	return len(nums)
}

func main() {
	number := []int{1, 3, 5, 6}
	target := 5

	result := searchItems(number, target)
	fmt.Printf("Index of the target %d: %d\n", target, result)

}
