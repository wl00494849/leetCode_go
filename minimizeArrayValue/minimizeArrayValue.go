package minimizearrayvalue

func max(nums []int) int {
	var a int
	for _, v := range nums {
		if a < v {
			a = v
		}
	}
	return a
}

func minimizeArrayValue(nums []int) int {
	left, right := nums[0], max(nums)
	for left < right {
		pivot := left + (right-left)/2
		var buff int
		sign := true
		for i := 0; i < len(nums); i++ {
			if pivot < nums[i] {
				buff -= nums[i] - pivot
			} else {
				buff += pivot - nums[i]
			}
			if buff < 0 {
				sign = false
				break
			}
		}
		if sign {
			right = pivot
		} else {
			left = pivot + 1
		}
	}
	return right
}
