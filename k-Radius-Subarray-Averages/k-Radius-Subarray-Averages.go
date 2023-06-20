package kradiussubarrayaverages

func getAverages(nums []int, k int) []int {
	var size = len(nums)
	var length = k*2 + 1
	var result []int
	// init
	for i := 0; i < len(nums); i++ {
		result = append(result, -1)
	}
	if length > size {
		return result
	}
	// init Sum
	var sum int
	for i := 0; i < length; i++ {
		sum += nums[i]
	}
	result[k] = sum / length
	var end = length
	for i := k + 1; end < size; i++ {
		sum -= nums[i-k-1]
		sum += nums[end]
		result[i] = sum / length
		end++
	}

	return result
}
