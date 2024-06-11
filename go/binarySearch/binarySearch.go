package binarysearch

import (
	"sort"
	"fmt"
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

// *********************************************** 
func getTotal (nums []int) int {
	total := 0
	for _, num := range nums {total += num}
	return total;
}

func isCapacityValid (capacity, days int, weights []int) bool {
	days_count, temp_weight := 1, 0

	for _, weight := range weights {
		if temp_weight + weight > capacity {
			temp_weight = 0
			days_count += 1
		}
		temp_weight += weight
		if days_count > days {
			return false
		}
		fmt.Printf("days: %d, capacity: %d, weight: %d, temp_weight: %d, days_count: %d\n",days , capacity, weight, temp_weight, days_count)
	}
	return days_count == days
}

func shipWithinDays (weights []int, days int) int {
	min, max := slices.Max(weights), getTotal(weights)

	for i := min; i <= max; i++ {
		if (isCapacityValid(i, days, weights)) {
			return i
		}
	}
	return max
}
// *********************************************** 
// potions has to be sorted from max to min
func findRightPotionsPoint (potions []int, spell, success int) int {
	left, right := 0, len(potions) - 1

	for left <= right {
		mid := (left + right) / 2
		total := potions[mid] * spell

		if total >= success && potions[mid - 1] * spell < success {
			return mid
		}

		if success > total {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 
}


func successfulPairs(spells, potions []int, success int) []int {
	output_slice := make([]int, 0)

	sort.Slice(potions, func(i, j int) bool {
		return potions[i] < potions[j] 
	})

	for _, spell := range spells {
		point := findRightPotionsPoint(potions, spell, success)
		if point != -1 {
			output_slice = append(output_slice, len(potions) - point)
		} else {
			output_slice = append(output_slice, 0)
		}
	}
	return output_slice
}

// *********************************************** 
func findRow (matrix[][]int, target int) int {
	for i, row := range matrix {
		if row[0] <= target && row[len(row) - 1] >= target {
			return i
		}
	}
	return -1
}
func findTarget(row []int, target int) bool {
	left, right := 0, len(row) - 1

	for left <= right {
		mid := (left + right) / 2
		curr := row[mid]

		if curr == target {
			return true
		}

		if target > curr {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
func searchMatrix (matrix [][]int, target int) bool {
	active_row := findRow(matrix, target)
	if active_row == -1 {
		return false
	}
	return findTarget(matrix[active_row], target)
}




