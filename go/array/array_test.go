package array
import (
	"reflect"
	"strconv"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	contains_duplicate_test_items := []struct {
		id     int
		nums   []int
		expect bool
	}{
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

func TestValidAnagram(t *testing.T) {

	valid_anagram_test_items := []struct {
		id   int
		s, t string
		exp  bool
	}{
		{1, "anagram", "nagaram", true},
		{2, "rat", "car", false},
		{3, "a", "ab", false},
	}

	for _, item := range valid_anagram_test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := validAnagram_2(item.s, item.t)
			if got != item.exp {
				t.Errorf("got: %t, exp: %t", got, item.exp)
			}

		})
	}

}

func TestConcatenationOfArray(t *testing.T) {

	test_items := []struct {
		id   int
		nums []int
		exp  []int
	}{
		{1, []int{1, 2, 3}, []int{1, 2, 3, 1, 2, 3}},
		{2, []int{1, 2}, []int{1, 2, 1, 2}},
	}

	for _, test_item := range test_items {
		t.Run(strconv.Itoa(test_item.id), func(t *testing.T) {
			got := ConcatenationOfArray(test_item.nums)

			if !reflect.DeepEqual(got, test_item.exp) {
				t.Errorf("got:%#v, exp: %#v", got, test_item.exp)
			}

		})
	}
}

func TestGreatestRightSide(t *testing.T) {
	test_items := []struct {
		id  int
		arr []int
		exp []int
	}{
		{1, []int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
		{2, []int{400}, []int{-1}},
	}

	for _, test_item := range test_items {
		t.Run(strconv.Itoa(test_item.id), func(t *testing.T) {
			got := GreatestRightSide(test_item.arr)
			if !reflect.DeepEqual(got, test_item.exp) {
				t.Errorf("got: %#v, exp: %#v", got, test_item.exp)
			}

		})

	}
}

func TestIsSubsequence(t *testing.T) {
	test_items := []struct {
		id   int
		s, t string
		exp  bool
	}{
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

func TestLengthOfLastWord(t *testing.T) {
	test_items := []struct {
		id  int
		s   string
		exp int
	}{
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

func TestTwoSum(t *testing.T) {
	test_items := []struct {
		id     int
		nums   []int
		target int
		exp    [2]int
	}{
		{1, []int{2, 7, 11, 15}, 9, [2]int{0, 1}},
		{2, []int{3, 2, 4}, 6, [2]int{1, 2}},
		{3, []int{3, 3}, 6, [2]int{0, 1}},
	}

	for _, item := range test_items {

		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := TwoSum(item.nums, item.target)
			if !reflect.DeepEqual(got, item.exp) {
				t.Errorf("got: %#v, exp: %#v", got, item.exp)
			}
		})
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	test_items := []struct {
		id  int
		s   []string
		exp string
	}{
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

func TestPascalsTriangle(t *testing.T) {
	test_items := []struct {
		id      int
		numRows int
		exp     [][]int
	}{
		{1, 5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{2, 1, [][]int{{1}}},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := PascalsTriangle(item.numRows)
			if !reflect.DeepEqual(got, item.exp) {
				t.Errorf("got: %#v, exp: %#v", got, item.exp)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	testing_items := []struct {
		id   int
		nums []int
		val  int
		exp  []int
	}{
		{id: 1, nums: []int{3, 2, 2, 3}, val: 3, exp: []int{2, 2, 3, 3}},
		{id: 2, nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, exp: []int{0, 1, 4, 0, 3, 2, 2, 2}},
	}

	for _, item := range testing_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := RemoveElement(item.nums, item.val)
			if !reflect.DeepEqual(got, item.exp) {
				t.Errorf("got: %#v, exp: %#v", got, item.exp)
			}
		})
	}

}

func TestValidEmails(t *testing.T) {
	test_items := []struct {
		id     int
		emails []string
		exp    int
	}{
		{id: 1, emails: []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}, exp: 2},
		{id: 2, emails: []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}, exp: 3},
	}

	for _, item := range test_items {
		got := ValidEmails(item.emails)
		if got != item.exp {
			t.Errorf("got: %d, exp: %d", got, item.exp)
		}
	}
}

func TestIsomorphicStrings(t *testing.T) {
	test_items := []struct {
		id   int
		s, t string
		exp  bool
	}{
		{id: 1, s: "egg", t: "add", exp: true},
		{id: 2, s: "foo", t: "bar", exp: false},
		{id: 3, s: "paper", t: "title", exp: true},
		{id: 4, s: "bbbaaaba", t: "aaabbbba", exp: false},
	}

	for _, item := range test_items {
		got := IsomorphicStrings(item.s, item.t)
		if got != item.exp {
			t.Errorf("s: %s, t: %s, got: %t, exp: %t", item.s, item.t, got, item.exp)
		}
	}
}

func TestCanPlantFlower(t *testing.T) {
	test_items := []struct {
		id        int
		flowerbed []int
		n         int
		exp       bool
	}{
		{id: 1, flowerbed: []int{1, 0, 0, 0, 1}, n: 1, exp: true},
		{id: 2, flowerbed: []int{1, 0, 0, 0, 1}, n: 2, exp: false},
		{id: 3, flowerbed: []int{0}, n: 1, exp: true},
		{id: 4, flowerbed: []int{0, 0}, n: 2, exp: false},
	}

	for _, item := range test_items {
		got := CanPlantFlower(item.flowerbed, item.n)
		if got != item.exp {
			t.Errorf("flowerbed: %#v, n: %d, got: %t, exp: %t, id: %d", item.flowerbed, item.n, got, item.exp, item.id)
		}
	}
}

func TestMajorityElement(t *testing.T) {
	test_items := []struct {
		id   int
		nums []int
		exp  int
	}{
		{id: 1, nums: []int{3, 2, 3}, exp: 3},
		{id: 2, nums: []int{2, 2, 1, 1, 1, 2, 2}, exp: 2},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := MajorityElement(item.nums)

			if got != item.exp {
				t.Errorf("got: %d, exp: %d", got, item.exp)
			}
		})

	}
}

func TestMajorityElement_2(t *testing.T) {
	test_items := []struct {
		id   int
		nums []int
		exp  int
	}{
		{id: 1, nums: []int{3, 2, 3}, exp: 3},
		{id: 2, nums: []int{2, 2, 1, 1, 1, 2, 2}, exp: 2},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := MajorityElement_2(item.nums)

			if got != item.exp {
				t.Errorf("got: %d, exp: %d", got, item.exp)
			}
		})

	}
}

