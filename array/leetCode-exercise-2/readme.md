# Search Insert Position


## Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

### Example 1:
Input: nums = [1,3,5,6], target = 5
Output: 2

### Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

### Example 3:
Input: nums = [1,3,5,6], target = 7
Output: 4

#### Notes:

First, I tried with this approach
```go
func searchItems(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return len(nums)
}
```

but it does not fulfill case 2