package slidingwindow

import "math"

func MaxProfit(prices []int) int {
	min := math.MaxInt32
	diff := 0

	for _, price := range prices {
		if price > min {
			diff = max(price-min, diff)
		} else {
			min = price
		}
	}
	return diff
}