func TestNextGreaterElement(t *testing.T) {
	test_items := []struct {
		id                int
		nums1, nums2, exp []int
	}{
		{id: 1, nums1: []int{4, 1, 2}, nums2: []int{1, 3, 4, 2}, exp: []int{-1, 3, -1}},
		{id: 2, nums1: []int{2, 4}, nums2: []int{1, 2, 3, 4}, exp: []int{3, -1}},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := NextGreaterElement(item.nums1, item.nums2)

			if !reflect.DeepEqual(got, item.exp) {
				t.Errorf("got: %#v, exp: %#v", got, item.exp)
			}
		})
	}
}

func TestFindPivotIndex(t *testing.T) {
	test_items := []struct {
		id   int
		nums []int
		exp  int
	}{
		{id: 1, nums: []int{1, 7, 3, 6, 5, 6}, exp: 3},
		{id: 2, nums: []int{1, 2, 3}, exp: -1},
		{id: 3, nums: []int{2, 1, -1}, exp: 0},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := FindPivotIndex(item.nums)

			if got != item.exp {
				t.Errorf("got: %#v, exp: %#v", got, item.exp)
			}
		})
	}
}

func TestFindDisappearedNumbers(t *testing.T) {
	test_items := []struct {
		id        int
		nums, exp []int
	}{
		{id: 1, nums: []int{4, 3, 2, 7, 8, 2, 3, 1}, exp: []int{5, 6}},
		{id: 1, nums: []int{1, 1}, exp: []int{2}},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := FindDisappearedNumbers(item.nums)

			if !reflect.DeepEqual(got, item.exp) {
				t.Errorf("got: %#v, exp: %#v", got, item.exp)
			}
		})
	}
}

