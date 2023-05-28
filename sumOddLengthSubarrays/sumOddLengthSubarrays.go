package sumoddlengthsubarrays

func SumOddLengthSubarrays(nums []int) int {
	var sum int
	odd := 1
	for odd <= len(nums) {
		for i := 0; i+odd <= len(nums); i++ {
			count := i
			for j := 0; j < odd; j++ {
				sum += nums[count]
				count++
			}
		}
		odd += 2
	}
	return sum
}
