package array

import (
	"reflect"
	"strconv"
	"testing"
)

func TestContainsDuplicate(t *testing.T)  {
	contains_duplicate_test_items := []struct{
		id int
		nums []int
		expect bool
	} {
		{1, []int{1, 2, 3, 1}, true},
		{2, []int{1, 2, 3}, false},
		{3, []int{1, 2, 3, 4}, false},
	}

	for _, test_item := range contains_duplicate_test_items {
		t.Run(strconv.Itoa(test_item.id), func(t *testing.T) {
			got := ContainsDuplicate(test_item.nums)

			if got != test_item.expect {
				t.Errorf("got: %t, expect: %t", got, test_item.expect)
			}

		})
	}
	
}

func TestValidAnagram (t *testing.T) {

	valid_anagram_test_items := []struct {
		id int
		s, t string
		exp bool
	}{
		{1, "anagram", "nagaram", true},
		{2, "rat", "car", false},
		{3, "a", "ab", false},
	}

	for _, item := range valid_anagram_test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := ValidAnagram(item.s, item.t)
			if (got != item.exp) {
				t.Errorf("got: %t, exp: %t", got, item.exp)
			}

		})
	}

}

func TestConcatenationOfArray(t *testing.T)  {

	test_items := []struct{
		id int
		nums []int
		exp []int
	} {
		{1, []int{1, 2, 3}, []int{1, 2, 3, 1, 2, 3}},
		{2, []int{1, 2}, []int{1, 2, 1, 2}},
	}

	for _, test_item := range test_items {
		t.Run(strconv.Itoa(test_item.id), func(t *testing.T) {
			got := ConcatenationOfArray(test_item.nums)

			if (!reflect.DeepEqual(got, test_item.exp)) {
				t.Errorf("got:%#v, exp: %#v", got, test_item.exp)
			}

		})
	}
}




func TestGreatestRightSide(t *testing.T)  {
	test_items := []struct {
		id int
		arr []int
		exp []int
	} {
		{1, []int{17, 18, 5, 4, 6, 1}, []int {18, 6, 6, 6, 1, -1}},
		{2, []int{400}, []int {-1}},
	}

	for _, test_item := range test_items {
		t.Run(strconv.Itoa(test_item.id), func(t *testing.T) {
			got := GreatestRightSide(test_item.arr)
			if (!reflect.DeepEqual(got, test_item.exp)) {
				t.Errorf("got: %#v, exp: %#v", got, test_item.exp)
			}

		})
	
	}
}








