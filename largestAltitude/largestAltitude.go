package largestaltitude

func largestAltitude(nums []int) int {
	var maxhigh int
	var sum int
	for _, v := range nums {
		sum += v
		if sum > maxhigh {
			maxhigh = sum
		}
	}
	return maxhigh
}
