package binarysearch

import "fmt"

// todo fix
func SearchInsert (nums []int, target int) int {
	left, right := 0, len(nums) - 1

	for left <= right {
		middle := (left + right) / 2
		middle_num := nums[middle]

		if middle_num == target {
			return middle
		}
		if (target > middle) {
			left = middle + 1
		} else {
			right = middle - 1
		}

		fmt.Printf("left: %d, right: %d, middle: %d\n", left, right, middle)
	}

	return left
} 

// 1,3,5,6
//
// 0 3 1
// 2 3 2
// 2 2 

// rest of the binary search problmes easy neetcode















