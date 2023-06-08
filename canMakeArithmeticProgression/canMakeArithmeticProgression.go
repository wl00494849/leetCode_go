package canmakearithmeticprogression

import "sort"

func canMakeArithmeticProgression(nums []int) bool {
	sort.Ints(nums)
	diff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] != diff {
			return false
		}
	}
	return true
}
