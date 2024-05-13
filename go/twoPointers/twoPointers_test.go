package twopointers

import (
	"strconv"
	"testing"
)

func TestValidPallindrome(t *testing.T)  {
	test_items := []struct{
		id int
		s string
		exp bool
	} {
		{id: 1, s: "A man, a plan, a canal: Panama", exp: true},
		{id: 2, s: "race a car", exp: false},
		{id: 3, s: "0P", exp: false},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := IsPalindrome_2(item.s)
			if got != item.exp {
				t.Errorf("got: %t, exp: %t", got, item.exp)
			}
		})
	}
	
}

func TestIsPalindromeII(t *testing.T)  {
	test_items := []struct{
		id int
		s string
		exp bool
	} {
		{id: 1, s: "aba", exp: true},
		{id: 2, s: "abca", exp: true},
		{id: 3, s: "abc", exp: false},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := IsPalindromeII(item.s)
			if got != item.exp {
				t.Errorf("got: %t, exp: %t", got, item.exp)
			}
		})
	}
	
}
