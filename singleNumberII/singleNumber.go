package singlenumberii

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
