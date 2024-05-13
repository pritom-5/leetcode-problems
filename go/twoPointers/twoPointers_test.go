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
	}

	for _, item := range test_items {
		t.Run(strconv.Itoa(item.id), func(t *testing.T) {
			got := ValidPallindrome(item.s)
			if got != item.exp {
				t.Errorf("got: %t, exp: %t", got, item.exp)
			}
		})
	}
	
}
