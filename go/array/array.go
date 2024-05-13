// ?????
// what is the difference between []int{} and make([]int, 0) and var name []int
// difference between byte and rune
// map[string]struct{} i dont want to store any value in map

package array

import (
	"slices"
	"strings"
)
func ContainsDuplicate(nums []int) bool {
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

func isAllItemsOfArrayZero(arr [26]int) bool {
	for _, item := range arr {
		if item != 0 {
			return false
		}
	}
	return true
}

func ValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s_arr := []rune(s)
	t_arr := []rune(t)
	var letters_arr [26]int

	for i, s_letter := range s_arr {
		letters_arr[s_letter-'a'] += 1
		letters_arr[t_arr[i]-'a'] -= 1
	}

	return isAllItemsOfArrayZero(letters_arr)

}

func ConcatenationOfArray(nums []int) []int {
	new_nums := make([]int, len(nums)*2)

	for i, num := range nums {
		new_nums[i] = num
		new_nums[len(nums)+i] = num
	}

	return new_nums
}

func GreatestRightSide(arr []int) []int {
	output_arr := make([]int, len(arr))
	greatest_item := arr[len(arr)-1]

	output_arr[len(arr)-1] = -1

	for i := len(arr) - 2; i >= 0; i-- {
		output_arr[i] = greatest_item
		num := arr[i]
		if num > greatest_item {
			greatest_item = num
		}
	}

	return output_arr

}

func IsSubsequence(s, t string) bool {
	left, right := 0, 0

	for left <= right && left < len(s) && right < len(t) {
		left_ch := s[left]
		right_ch := t[right]

		// fmt.Printf("left_ch: %d,right_ch: %d",left_ch,right_ch)

		if left_ch == right_ch {
			left += 1
		}
		right += 1
	}

	return left == len(s)
}

func LengthOfLastWord(s string) int {
	trimed_s := strings.Trim(s, " ")
	split_s_slice := strings.Split(trimed_s, " ")
	return len(split_s_slice[len(split_s_slice)-1])
}

func TwoSum(nums []int, target int) [2]int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{}
}

func smallestWord(strs []string) string {
	smallest_word := strs[0]

	for _, str := range strs {
		if len(str) < len(smallest_word) {
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

			if rune(str[i]) != letter {
				return prefix
			}
		}
		prefix += string(letter)
	}

	return prefix
}

func PascalsTriangle(numRows int) [][]int {
	var pascals_triangle [][]int = [][]int{{1}}

	for i := 1; i < numRows; i++ {
		prev_row := pascals_triangle[i-1]
		curr_row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			var val_1, val_2 int

			if j <= 0 {
				val_1 = 0
			} else {
				val_1 = prev_row[j-1]
			}

			if j >= len(prev_row) {
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

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func RemoveElement(nums []int, val int) []int {
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left] == val {
			swap(nums, left, right)
			right -= 1
			continue
		}
		left += 1
	}

	return nums
}

func strippedEmail(email string) string {
	parts := strings.Split(email, "@")
	should_break := false

	updated_email_first_part := ""

	for _, letter := range parts[0] {
		switch letter {
		case '.':
			continue
		case '+':
			should_break = true
		default:
			updated_email_first_part += string(letter)
		}

		if should_break {
			break
		}
	}

	return updated_email_first_part + "@" + parts[1]
}

func strippedEmail_2(email string) string {
	parts := strings.Split(email, "@")
	local, domain := parts[0], parts[1]

	local = strings.ReplaceAll(local, ".", "")
	plus_idx := strings.Index(local, "+")

	if plus_idx >= 0 {
		local = local[:plus_idx]
	}

	return local + "@" + domain
}

func ValidEmails(emails []string) int {
	email_map := make(map[string]int)
	counter := 0

	for _, email := range emails {
		stripped_email := strippedEmail_2(email)
		_, ok := email_map[stripped_email]
		if !ok {
			email_map[stripped_email] = 1
			counter += 1
		}
	}
	return counter
}

func getLetterMap(s, t string) map[rune]rune {
	m := make(map[rune]rune)

	for i, l := range s {
		m[l] = rune(t[i])
	}

	return m
}

func IsomorphicStrings(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s_to_t_map := make(map[rune]rune)
	t_to_s_map := make(map[rune]rune)

	for i, l := range s {
		_, ok_s := s_to_t_map[l]

		if ok_s {
			if s_to_t_map[l] != rune(t[i]) {
				return false
			}
		} else {
			s_to_t_map[l] = rune(t[i])
		}

		_, ok_t := t_to_s_map[rune(t[i])]

		if ok_t {
			if t_to_s_map[rune(t[i])] != l {
				return false
			}
		} else {
			t_to_s_map[rune(t[i])] = l
		}
	}

	return true
}

func CanPlantFlower(flowerbed []int, n int) bool {
	counter := 0

	updateCounter := func(idx int) {
		counter += 1
		flowerbed[idx] = 1

	}

	for i := 0; i < len(flowerbed); i++ {
		switch i {

		case 0:
			if len(flowerbed) == 1 && flowerbed[i] == 0 {
				updateCounter(i)
			} else if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				updateCounter(i)
			}

		case len(flowerbed) - 1:
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
				updateCounter(i)
			}

		default:
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				updateCounter(i)
			}
		}

	}

	return counter >= n
}


func MajorityElement(nums []int) int  {
	nums_map := make(map[int]int)

	for _, num := range nums {
		_, ok := nums_map[num]
		if ok {
			nums_map[num] += 1
		} else {
			nums_map[num] = 1
		}
	}

	majority_element, majority_key := 0, 0

	for k, v := range nums_map {
		if v > majority_element {
			majority_element = v
			majority_key = k
		}
	}
	return majority_key
}
 
func MajorityElement_2(nums []int) int  {
	counter, majority_element := 0, nums[0]

	for _, num := range nums {
		if (num == majority_element) {
			counter += 1
			continue
		}

		if (counter == 0) {
			counter += 1
			majority_element = num
		} else {
			counter -= 1
		}

	}

	return majority_element
}


func NextGreaterElement(nums1, nums2 []int) []int  {
	greater_element_slice := make([]int, len(nums1))

	for i:= range greater_element_slice {
		greater_element_slice[i] = -1
	}

	for i, num1 := range nums1 {
		nums1_idx := slices.Index(nums2, num1)
		for j := nums1_idx; j < len(nums2); j++ {
			num2 := nums2[j]
			if num2 > num1 {
				greater_element_slice[i] = num2
				break
			}
		}


	}
	
	return greater_element_slice
}

func findTotal (nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func FindPivotIndex(nums []int) int  {
	total := findTotal(nums)
	left_total := 0
	for i, num := range nums {
		if left_total == total - left_total - num {
			return i
		}
		left_total += num
	}
	return -1
}



























