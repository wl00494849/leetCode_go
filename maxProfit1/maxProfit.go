package maxprofit1

func maxProfit(price []int) int {
	var profit int
	for i := 1; i < len(price); i++ {
		diff := price[i] - price[i-1]
		if diff > 0 {
			profit += diff
		}
	}
	return profit
}
