package maxprofit

import "math"

func maxProfit(prices []int) int {
	min := math.MaxInt
	var overProfit int
	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if profit := min - prices[i]; profit > overProfit {
			overProfit = profit
		}
	}
	return overProfit
}
