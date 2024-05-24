package binarysearch

import (
	"slices"
)

func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		curr_num := nums[mid] 
		if curr_num == target {
			return mid
		}
		if target > curr_num {
			left = mid + 1
		} else {
			right = mid - 1 
		}
	}
	return left
}

func squaresOfSortedArray(nums []int) []int {
	new_nums := make([]int, 0)
	left, right, i := 0, len(nums) - 1, 0

	for left <= right {
		left_sq := nums[left] * nums[left]
		right_sq := nums[right] * nums[right]
		if left_sq > right_sq {
			new_nums = append(new_nums, left_sq)
			left += 1
		} else {
			new_nums = append(new_nums, right_sq)
			right -= 1
		}
		i += 1
	}
	slices.Reverse(new_nums)
	return new_nums
}


