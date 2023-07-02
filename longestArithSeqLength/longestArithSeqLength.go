package longestarithseqlength

func longestArithSeqLength(nums []int) int {
	var maxLen int
	// dp[right][diff]
	var dp []map[int]int
	// init dp
	for i := 0; i <= len(nums); i++ {
		dp = append(dp, map[int]int{})
	}
	//right
	for i := 1; i < len(nums); i++ {
		// left
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			dp[i][diff] = dp[j][diff] + 1
			maxLen = max(dp[i][diff], maxLen)
		}
	}
	return maxLen + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
