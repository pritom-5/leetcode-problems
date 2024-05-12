// ?????
// what is the difference between []int{} and make([]int, 0) and var name []int
// difference between byte and rune

package array

import (
	"strings"
)

func ContainsDuplicate(nums []int) bool  {
	nums_map := make(map[int]int)

	for _, num := range nums {
		_, ok := nums_map[num]

		if ok {
			return true
		} 

		nums_map[num] = 1
	}

	return false
}

func isAllItemsOfArrayZero (arr [26]int) bool {
	for _, item := range arr {
		if item != 0 {
			return false
		}
	}
	return true
}

func ValidAnagram (s, t string) bool {
	if (len(s) != len(t)) {
		return false
	}

	s_arr := []rune(s)
	t_arr := []rune(t)
	var letters_arr [26]int


	for i, s_letter := range s_arr {
		letters_arr[s_letter - 'a'] += 1
		letters_arr[t_arr[i] - 'a'] -= 1
	}

	return isAllItemsOfArrayZero(letters_arr)

}

func ConcatenationOfArray (nums []int) []int {
	new_nums := make([]int, len(nums) * 2)

	for i, num := range nums {
		new_nums[i] = num
		new_nums[len(nums) + i] = num
	}

	return new_nums
}

func GreatestRightSide (arr []int) []int {
	output_arr := make([]int, len(arr))
	greatest_item := arr[len(arr) - 1]

	output_arr[len(arr) - 1] = -1 

	for i := len(arr) - 2; i >= 0; i-- {
		output_arr[i] = greatest_item
		num := arr[i]
		if (num > greatest_item) {
			greatest_item = num
		}
	}

	return output_arr

}


func IsSubsequence (s, t string) bool {
	left, right := 0, 0

	for (left <= right && left < len(s) && right < len(t)){
		left_ch := s[left]
		right_ch := t[right]

		// fmt.Printf("left_ch: %d,right_ch: %d",left_ch,right_ch)

		if (left_ch == right_ch) {
			left += 1
		}
		right += 1
	}

	return left == len(s)
}



func LengthOfLastWord (s string) int {
	trimed_s := strings.Trim(s, " ")
	split_s_slice := strings.Split(trimed_s, " ")
	return len(split_s_slice[len(split_s_slice) - 1])
}


func TwoSum(nums []int, target int) [2]int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if (nums[i] + nums[j] == target) {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{}
} 

func smallestWord (strs []string) string {
	smallest_word := strs[0]

	for _, str := range strs {
		if (len(str) < len(smallest_word)) {
			smallest_word = str
		}
	}
	return smallest_word
}

func LongestCommonPrefix(strs []string) string {

	prefix := ""
	smallest_word := smallestWord(strs)

	for i, letter := range smallest_word {
		for _, str := range strs {

			if (rune(str[i]) != letter) {
				return prefix
			}
		}
		prefix += string(letter)
	}

	return prefix
}

func PascalsTriangle (numRows int) [][]int {
	var pascals_triangle [][]int = [][]int{{1}}

	for i := 1; i < numRows; i++ {
		prev_row := pascals_triangle[i - 1]
		curr_row := make([]int, i + 1)
		for j := 0; j <= i; j++ {
			var val_1, val_2 int

			if (j <= 0) {
				val_1 = 0
			} else {
				val_1 = prev_row[j - 1]
			}

			if (j >= len(prev_row)) {
				val_2 = 0
			} else {
				val_2 = prev_row[j]
			}

			curr_row[j] = val_1 + val_2
		}
		pascals_triangle = append(pascals_triangle, curr_row)
	}

	return pascals_triangle

}







