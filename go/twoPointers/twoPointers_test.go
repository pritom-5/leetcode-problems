package twopointers

import (
	"reflect"
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



func TestDiffOfKScores(t *testing.T)  {
	test_items := []struct{
		id int
		nums []int
		exp,k int
	} {
		{id: 1, nums: []int{90}, k: 1, exp: 0},
		{id: 2, nums: []int{9,4,1,7}, k: 2, exp: 2},
		{id: 3, nums: []int{87063,61094,44530,21297,95857,93551,9918}, k: 6, exp: 74560},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := DiffOfKScores(item.nums, item.k)
			if got != item.exp {
				t.Errorf("got: %d, exp: %d", got, item.exp)
			}
		})
	}
}

func TestMergeStringsAlternatively(t *testing.T)  {
	testing_items := []struct{
		id int
		word1, word2, exp string
	}{
		{1, "abc", "pqr", "apbqcr"},
		{2, "ab", "pqrs", "apbqrs"},
		{3, "abcd", "pq", "apbqcd"},
	}

	for _, item := range testing_items {
		got := MergeStringsAlternatively(item.word1, item.word2)
		if (got != item.exp) {
			t.Errorf("got: %s, exp: %s", got, item.exp)
		}
	}
}

func TestReverseString(t *testing.T)  {
	testing_items := []struct{
		id int
		s, exp []string
	}{
		{1, []string{"h","e","l","l","o"}, []string{"o","l","l","e","h"}},
		{2, []string{"H","a","n","n","a","h"}, []string{"h","a","n","n","a","H"}},
	}

	for _, item := range testing_items {
		got := ReverseString(item.s)
		if (!reflect.DeepEqual(got, item.exp)) {
			t.Errorf("got: %#v, exp: %#v", got, item.exp)
		}
	}
}


func TestRemoveDuplicateFromSortedArray(t *testing.T)  {
	testing_items := []struct{
		id int
		s []int
		exp int
	}{
		{1, []int{1,1,2}, 2},
		{2, []int{0,0,1,1,1,2,2,3,3,4}, 5},
	}

	for _, item := range testing_items {
		got := RemoveDuplicateFromSortedArray(item.s)
		if got != item.exp {
			t.Errorf("got: %d, exp: %d", got, item.exp)
		}
	}
}








