package arraysign

func arraySign(nums []int) int {
	var negative int
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		if v < 0 {
			negative++
		}
	}

	if negative%2 == 0 {
		return 1
	}

	return -1
}
