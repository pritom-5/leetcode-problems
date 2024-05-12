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


func TestIsSubsequence(t *testing.T)  {
	test_items := []struct{
		id int
		s, t string
		exp bool
	} {
		{1, "abc", "ahbgdc", true},
		{2, "axc", "ahbgdc", false},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := IsSubsequence(item.s, item.t)
			if got != item.exp {
				t.Errorf("got: %t, exp: %t", got, item.exp)
			}
		})
	}
}



func TestLengthOfLastWord(t *testing.T)  {
	test_items := []struct{
		id int
		s string
		exp int
	} {
		{1, "Hello World", 5},
		{2, "   fly me   to   the moon  ", 4},
		{3, "luffy is still joyboy", 6},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := LengthOfLastWord(item.s)
			if got != item.exp {
				t.Errorf("got: %d, exp: %d", got, item.exp)
			}
		})
	}
}

func TestTwoSum(t *testing.T)  {
	test_items := []struct{
		id int
		nums []int
		target int
		exp [2]int
	} {
		{1, []int{2, 7, 11, 15}, 9, [2]int{0, 1}},
		{2, []int{3, 2, 4}, 6, [2]int{1, 2}},
		{3, []int{3, 3}, 6, [2]int{0, 1}},
	}

	for _, item := range test_items {

		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := TwoSum(item.nums, item.target)
			if (!reflect.DeepEqual(got, item.exp)) {
				t.Errorf("got: %#v, exp: %#v", got, item.exp)
			}
		})
	}
}



func TestLongestCommonPrefix(t *testing.T)  {
	test_items := []struct{
		id int
		s []string
		exp string
	} {
		{1, []string{"flower", "flow", "flight"}, "fl"},
		{1, []string{"dog", "racecar", "car"}, ""},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := LongestCommonPrefix(item.s)
			if got != item.exp {
				t.Errorf("got: %s, exp: %s", got, item.exp)
			}
		})
	}
}









