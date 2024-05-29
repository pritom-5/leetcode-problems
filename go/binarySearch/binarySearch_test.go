package binarysearch

import (
	"reflect"
	"testing"
)

func TestSearchInsert(t *testing.T)  {
	testing_items := []struct {
		id string
		target, exp int
		nums []int
	} {
		{id: "1", nums: []int {1,3,5,6}, target: 5, exp: 2},
		{id: "2", nums: []int {1,3,5,6}, target: 2, exp: 1},
		{id: "3", nums: []int {1,3,5,6}, target: 7, exp: 4},
	}

	for _, item := range testing_items {
		t.Run(item.id, func(t *testing.T) {
			got := SearchInsert(item.nums, item.target)

			if got != item.exp {
				t.Errorf("got: %d, exp: %d, target: %d", got, item.exp, item.target)

			}
		})
	}
}

func TestSquaresOfSortedArray (t *testing.T) {
	test_items := []struct{
		id string
		nums, exp []int
	}{
		{id: "1", nums: []int{-4,-1,0,3,10}, exp: []int{0,1,9,16,100}},
		{id: "2", nums: []int{-7,-3,2,3,11}, exp: []int{4,9,9,49,121}},
	}

	for _, item := range test_items {
		t.Run(item.id, func(t *testing.T) {
			got := squaresOfSortedArray(item.nums)
			if !reflect.DeepEqual(got, item.exp) {
				t.Errorf("got: %#v, exp: %#v", got, item.exp)
			}
		})
	}
}




