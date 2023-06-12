package twosum

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if v, found := m[target-nums[i]]; found {
			return []int{i, v}
		}
		m[target-nums[i]] = i
	}
	return []int{}
}
