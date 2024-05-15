package slidingwindow

import "testing"

func TestMaxProfit(t *testing.T) {
	test_items := []struct {
		id, exp int
		prices  []int
	}{
		{id: 1, prices: []int{7, 1, 5, 3, 6, 4}, exp: 5},
		{id: 2, prices: []int{7, 6, 4, 3, 1}, exp: 0},
	}

	for _, item := range test_items {
		got := MaxProfit(item.prices)

		if got != item.exp {
			t.Errorf("got: %d, exp: %d", got, item.exp)
		}
	}

}
