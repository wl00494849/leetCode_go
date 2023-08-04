package minsubarraylen

func minSubArrayLen(target int, nums []int) int {
	var result = int(1e9)
	var length = len(nums)
	var sum int
	var i int
	var j int
	for j < length {
		sum += nums[j]
		for sum >= target {
			sum -= nums[i]
			if v := j - i + 1; v < result {
				result = v
			}
			i++
		}
		j++
	}

	if result == int(1e9) {
		return 0
	}

	return result
}
