package longestarithseqlength

func longestArithSeqLength(nums []int) int {
	var max int
	for i := 1; i < len(nums); i++ {
		diff := nums[i-1] - nums[i]
		m := map[int]int{}
		m[diff]++
		for _, v := range m {
			if v > max {
				max = v
			}
		}
	}

	return max
}
