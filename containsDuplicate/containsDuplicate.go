package containsduplicate

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 1 {
			return false
		}
		m[nums[i]] = 1
	}
	return true
}