func TestNosBalloons(t *testing.T) {
	test_items := []struct {
		id, exp int
		text    string
	}{
		{id: 1, text: "nlaebolko", exp: 1},
		{id: 2, text: "loonbalxballpoon", exp: 2},
		{id: 3, text: "leetcode", exp: 0},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := NosBalloons(item.text)

			if got != item.exp {
				t.Errorf("got: %d, exp: %d", got, item.exp)
			}
		})
	}
}

func TestWordPattern(t *testing.T) {
	test_items := []struct {
		id         int
		pattern, s string
		exp        bool
	}{
		{id: 1, pattern: "abba", s: "dog cat cat dog", exp: true},
		{id: 2, pattern: "abba", s: "dog cat cat fish", exp: false},
		{id: 3, pattern: "aaaa", s: "dog cat cat dog", exp: false},
		{id: 4, pattern: "aaa", s: "cat cat cat cat", exp: false},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := WordPattern(item.pattern, item.s)

			if got != item.exp {
				t.Errorf("got: %t, exp: %t", got, item.exp)
			}
		})
	}
}

func TestFindDiffTwoArray(t *testing.T) {
	test_items := []struct {
		id           int
		nums1, nums2 []int
		exp          [2][]int
	}{
		{id: 1, nums1: []int{1, 2, 3}, nums2: []int{2, 4, 6}, exp: [2][]int{{1, 3}, {4, 6}}},
		{id: 2, nums1: []int{1, 2, 3, 3}, nums2: []int{1, 1, 2, 2}, exp: [2][]int{{3}, {}}},
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := FindDiffTwoArray(item.nums1, item.nums2)

			if !reflect.DeepEqual(got, item.exp) {
				t.Errorf("got: %#v, exp: %#v", got, item.exp)
			}
		})
	}
}

func TestIsMonotonic(t *testing.T) {
	test_items := []struct {
		id string
		nums []int
		exp bool
	} {
		{id: "01", nums: []int {1, 2, 2, 3}, exp: true},
		{id: "02", nums: []int {6, 5, 5, 4}, exp: true},
		{id: "03", nums: []int {1, 3, 2}, exp: false},
	}

	for _, item := range test_items {
		t.Run(item.id, func(t *testing.T) {
			got := isMonotonic(item.nums)

			if got != item.exp {
				t.Errorf("got: %t, exp: %t", got, item.exp)
			}
		})

	}
}


func TestLargestGoodInteger(t *testing.T) {
	test_data := []struct {
		id string 
		num, exp string
	} {
		{id: "1", num: "6777133339", exp: "777"},
		{id: "2", num: "2300019", exp: "000"},
		{id: "3", num: "42352338", exp: ""},
	}

	for _, item := range test_data {
		t.Run(item.id, func(t *testing.T) {
			got := largestGoodInteger(item.num)
			if got != item.exp {
				t.Errorf("got: %s, exp: %s", got, item.exp)
			}
		})
	}
}

func TestMaxScore(t *testing.T) {
	test_data := []struct {
		id, num string 
		exp int
	} {
		{id: "1", num: "011101", exp: 5},
	}

	for _, item := range test_data {
		t.Run(item.id, func(t *testing.T) {
			got := maxScore(item.num)
			if got != item.exp {
				t.Errorf("got: %d, exp: %d", got, item.exp)
			}
		})
	}
}



func TestPathCrossed(t *testing.T) {
	test_data := []struct {
		id, path string 
		exp bool
	} {
		// {id: "1", path: "NES", exp: false},
		{id: "2", path: "NESWW", exp: true},
	}

	for _, item := range test_data {
		t.Run(item.id, func(t *testing.T) {
			got := pathCrossed(item.path)
			if got != item.exp {
				t.Errorf("got: %t, exp: %t", got, item.exp)
			}
		})
	}
}
