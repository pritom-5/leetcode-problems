package twopointers

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"unicode"
)

func GetCleanString(s string) string {
	clean_string := ""

	for start := 0; start < len(s); start++ {
		letter := rune(s[start])

		println(letter)

		if letter >= 48 && letter <= 57 {
			clean_string += strconv.Itoa(int(letter))
			continue
		}

		letter_lowercase := unicode.ToLower(letter)

		if letter_lowercase >= 97 && letter_lowercase <= 122 {
			clean_string += string(letter_lowercase)
		}
	}

	println(clean_string)

	return clean_string
}

func GetReversedString(s string) string {
	reversed_string := ""

	for i := len(s) - 1; i >= 0; i-- {
		reversed_string += string(s[i])
	}

	return reversed_string
}

func IsPalindrome(s string) bool {
	clean_string := GetCleanString(s)
	reversed_string := GetReversedString(s)
	reverse_clean_string := GetCleanString(reversed_string)

	return clean_string == reverse_clean_string
}

func isDigitOrLetter(letter rune) bool {
	return unicode.IsLetter(letter) || unicode.IsDigit(letter)
}

func IsPalindrome_2(s string) bool {
	arr_s := []rune(s)
	left, right := 0, len(s)-1

	for left <= right {

		left_letter := unicode.ToLower(arr_s[left])
		right_letter := unicode.ToLower(arr_s[right])

		if !isDigitOrLetter(left_letter) {
			left += 1
			continue
		}
		if !isDigitOrLetter(right_letter) {
			right -= 1
			continue
		}

		if left_letter != right_letter {
			return false
		}

		// println(left, right, left_letter, right_letter)
		left += 1
		right -= 1
	}
	return true
}

func checkPalindrome(s string, i, j int) (error, []int) {
	for i <= j {
		left_item := s[i]
		right_item := s[j]

		if left_item != right_item {
			return errors.New("not palindrome"), []int{i, j}
		}

		i += 1
		j -= 1
	}
	return nil, []int{-1, -1}
}

func IsPalindromeII(s string) bool {

	err, left_right := checkPalindrome(s, 0, len(s)-1)
	if err == nil {
		return true
	}
	i, j := left_right[0], left_right[1]

	err, _ = checkPalindrome(s, i+1, j)
	if err == nil {
		return true
	}
	err, _ = checkPalindrome(s, i, j-1)
	if err == nil {
		return true
	}

	return false

}

func DiffOfKScores(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	}

	sort.Ints(nums)
	l, r := 0, k-1

	diff := math.MaxInt32

	for r < len(nums) {
		tmp_diff := nums[r] - nums[l]
		if diff > tmp_diff {
			diff = tmp_diff
		}
		l += 1
		r += 1
	}
	return diff
}

func MergeStringsAlternatively(word1, word2 string) string {
	output_string := ""
	l, r := 0, 0

	for l < len(word1) && r < len(word2) {
		output_string += string(word1[l])
		output_string += string(word2[r])
		l += 1
		r += 1
	}

	if l < len(word1) {
		output_string += word1[l:]
	}
	if r < len(word2) {
		output_string += word2[r:]
	}

	return output_string
}

func ReverseString(s []string) []string {
	s_copy := slices.Clone(s)
	slices.Reverse(s_copy)
	return s_copy
}

func RemoveDuplicateFromSortedArray(nums []int) int {
	count := 1
	curr_unique, l, r := nums[0], 1, 1

	for l < len(nums) {
		if nums[l] != curr_unique {
			count += 1
			curr_unique = nums[l]
			nums[r] = curr_unique
			r += 1
		}
		l += 1
	}

	fmt.Printf("nums: %#v", nums)
	return count
}
