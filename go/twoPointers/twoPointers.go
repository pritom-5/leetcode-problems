package twopointers

import (
	"fmt"
	"unicode"
)

func GetCleanString(s string) string {
	clean_string := ""

	for start := 0;start < len(s); start++ { 
		letter := rune(s[start])
		letter_lowercase := unicode.ToLower(letter)

		fmt.Printf("%d lower: %d\n", letter, letter_lowercase)

		if (letter_lowercase >= 97 && letter_lowercase <= 122) {
			print("hit")
			clean_string += string(letter_lowercase)
		}
	}

	return clean_string
}

func GetReversedString(s string) string {
	reversed_string := ""

	for i := len(s) - 1; i >= 0; i-- {
		reversed_string += string(s[i])
	}

	return reversed_string
}

func ValidPallindrome(s string) bool  {
	clean_string := GetCleanString(s)
	reversed_string := GetReversedString(s)
	reverse_clean_string := GetCleanString(reversed_string)

	fmt.Printf("clean_string: %s, \treverse_clean_string: %s\n",clean_string,reverse_clean_string)

	return clean_string == reverse_clean_string
}
