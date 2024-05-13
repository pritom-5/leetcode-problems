package twopointers

import (
	"strconv"
	"unicode"
)

func GetCleanString(s string) string {
	clean_string := ""

	for start := 0;start < len(s); start++ { 
		letter := rune(s[start])
		 
		println(letter)


		if (letter >= 48 && letter <= 57) {
			clean_string += strconv.Itoa(int(letter))
			continue
		}

		letter_lowercase := unicode.ToLower(letter)


		if (letter_lowercase >= 97 && letter_lowercase <= 122) {
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

func IsPalindrome(s string) bool  {
	clean_string := GetCleanString(s)
	reversed_string := GetReversedString(s)
	reverse_clean_string := GetCleanString(reversed_string)

	return clean_string == reverse_clean_string
}


func isDigitOrLetter(letter rune) bool {
	return unicode.IsLetter(letter) || unicode.IsDigit(letter)
}

func IsPalindrome_2 (s string) bool {
	arr_s := []rune(s)
	left, right := 0, len(s) - 1


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


















