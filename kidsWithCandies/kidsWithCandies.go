package kidswithcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	res := make([]bool, len(candies))
	for _, v := range candies {
		if v > max {
			max = v
		}
	}
	for i, v := range candies {
		if v+extraCandies >= max {
			res[i] = true
		}
	}

	return res
}
