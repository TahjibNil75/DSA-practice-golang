## Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target. You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

 

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

## Notes:
1. 
```go
nums := []int{1, 2, 3, 5, 7, 9}
var ptr1, ptr2 int = 0, len(nums)-1

for ptr1 < ptr2 {
    current := nums[ptr2] + nums[ptr1]
    // Rest of the code...
}

```
In the first iteration:
ptr1 points to the element 1 at index 0.
ptr2 points to the element 9 at index 5.
current is calculated as 9 + 1, which is 10.


In the second iteration:
ptr1 is incremented to point to the element 2 at index 1.
ptr2 is decremented to point to the element 7 at index 4.
current is calculated as 7 + 2, which is 9.


The loop continues to iterate until ptr1 becomes equal to or greater than ptr2. In each iteration, current is calculated based on the elements at the current positions of ptr1 and ptr2.

This loop is part of the twoSum function designed to find two indices in the array such that the elements at those indices add up to a specific target value. The loop moves towards the center of the array by adjusting the pointers and examining pairs of elements along the way.
