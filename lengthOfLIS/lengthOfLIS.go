package lengthoflis

func lengthOfLIS(nums []int) int {
	var result int
	var dp []int
	for i := 0; i < len(nums); i++ {
		dp = append(dp, 1)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(dp); j++ {
			if nums[j] > nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		result = max(result, dp[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
