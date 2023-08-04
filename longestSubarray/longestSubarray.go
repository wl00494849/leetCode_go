package longestsubarray

func longestSubarray(nums []int) int {
	var maxlen, lastZero, left int = 0, -1, 0

	for right, v := range nums {
		if v == 0 {
			left = lastZero + 1
			lastZero = right
		}
		maxlen = max(maxlen, right-left)
	}

	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
